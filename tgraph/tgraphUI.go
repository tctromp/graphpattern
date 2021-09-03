package tgraph

import (
	"image/color"
	"math/rand"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tctromp/graphpattern/util"
)

type Graph struct {
	Vertices []*Vertex
	Pinned   *Vertex
}

func (g *Graph) RandomizeVerticesLocs(minRandInt, maxRandInt int) {
	vertices := g.Vertices
	for _, v := range vertices {
		randX := rand.Intn(maxRandInt-minRandInt) + minRandInt
		randY := rand.Intn(maxRandInt-minRandInt) + minRandInt
		v.Loc = &util.Vec2{X: float64(randX), Y: float64(randY)}
	}
}

//var vertexFont font.Face

/*
func setFont() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72

	curFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	vertexFont = curFont
}
*/

func getFont() {

}

func (graph *Graph) DrawGraph(screen *ebiten.Image, scale float64, vertexSize float64, edgeColor, vertexColor, pinnedVertexColor color.Color) {
	vertices := graph.Vertices

	for _, v := range vertices {
		for _, e := range v.OutEdges {
			ebitenutil.DrawLine(screen, (offsetX+v.Loc.X)/scale, (offsetY+v.Loc.Y)/scale, (offsetX+e.Loc.X)/scale, (offsetY+e.Loc.Y)/scale, edgeColor)
		}
	}

	/*
		for _, v := range vertices {
			clr := vertexColor
			if v == graph.Pinned {
				clr = pinnedVertexColor
			}
			ebitenutil.DrawRect(screen, (offsetX+v.Loc.X-vertexSize/2)/scale, (offsetY+v.Loc.Y-vertexSize/2)/scale, vertexSize, vertexSize, clr)

		}
	*/

}

var clicked *Vertex
var lastX float64
var lastY float64
var offsetX = 0.0
var offsetY = 0.0

func (g *Graph) UpdateMouseEvents(vertexSize, scale float64) {
	x, y := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.GraphClick(float64(x), float64(y), vertexSize, scale)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.GraphRelease(float64(x), float64(y))
	} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.UpdatePinned(x, y, vertexSize)
	} else {
		g.UpdateMousePos(float64(x), float64(y), scale)
	}
}

func (g *Graph) UpdatePinned(x, y int, vertexSize float64) {
	xx := float64(x) - offsetX
	yy := float64(y) - offsetY

	for _, v := range g.Vertices {
		if xx >= v.Loc.X-vertexSize && xx < v.Loc.X+vertexSize/2 {
			if yy >= v.Loc.Y-vertexSize/2 && yy < v.Loc.Y+vertexSize/2 {
				g.Pinned = v
				break
			}
		}
	}
}

func (g *Graph) GraphClick(x, y, vertexSize, scale float64) {
	x = (x*scale - offsetX)
	y = (y*scale - offsetY)
	for _, v := range g.Vertices {
		if x >= v.Loc.X-vertexSize && x < v.Loc.X+vertexSize/2 {
			if y >= v.Loc.Y-vertexSize/2 && y < v.Loc.Y+vertexSize/2 {
				clicked = v
				break
			}
		}
	}
}

func (g *Graph) GraphRelease(x, y float64) {
	clicked = nil
}

func (g *Graph) UpdateMousePos(x, y, scale float64) {
	if clicked != nil {
		clicked.Loc = clicked.Loc.Add(&util.Vec2{X: (x - lastX) * scale, Y: (y - lastY) * scale})
	} else {
		//Drag whole screen
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			offsetX += (x - lastX) * scale
			offsetY += (y - lastY) * scale
		}
	}
	lastX = x
	lastY = y
}

func (g *Graph) Compute(v1 *Vertex, i int, forces *[]*util.Vec2, length float64, wg *sync.WaitGroup) {
	forces1 := *forces
	for _, v2 := range g.Vertices {
		if v1 == v2 {
			continue
		}
		*forces1[i] = *forces1[i].Add(fRep(v1, v2, length))
	}
	for _, v2 := range v1.OutEdges {
		forces1[i] = forces1[i].Add(fAttr(v1, v2, length))
	}
	for _, v2 := range v1.InEdges {
		forces1[i] = forces1[i].Add(fAttr(v1, v2, length))
	}
	wg.Done()
}

func (g *Graph) SpringUpdate(length float64) {
	forces1 := make([]*util.Vec2, len(g.Vertices))
	//forces2 := make([]*util.Vec2, len(g.Edges))
	for i, _ := range g.Vertices {
		forces1[i] = &util.Vec2{X: 0, Y: 0}
	}

	var wg sync.WaitGroup

	for i, v1 := range g.Vertices {
		wg.Add(1)
		go g.Compute(v1, i, &forces1, length, &wg)
		/*
			for _, v2 := range g.Vertices {
				if v1 == v2 {
					continue
				}
				forces1[i] = forces1[i].Add(fRep(v1, v2, length))
			}
			for _, v2 := range v1.OutEdges {
				forces1[i] = forces1[i].Add(fAttr(v1, v2, length))
			}
			for _, v2 := range v1.InEdges {
				forces1[i] = forces1[i].Add(fAttr(v1, v2, length))
			}
		*/
	}

	wg.Wait()

	/*
		for i, e := range g.Edges {
			forces1[i] = forces1[i].Add(fAttr(e.V1, e.V2, length))
			//forces2[i] = fRep(e.V2, e.V1, length).Add(fAttr(e.V2, e.V1, length))
		}


		for i, e := range g.Edges {
			e.V1.Loc = e.V1.Loc.Add(forces1[i].Mul(0.01))
			//e.V2.Loc = e.V2.Loc.Add(forces2[i].Mul(0.01))
		}
	*/

	for i, v := range g.Vertices {
		if g.Pinned == v {
			continue
		}
		v.Loc = v.Loc.Add(forces1[i].Mul(0.01))
	}

}

func fRep(v1, v2 *Vertex, length float64) *util.Vec2 {
	return v2.Loc.Norm(v1.Loc).Mul(length * length / v1.Loc.Dist(v2.Loc))
}

func fAttr(v1, v2 *Vertex, length float64) *util.Vec2 {
	dist2 := v1.Loc.Dist2(v2.Loc)
	return v1.Loc.Norm(v2.Loc).Mul(dist2 / length)
}
