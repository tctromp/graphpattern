package tgraph

import (
	"fmt"

	"github.com/tctromp/graphpattern/util"
)

type Vertex struct {
	Name     string
	Loc      *util.Vec2
	InEdges  []*Vertex // Directed edges into vertex
	OutEdges []*Vertex // Directed edges out of vertex
}

func AddEdge(v1, v2 *Vertex) {
	v1.OutEdges = append(v1.OutEdges, v2)
	v2.InEdges = append(v2.InEdges, v1)
}

type Step struct {
	A int // Ex: {A, B}
	B int
}

func GetNumberOfSteps(steps []*Step) int {
	max := 0
	for _, step := range steps {
		if step.A > max {
			max = step.A
		}
		if step.B > max {
			max = step.B
		}
	}
	return max
}

type Solution struct {
	Soln []*Vertex
}

type Solutions struct {
	Solns []*Solution
}

func (solns *Solutions) AddSolution(soln *Solution) {
	solns.Solns = append(solns.Solns, soln)
}

func (solns *Solutions) PrintSolutions() {
	fmt.Printf("Found: %d possible solution\n", len(solns.Solns))
	for i, o := range solns.Solns {
		fmt.Printf("%d: ", i)
		for i, v := range o.Soln {
			fmt.Print("[" + v.Name + "]")
			if i-1 < len(solns.Solns) {
				fmt.Print(" + ")
			}
		}
		fmt.Println()
	}
}

//Note: assuming steps are sorted, e.g {{1,2}, {1,3}, {3,2}}
func GetPatternSolutions(vertices []*Vertex, steps []*Step) *Solutions {
	soln := make([]*Vertex, GetNumberOfSteps(steps)) // Number of unique vertices
	outputSolutions := &Solutions{make([]*Solution, 0)}

	for _, v := range vertices {
		soln[0] = v
		recur(0, 1, vertices, soln, steps, outputSolutions)
	}

	return outputSolutions
}

func recur(step, size int, vertices []*Vertex, soln []*Vertex, steps []*Step, outputSolutions *Solutions) {

	if step == len(steps) {
		newAray := make([]*Vertex, len(soln))
		copy(newAray, soln)

		outputSolutions.AddSolution(&Solution{newAray})
		return
	}

	A := steps[step].A
	B := steps[step].B

	if A <= size && B <= size {
		//No new nodes added, just checking for edges
		if A > B {
			if !soln[B-1].IsOutEdge(soln[A-1]) {
				return
			}
			recur(step+1, size, vertices, soln, steps, outputSolutions)
		} else {
			if !soln[A-1].IsOutEdge(soln[B-1]) {
				return
			}
			recur(step+1, size, vertices, soln, steps, outputSolutions)
		}
	} else {
		//New vertex needs to be added, loop through posibilities
		for _, v := range soln[A-1].OutEdges {
			if v.AlreadyUsed(size, soln) {
				continue
			}
			//tmp := soln[B-1]
			soln[B-1] = v
			recur(step+1, size+1, vertices, soln, steps, outputSolutions)
			//soln[B-1] = tmp
		}
	}

}

func (v1 *Vertex) AlreadyUsed(size int, soln []*Vertex) bool {
	for i := 0; i < size; i++ {
		if soln[i] == v1 {
			return true
		}
	}
	return false
}

func (v1 *Vertex) IsOutEdge(b *Vertex) bool { // Then a IS an outedge of b (a -> b)
	for _, v2 := range b.OutEdges {
		if v1 == v2 {
			return true
		}
	}
	return false
}
