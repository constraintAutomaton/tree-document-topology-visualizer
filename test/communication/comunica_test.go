package comunication_test

import (
	"testing"
	"tree-document-topology-visualizer/communication"
)

func TestGetTreeRelations(t *testing.T) {
	A_DATASOURCE := "foo"                             // doesn't matter
	A_LIMIT := uint(0)                                // doesn't matter always return 5 results
	communication.SetComunicaBinaryPath("./mock.mjs") // the program will fail if no query or data source was passed
	sparqlOutput, err := communication.GetTreeRelation(A_DATASOURCE, A_LIMIT)
	if err != nil {
		t.Fatal(err)
	}

	if len(sparqlOutput) != 5 {
		t.Fatalf("should had 5 outputs but has %v output", len(sparqlOutput))
	}
}

func TestGetTreeRelationProgramReturnError(t *testing.T) {
	A_DATASOURCE := "foo" // doesn't matter
	A_LIMIT := uint(0)    //doesn't matter
	communication.SetComunicaBinaryPath("./mock_fail.js")
	_, err := communication.GetTreeRelation(A_DATASOURCE, A_LIMIT)
	if err == nil {
		t.Fatal("Since the program fail, we should return an error")
	}
	programFailError, valid := err.(communication.ProgramFailedError)
	if !valid {
		t.Errorf("impossible to cast the error into a program failure, the error is {%v}", err.Error())
	}

	if programFailError.Program != "./mock_fail.js" {
		t.Errorf("didn't failed on the right program, should be {%v} but was {%v}", "./mock_exit_code_1.js", programFailError.Program)
	}

}

func TestGetTreeRelationProgramMalformedJson(t *testing.T) {
	A_DATASOURCE := "foo" // doesn't matter
	A_LIMIT := uint(0)    //doesn't matter
	communication.SetComunicaBinaryPath("./mock_malformed_json.js")
	_, err := communication.GetTreeRelation(A_DATASOURCE, A_LIMIT)
	if err == nil {
		t.Fatal("Since the program fail, we should return an error")
	}
	_, valid := err.(communication.UnableToDecodeJsonError)
	if !valid {
		t.Errorf("impossible to cast the error into a program failure, the error is {%v}", err.Error())
	}

}
