package ui

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tctromp/graphpattern/tgraph"
)

//https://flatuicolors.com/palette/defo

//rgb(44, 62, 80)
var backgroundColor = color.RGBA{R: 44, G: 62, B: 80, A: 255}

//rgb(189, 195, 199)
var edgeColor = color.RGBA{R: 186, G: 195, B: 199, A: 255}

//rgb(231, 76, 60)
var vertexColor = color.RGBA{R: 231, G: 76, B: 60, A: 255}

//rgb(46, 204, 113)
var pinnedVertexColor = color.RGBA{R: 46, G: 204, B: 113, A: 255}

var CurGame *Game
var vertexSize = 10.0
var springLength = 100.0
var totalSteps = 1
var scale = 1.0
var scaleSteps = 1.0

type Game struct {
	Graph    *tgraph.Graph
	Steps    []*tgraph.Step
	NewSteps []*tgraph.Step
}

func (g *Game) Update() error {
	CurGame.Graph.UpdateMouseEvents(vertexSize, scale)

	for i := 0; i < 10; i++ {
		CurGame.Graph.SpringUpdate(springLength)

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		CurGame.Graph.RandomizeVerticesLocs(100, 600)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		CurGame.Graph.Vertices = tgraph.UpdateFromPattern(CurGame.Graph.Vertices, CurGame.Steps, CurGame.NewSteps)
		CurGame.Graph.RandomizeVerticesLocs(100, 600)
		totalSteps++
	}

	_, y := ebiten.Wheel()
	if y == -1 || inpututil.IsKeyJustPressed(ebiten.KeyNumpad8) {
		//scale -= 0.1
		scaleSteps += 0.25
	} else if y == 1 || inpututil.IsKeyJustPressed(ebiten.KeyNumpad2) {
		//scale += 0.1
		scaleSteps -= 0.25

	}
	scale = math.Pow(scaleSteps, 2)
	if scaleSteps < 0 {
		scale = 1 / scale
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Steps: %d\nVertices: %d", totalSteps, len(CurGame.Graph.Vertices)))

	if CurGame.Graph != nil {
		CurGame.Graph.DrawGraph(screen, scale, vertexSize, edgeColor, vertexColor, pinnedVertexColor)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func InitUI(width, height int) {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Graph")
	CurGame = &Game{Graph: nil}
}

func StartUI() {

	if err := ebiten.RunGame(CurGame); err != nil {
		log.Fatal(err)
	}
}
