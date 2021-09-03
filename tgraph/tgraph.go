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

type Edge struct {
	V1 *Vertex
	V2 *Vertex
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

func GetNonOverlappingSolutions(vertices []*Vertex, steps []*Step) *Solutions {
	sols := GetPatternSolutions(vertices, steps)

	realSols := &Solutions{Solns: make([]*Solution, 0)}

	usedEdges := make([]*Edge, 0)

	for _, sol := range sols.Solns {
		edgeExists := false
		curUsedEdges := make([]*Edge, len(usedEdges))
		copy(curUsedEdges, usedEdges)
		for _, step := range steps {
			v1 := sol.Soln[step.A-1]
			v2 := sol.Soln[step.B-1]

			if doesEdgeExist(v1, v2, curUsedEdges) {
				edgeExists = true
				break
			}
			curUsedEdges = append(curUsedEdges, &Edge{V1: v1, V2: v2})
		}
		if edgeExists {
			continue
		}
		realSols.Solns = append(realSols.Solns, sol)
		usedEdges = curUsedEdges
	}
	return realSols
}

func doesEdgeExist(v1, v2 *Vertex, edges []*Edge) bool {
	for _, e := range edges {
		if e.V1 == v1 && e.V2 == v2 {
			return true
		}
	}
	return false
}

func UpdateFromPattern(vertices []*Vertex, steps, newSteps []*Step) []*Vertex {
	sols := GetNonOverlappingSolutions(vertices, steps)

	//remove solutions that use the same edges

	fmt.Println("UpdateFromPattern")
	for _, sol := range sols.Solns {
		for _, step := range steps {
			v1 := sol.Soln[step.A-1]
			v2 := sol.Soln[step.B-1]
			RemoveEdge(v1, v2)
		}
	}

	for _, sol := range sols.Solns {
		numSteps := GetNumberOfSteps(steps)

		tmpSol := make([]*Vertex, len(sol.Soln))

		copy(tmpSol, sol.Soln)

		for _, step := range newSteps {
			A := step.A
			B := step.B
			//Does not check if A == B

			if A > numSteps {
				//New Node
				fmt.Println("New Node")
				v1 := &Vertex{Name: fmt.Sprintf("%d", len(vertices)), InEdges: make([]*Vertex, 0), OutEdges: make([]*Vertex, 0)}
				v2 := tmpSol[step.B-1]

				v1.Loc = &util.Vec2{X: v2.Loc.X, Y: v2.Loc.Y}
				AddEdge(v1, v2)

				vertices = append(vertices, v1)

				tmpSol = append(tmpSol, v1)

				numSteps++
			} else if B > numSteps {
				//New Node
				fmt.Println("New Node")
				v1 := tmpSol[step.A-1]
				v2 := &Vertex{Name: fmt.Sprintf("%d", len(vertices)), InEdges: make([]*Vertex, 0), OutEdges: make([]*Vertex, 0)}

				v2.Loc = &util.Vec2{X: v1.Loc.X, Y: v1.Loc.Y}
				AddEdge(v1, v2)

				vertices = append(vertices, v2)

				tmpSol = append(tmpSol, v2)

				numSteps++
			} else {

				v1 := tmpSol[step.A-1]
				v2 := tmpSol[step.B-1]
				AddEdge(v1, v2)

			}
		}
	}
	return vertices
}

func RemoveEdge(v1, v2 *Vertex) {
	outEdges := v1.OutEdges
	fmt.Println("RemoveEdge")
	i := 0
	for _, v := range v1.OutEdges {
		if v == v2 {
			//Remove
		} else {
			outEdges[i] = v
			i++
		}
	}
	v1.OutEdges = outEdges[:i]

	inEdges := v2.InEdges
	i = 0
	for _, v := range v2.InEdges {
		if v == v1 {
			//Remove
		} else {
			inEdges[i] = v
			i++
		}
	}
	v2.InEdges = inEdges[:i]
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
