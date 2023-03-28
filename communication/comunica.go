package communication

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	DEFAULT_BINARY_PATH           = "./comunica-feature-link-traversal/engines/query-sparql-link-traversal/bin/query-dynamic.js"
	SPARQL_QUERY_GET_ALL_RELATION = `
PREFIX tree: <https://w3id.org/tree#>
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

func GetTreeRelation(datasource string, limit uint) ([]SparqlRelationOutput, error) {
	command := fmt.Sprintf("node %v %v -q %v --lenient", BINARY_PATH, datasource, fmt.Sprintf(SPARQL_QUERY_GET_ALL_RELATION, limit))
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "COMUNICA_CONFIG='./comunica-feature-link-traversal/engines/config-query-sparql-link-traversal/config/config-tree.json'")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	sparqlRelation := []SparqlRelationOutput{}
	if err := json.NewDecoder(stdout).Decode(&sparqlRelation); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}
	return sparqlRelation, nil
}

type SparqlRelationOutput struct {
	Operator string
	Value    string
	NextNode string
	Node     string
}

func SetComunicaBinaryPath(path string) {
	BINARY_PATH = path
}
