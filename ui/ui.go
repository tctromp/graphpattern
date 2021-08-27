package ui

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
var vertexSize = 20.0
var springLength = 50.0

type Game struct {
	Graph *tgraph.Graph
}

func (g *Game) Update() error {
	CurGame.Graph.UpdateMouseEvents(vertexSize)
	CurGame.Graph.SpringUpdate(springLength)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World!")
	if CurGame.Graph != nil {
		CurGame.Graph.DrawGraph(screen, vertexSize, edgeColor, vertexColor, pinnedVertexColor)
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
