package visualization

import "fmt"

type GraphFilePathNoExtension struct {
	Path string
}

func (g GraphFilePathNoExtension) Error() string {
	return fmt.Sprintf("The path of the output graph {%v} as no extension", g.Path)
}
