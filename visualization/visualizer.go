package visualization

// Visualizer is a general interface for the visualization of the TREE graph.
type Visualizer interface {
	GenerateFile(path string) error
}
