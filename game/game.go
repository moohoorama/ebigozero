package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 320
	ScreenHeight = 240
)

type Game struct {
	i int
}

func NewGame() (*Game, error) {
	return &Game{}, nil
}

func (g *Game) Update() error {
	g.i++
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	// draw red rectangle
	screen.Fill(color.RGBA{0x00, 0x80, 0x80, 0xff})
	for j := range 50 {
		idx := g.i + j*10
		R := uint8(idx % 255)
		G := uint8(idx % 250)
		B := uint8(idx % 245)
		A := uint8(idx%50) + 50
		X := (idx + j) % ScreenWidth
		Y := idx % ScreenHeight
		screen.Set(X, Y, color.RGBA{R, G, B, A})
	}
	ebitenutil.DebugPrint(screen, "Hello, Ebiten-Mobile!")
	// draw text

}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// 가상의 논리 해상도 설정
	return 320, 240
}
