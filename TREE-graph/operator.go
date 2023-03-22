package treegraph

var operatorMapping = map[string]string{
	"tree:PrefixRelation":               "PrefixRelation",
	"tree:SubstringRelation":            "SubstringRelation",
	"tree:SuffixRelation":               "SuffixRelation",
	"tree:GreaterThanRelation":          ">",
	"tree:GreaterThanOrEqualToRelation": "≥",
	"tree:LessThanRelation":             "<",
	"tree:LessThanOrEqualToRelation":    "≤",
	"tree:EqualToRelation":              "=",
	"tree:GeospatiallyContainsRelation": "GeospatiallyContainsRelation",
}
