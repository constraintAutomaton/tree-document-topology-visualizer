package treegraph

// OperatorMapping returned a map associating the IRI of an operator of a tree:relation with a more ergonomic label.
func OperatorMapping() *map[string]string {
	return &operatorMapping
}

var operatorMapping = map[string]string{
	"https://w3id.org/tree#PrefixRelation":               "HasPrefixRelationWith",
	"https://w3id.org/tree#SubstringRelation":            "HasSubstringRelationWith",
	"https://w3id.org/tree#SuffixRelation":               "HasSuffixRelationWith",
	"https://w3id.org/tree#GreaterThanRelation":          ">",
	"https://w3id.org/tree#GreaterThanOrEqualToRelation": "≥",
	"https://w3id.org/tree#LessThanRelation":             "<",
	"https://w3id.org/tree#LessThanOrEqualToRelation":    "≤",
	"https://w3id.org/tree#EqualToRelation":              "=",
	"https://w3id.org/tree#GeospatiallyContainsRelation": "GeospatiallyContainsRelation",
}
