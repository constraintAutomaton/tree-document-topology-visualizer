package visualization

type Visualizer interface {
	GenerateFile(path string) error
}
