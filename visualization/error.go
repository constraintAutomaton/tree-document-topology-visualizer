package visualization

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type GraphFilePathNoExtension struct {
	Path string
}

func (g GraphFilePathNoExtension) Error() string {
	return fmt.Sprintf("the path of the output graph {%v} as no extension", g.Path)
}

type GraphFilePathInvalidFormat struct {
	format string
}

func (g GraphFilePathInvalidFormat) Error() string {
	return fmt.Sprintf("the format provided {%v} is invalid, we support {%v}", g.format, maps.Keys(graphvizValidFormat))
}
