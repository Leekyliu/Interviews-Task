package main

import (
	"fmt"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle = "The Snake Game"

	windowHeight = 600
	windowWidth  = 800
)

func main() {
	run()
}

func run() error {



	window, err := sdl.CreateWindow(
		windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN,
	)

	scene, err := newScene(window)
	if err != nil {
		return fmt.Errorf("could not create scene: %v", err)
	}

	events := make(chan sdl.Event)
	errc := scene.run(events)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}


}
