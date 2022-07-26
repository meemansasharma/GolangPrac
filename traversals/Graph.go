package main

import "fmt"

type Graph struct {
	V         int
	vertexMap map[int][]int
}

func (graph *Graph) addVertex(newVertex int) {
	graph.V++
	graph.vertexMap[newVertex] = make([]int, 0)

}

func (graph *Graph) addEdge(source int, destination int) {
	graph.vertexMap[source] = append(graph.vertexMap[source], destination)
	graph.vertexMap[destination] = append(graph.vertexMap[destination], source)
}

func (graph *Graph) printGraph() {
	fmt.Println("The graph")
	for key, value := range graph.vertexMap {
		fmt.Println(key, "-->", value)
	}
}
func (graph *Graph) getVertexCount() {
	fmt.Println("The count of the vertex in the graph is : ", graph.V)
}
func (graph *Graph) Creategraph() {

	graph.vertexMap = make(map[int][]int)
	graph.V = 0
	fmt.Println("Create graph")
	fmt.Println("Enter vertex")
	graph.addVertex(0)
	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(3)
	graph.addVertex(4)
	graph.addVertex(5)
	graph.addVertex(6)
	graph.addEdge(3, 1)
	graph.addEdge(3, 4)
	graph.addEdge(4, 2)
	graph.addEdge(4, 5)
	graph.addEdge(1, 2)
	graph.addEdge(1, 0)
	graph.addEdge(0, 2)
	graph.addEdge(6, 5)
	graph.printGraph()
}
