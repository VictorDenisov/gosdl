package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	rect1 := &sdl.Rect{288, 208, 100, 100}
	rect2 := &sdl.Rect{50, 50, 100, 80}
	rects := []*sdl.Rect{rect1, rect2}
	var selectedRect *sdl.Rect
	var leftMouseButtonDown bool
	var mousePos sdl.Point
	var clickOffset sdl.Point

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			sdl.Delay(10)
			switch e := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.MouseMotionEvent:
				mousePos = sdl.Point{e.X, e.Y}
				if leftMouseButtonDown && selectedRect != nil {
					selectedRect.X = mousePos.X - clickOffset.X
					selectedRect.Y = mousePos.Y - clickOffset.Y
				}
			case *sdl.MouseButtonEvent:
				if e.Type == sdl.MOUSEBUTTONUP {
					if leftMouseButtonDown && e.Button == sdl.BUTTON_LEFT {
						leftMouseButtonDown = false
						selectedRect = nil
					}
				} else if e.Type == sdl.MOUSEBUTTONDOWN {
					if !leftMouseButtonDown && e.Button == sdl.BUTTON_LEFT {
						leftMouseButtonDown = true
						for _, r := range rects {
							if mousePos.InRect(r) {
								selectedRect = r
								clickOffset.X = mousePos.X - r.X
								clickOffset.Y = mousePos.Y - r.Y
								break
							}
						}
					}
				}
			}
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.Clear()

			for _, r := range rects {
				if r == selectedRect {
					renderer.SetDrawColor(0, 0, 255, 255)
				} else {
					renderer.SetDrawColor(0, 255, 0, 255)
				}
				renderer.FillRect(r)
			}
			for i := int32(0); i < 100; i++ {
				r := &sdl.Rect{300, i * 10, 9, 9}
				renderer.SetDrawColor(255-uint8(i*2), 255, 255-uint8(i*2), 255)
				renderer.FillRect(r)
			}
			renderer.Present()
		}
	}
	sdl.Quit()
}
