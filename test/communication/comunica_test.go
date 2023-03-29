package comunication_test

import (
	"testing"
	"tree-document-topology-visualizer/communication"
)

func TestGetTreeRelations(t *testing.T) {
	A_DATASOURCE := "foo" // doesn't matter
	A_LIMIT := uint(0)    // doesn't matter always return 5 results
	communication.SetComunicaBinaryPath("./mock.mjs")
	sparqlOutput, err := communication.GetTreeRelation(A_DATASOURCE, A_LIMIT)
	if err != nil {
		t.Fatal(err)
	}

	if len(sparqlOutput) != 5 {
		t.Fatalf("should had 5 outputs but has %v output", len(sparqlOutput))
	}
}
