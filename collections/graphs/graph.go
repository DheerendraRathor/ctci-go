package graphs

type Graph struct {
    nodes map[string]bool       // Set of nodes

    // Map of neighbours. Key is source node and value another map where key is destination node and value is cost
    edges map[string]map[string]int

    // Destination map represents nodes which are pointing edges towards key
    reverseEdges map[string]map[string]bool
}

func New() *Graph {
    return &Graph{
        nodes: make(map[string]bool),
        edges: make(map[string]map[string]int),
        reverseEdges: make(map[string]map[string]bool),
    }
}

// This initialized edges map for new node. If node already exists then it won't do anything
func (g *Graph) AddNode(name string) {
    g.nodes[name] = true
    if g.edges[name] == nil {
        g.edges[name] = make(map[string]int)
    }
    if g.reverseEdges[name] == nil {
        g.reverseEdges[name] = make(map[string]bool)
    }
}

func (g *Graph) DoesNodeExists(name string) bool {
    return g.nodes[name]
}

// Caveat: This function will add nodes if not present already
func (g *Graph) AddEdge(source, destination string, cost int) {
    g.AddNode(source)
    g.AddNode(destination)
    g.edges[source][destination] = cost
    g.reverseEdges[destination][source] = true
}

// Remove a node from graph
// This will also delete all edges coming to node and all edges going from node
func (g *Graph) RemoveNode(name string) {
    delete(g.nodes, name)           // Delete from set of nodes
    reverseEdges := g.reverseEdges[name]
    for k, _ := range reverseEdges {
        delete(g.edges[k], name)            // Delete all edges which are pointing to node
    }
    delete(g.reverseEdges, name)
}

func (g *Graph) RemoveEdge(source, destination string) {
    delete(g.edges[source], destination)
    delete(g.reverseEdges[destination], source)
}

// Return all edges going from node name
// Second return value indicates whether node is present in graph or not
func (g *Graph) GetEdges(name string) (map[string]int, bool) {
    return g.edges[name], g.nodes[name]
}

func (g *Graph) BreadthFirstTraversal(startNode string) {
    // TODO (Dheerendra): Implement this properly
}
