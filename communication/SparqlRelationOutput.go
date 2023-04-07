package communication

// SparqlRelationOutput is the output of the SPARQL query to get all the relations of a TREE view
type SparqlRelationOutput struct {
	Operator string
	Value    string
	NextNode string
	Node     string
}
