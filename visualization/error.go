package visualization

import (
	"fmt"

	"golang.org/x/exp/maps"
)

// GraphFilePathNoExtensionError is an error that describes that the path of the graph defined by the user has no extension.
type GraphFilePathNoExtensionError struct {
	Path string
}

func (g GraphFilePathNoExtensionError) Error() string {
	return fmt.Sprintf("the path of the output graph {%v} has no extension", g.Path)
}

// GraphFilePathInvalidFormatError is an error that describes that the format of the graph defined by the user is not supported.
type GraphFilePathInvalidFormatError struct {
	format string
}

func (g GraphFilePathInvalidFormatError) Error() string {
	return fmt.Sprintf("the format provided {%v} is invalid, we support {%v}", g.format, maps.Keys(graphvizValidFormat))
}
