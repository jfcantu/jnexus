package db

// Network represents all nodes and links
type Network struct {
	Nodes []*Node `json:"nodes"`
	Links []*Link `json:"links"`
}

// Node represents a node in the graph
type Node struct {
	ID         int               `json:"id"`
	Labels     []string          `json:"labels"`
	Properties map[string]string `json:"properties"`
}

// Link represents a link between nodes
type Link struct {
	ID         int               `json:"id"`
	Start      int               `json:"start"`
	End        int               `json:"end"`
	Type       string            `json:"type"`
	Properties map[string]string `json:"properties"`
}
