package communication

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const BINARY_PATH = "./comunica-feature-link-traversal/engines/query-sparql-link-traversal/bin/query.js"
const PATH_FILE_SPARQL_QUERY_GETTING_TREE_RELATIONS = "./communication/comunica_getting_relation_query"

func GetTreeRelation(datasource string) ([]SparqlRelationOutput, error) {
	command := fmt.Sprintf("node %v %v -f %v --lenient", BINARY_PATH, datasource, PATH_FILE_SPARQL_QUERY_GETTING_TREE_RELATIONS)
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
