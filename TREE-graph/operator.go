package treegraph

func OperatorMapping() *map[string]string {
	return &operatorMapping
}

var operatorMapping = map[string]string{
	"https://w3id.org/tree#PrefixRelation":               "IsPrefixRelationOf",
	"https://w3id.org/tree#SubstringRelation":            "IsSubstringRelationOf",
	"https://w3id.org/tree#SuffixRelation":               "IsSuffixRelationOf",
	"https://w3id.org/tree#GreaterThanRelation":          ">",
	"https://w3id.org/tree#GreaterThanOrEqualToRelation": "≥",
	"https://w3id.org/tree#LessThanRelation":             "<",
	"https://w3id.org/tree#LessThanOrEqualToRelation":    "≤",
	"https://w3id.org/tree#EqualToRelation":              "=",
	"https://w3id.org/tree#GeospatiallyContainsRelation": "GeospatiallyContainsRelation",
}
