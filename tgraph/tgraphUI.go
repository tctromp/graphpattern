package tgraph

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tctromp/graphpattern/util"
)

type Graph struct {
	Vertices []*Vertex
	Pinned   *Vertex
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

func (graph *Graph) DrawGraph(screen *ebiten.Image, vertexSize float64, edgeColor, vertexColor, pinnedVertexColor color.Color) {
	vertices := graph.Vertices

	for _, v := range vertices {
		for _, e := range v.OutEdges {
			ebitenutil.DrawLine(screen, offsetX+v.Loc.X, offsetY+v.Loc.Y, offsetX+e.Loc.X, offsetY+e.Loc.Y, edgeColor)
		}
	}
	for _, v := range vertices {
		clr := vertexColor
		if v == graph.Pinned {
			clr = pinnedVertexColor
		}
		ebitenutil.DrawRect(screen, offsetX+v.Loc.X-vertexSize/2, offsetY+v.Loc.Y-vertexSize/2, vertexSize, vertexSize, clr)

	}

}

var clicked *Vertex
var lastX float64
var lastY float64
var offsetX = 0.0
var offsetY = 0.0

func (g *Graph) UpdateMouseEvents(vertexSize float64) {
	x, y := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.GraphClick(float64(x), float64(y), vertexSize)
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		g.GraphRelease(float64(x), float64(y))
	} else {
		g.UpdateMousePos(float64(x), float64(y))
	}
}

func (g *Graph) GraphClick(x, y, vertexSize float64) {
	x = x - offsetX
	y = y - offsetY
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

func (g *Graph) UpdateMousePos(x, y float64) {
	if clicked != nil {
		clicked.Loc = clicked.Loc.Add(&util.Vec2{X: x - offsetX - clicked.Loc.X, Y: y - offsetY - clicked.Loc.Y})
	} else {
		//Drag whole screen
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			offsetX += x - lastX
			offsetY += y - lastY
		}
	}
	lastX = x
	lastY = y
}

func (g *Graph) SpringUpdate(length float64) {
	forces1 := make([]*util.Vec2, len(g.Vertices))
	//forces2 := make([]*util.Vec2, len(g.Edges))
	for i, _ := range g.Vertices {
		forces1[i] = &util.Vec2{X: 0, Y: 0}
	}

	for i, v1 := range g.Vertices {
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
	}

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
		v.Loc = v.Loc.Add(forces1[i].Mul(0.05))
	}

}

func fRep(v1, v2 *Vertex, length float64) *util.Vec2 {
	return v2.Loc.Norm(v1.Loc).Mul(length * length / v1.Loc.Dist(v2.Loc))
}

func fAttr(v1, v2 *Vertex, length float64) *util.Vec2 {
	return v1.Loc.Norm(v2.Loc).Mul(math.Pow(v1.Loc.Dist(v2.Loc), 2) / length)
}
