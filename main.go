package main

import (
	"fmt"
	"math/rand"

	"github.com/tctromp/graphpattern/tgraph"
	"github.com/tctromp/graphpattern/ui"
)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println("Hello world")

	step1 := tgraph.Step{A: 1, B: 2}
	step2 := tgraph.Step{A: 1, B: 3}
	//step3 := tgraph.Step{A: 2, B: 3}

	newStep1 := tgraph.Step{A: 1, B: 2}
	newStep2 := tgraph.Step{A: 1, B: 4}
	newStep3 := tgraph.Step{A: 2, B: 4}
	newStep4 := tgraph.Step{A: 3, B: 4}

	v1 := tgraph.Vertex{Name: "1", InEdges: make([]*tgraph.Vertex, 0), OutEdges: make([]*tgraph.Vertex, 0)}
	v2 := tgraph.Vertex{Name: "2", InEdges: make([]*tgraph.Vertex, 0), OutEdges: make([]*tgraph.Vertex, 0)}
	v3 := tgraph.Vertex{Name: "3", InEdges: make([]*tgraph.Vertex, 0), OutEdges: make([]*tgraph.Vertex, 0)}
	//v4 := tgraph.Vertex{Name: "4", InEdges: make([]*tgraph.Vertex, 0), OutEdges: make([]*tgraph.Vertex, 0)}
	//v5 := tgraph.Vertex{Name: "5", InEdges: make([]*tgraph.Vertex, 0), OutEdges: make([]*tgraph.Vertex, 0)}
	//v6 := tgraph.Vertex{Name: "6", InEdges: make([]*tgraph.Vertex, 0), OutEdges: make([]*tgraph.Vertex, 0)}
	//v7 := tgraph.Vertex{Name: "7", InEdges: make([]*tgraph.Vertex, 0), OutEdges: make([]*tgraph.Vertex, 0)}
	//v8 := tgraph.Vertex{Name: "8", InEdges: make([]*tgraph.Vertex, 0), OutEdges: make([]*tgraph.Vertex, 0)}

	tgraph.AddEdge(&v1, &v2)
	tgraph.AddEdge(&v1, &v3)
	//tgraph.AddEdge(&v1, &v4)
	//tgraph.AddEdge(&v4, &v3)
	//tgraph.AddEdge(&v5, &v4)
	//tgraph.AddEdge(&v6, &v4)
	//tgraph.AddEdge(&v6, &v7)
	//tgraph.AddEdge(&v7, &v8)

	steps := make([]*tgraph.Step, 0)
	steps = append(steps, &step1)
	steps = append(steps, &step2)
	//steps = append(steps, &step3)

	newSteps := make([]*tgraph.Step, 0)
	newSteps = append(newSteps, &newStep1)
	newSteps = append(newSteps, &newStep2)
	newSteps = append(newSteps, &newStep3)
	newSteps = append(newSteps, &newStep4)

	vertices := make([]*tgraph.Vertex, 0)
	vertices = append(vertices, &v1)
	vertices = append(vertices, &v2)
	vertices = append(vertices, &v3)
	//vertices = append(vertices, &v4)
	//vertices = append(vertices, &v5)
	//vertices = append(vertices, &v6)
	//vertices = append(vertices, &v7)
	//vertices = append(vertices, &v8)

	minRandInt := 100
	maxRandInt := 600

	rand.Seed(2)

	outputSolutions := tgraph.GetPatternSolutions(vertices, steps)
	nonOverlap := tgraph.GetNonOverlappingSolutions(vertices, steps)

	outputSolutions.PrintSolutions()
	nonOverlap.PrintSolutions()

	fmt.Printf("NumSteps: %d\n", tgraph.GetNumberOfSteps(steps))
	ui.InitUI(1920/2+1920/4, 1080/2+1080/4)

	ui.CurGame.Graph = &tgraph.Graph{Vertices: vertices, Pinned: nil}
	ui.CurGame.Steps = steps
	ui.CurGame.NewSteps = newSteps
	ui.CurGame.Graph.RandomizeVerticesLocs(minRandInt, maxRandInt)

	ui.StartUI()

	fmt.Println("EndMain")

}

/*

	step1 := tgraph.Step{1, 2}
	step2 := tgraph.Step{3, 2}

	v1 := tgraph.Vertex{"1", make([]*tgraph.Vertex, 0), make([]*tgraph.Vertex, 0)}
	v2 := tgraph.Vertex{"2", make([]*tgraph.Vertex, 0), make([]*tgraph.Vertex, 0)}
	v3 := tgraph.Vertex{"3", make([]*tgraph.Vertex, 0), make([]*tgraph.Vertex, 0)}

	AddEdge(&v1, &v2)
	AddEdge(&v3, &v2)

	steps := make([]*tgraph.Step, 0)
	steps = append(steps, &step1)
	steps = append(steps, &step2)

	vertices := make([]*tgraph.Vertex, 0)
	vertices = append(vertices, &v1)
	vertices = append(vertices, &v2)
	vertices = append(vertices, &v3)


*/
