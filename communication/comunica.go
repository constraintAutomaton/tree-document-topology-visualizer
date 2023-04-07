package communication

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

const (
	DEFAULT_BINARY_PATH           = "./comunica-js/index.mjs"
	SPARQL_QUERY_GET_ALL_RELATION = `PREFIX tree: <https://w3id.org/tree#>
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>

SELECT ?node ?nextNode ?operator ?value WHERE {
  ?node tree:relation ?relation .
  ?relation tree:node ?nextNode .
  
  ?relation rdf:type ?operator.
  ?relation tree:value ?value .
}LIMIT %v`
)

var (
	BINARY_PATH = DEFAULT_BINARY_PATH
)

// GetTreeRelation fetch the relation of a TREE view with a SPARQL query using the SPARQL query engine Comunica.
func GetTreeRelation(datasource string, limit uint) ([]SparqlRelationOutput, error) {
	command := fmt.Sprintf(`node %v -d %v -q %v`, BINARY_PATH, datasource, fmt.Sprintf(SPARQL_QUERY_GET_ALL_RELATION, limit))
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("was not able to get the stdout with error {%v}", err.Error())
	}
	defer stdout.Close()
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("was not able to get the stdout with error {%v}", err.Error())
	}
	defer stderr.Close()

	if err := cmd.Start(); err != nil {
		return nil, ProgramFailedError{Program: BINARY_PATH, Message: err.Error()}
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, stderr)
	if err != nil {
		return nil, fmt.Errorf("unable to copy the error buffer return error {%v}", err)
	}
	if stringError := buf.String(); stringError != "" {
		return nil, ProgramFailedError{Program: BINARY_PATH, Message: stringError}
	}

	sparqlRelation := []SparqlRelationOutput{}
	if err := json.NewDecoder(stdout).Decode(&sparqlRelation); err != nil {
		return nil, UnableToDecodeJsonError{Message: err.Error()}
	}

	if err := cmd.Wait(); err != nil {
		return nil, ProgramFailedError{Program: BINARY_PATH, Message: err.Error()}
	}
	return sparqlRelation, nil
}

// SetComunicaBinaryPath set a new path for the JavaScript binary of Comunica.
func SetComunicaBinaryPath(path string) {
	BINARY_PATH = path
}
