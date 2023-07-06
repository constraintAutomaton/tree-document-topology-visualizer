package communication

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	DEFAULT_BINARY_PATH = "./comunica-js/index.mjs"
)

var (
	BINARY_PATH = DEFAULT_BINARY_PATH
)

// GetTreeRelation fetch the relation of a TREE view with a SPARQL query using the SPARQL query engine Comunica.
func GetTreeRelation(datasource string, limit uint) ([]SparqlRelationOutput, error) {
	command := fmt.Sprintf(`node --no-warnings %v -d %v -l %v`, BINARY_PATH, datasource, limit)
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)
	result_saver := saveSparqlResult{Results: []SparqlRelationOutput{}}

	cmd.Stdout = &result_saver

	if err := cmd.Start(); err != nil {
		//return nil, ProgramFailedError{Program: BINARY_PATH, Message: err.Error()}
	}

	if err := cmd.Wait(); err != nil {
		//return nil, ProgramFailedError{Program: BINARY_PATH, Message: err.Error()}
	}
	return result_saver.Results, nil
}

// SetComunicaBinaryPath set a new path for the JavaScript binary of Comunica.
func SetComunicaBinaryPath(path string) {
	BINARY_PATH = path
}

type saveSparqlResult struct {
	Results []SparqlRelationOutput
}

func (s *saveSparqlResult) Write(p []byte) (n int, err error) {
	currentRelation := SparqlRelationOutput{}
	if err := json.Unmarshal(p, &currentRelation); err != nil {
		return 0, err
	}
	fmt.Println(currentRelation)
	s.Results = append(s.Results, currentRelation)
	return os.Stdout.Write(p)
}
