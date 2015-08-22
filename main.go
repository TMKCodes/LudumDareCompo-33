package main

import (
	"runtime"
)

func main() {
	runtime.LockOSThread()
	Game := NewGame("Child Service", 1920, 1080, 32, true);
	for Game.IsOpen() {
		Game.Update()
		Game.Draw()
	}
}