package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"

	"github.com/hajimehoshi/ebiten/v2"
	"kr.rojero.ebitengo/game"
)

var (
	memProfile  = flag.String("memprofile", "", "write memory profile to file")
	traceOut    = flag.String("trace", "", "write trace to file")
	transparent = flag.Bool("transparent", false, "background transparency")
)

func main() {
	flag.Parse()

	if *traceOut != "" {
		f, err := os.Create(*traceOut)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		trace.Start(f)
		defer trace.Stop()
	}

	g, err := game.NewGame()
	if err != nil {
		panic(err)
	}

	if *transparent {
		ebiten.SetScreenTransparent(true)
		ebiten.SetWindowDecorated(false)
	}

	const scale = 2
	ebiten.SetWindowSize(game.ScreenWidth*scale, game.ScreenHeight*scale)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := pprof.WriteHeapProfile(f); err != nil {
			panic(fmt.Sprintf("could not write memory profile: %s", err))
		}
	}
}
