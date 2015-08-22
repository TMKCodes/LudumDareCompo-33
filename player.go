package main

import (
	"math/rand"
	sf "bitbucket.org/krepa098/gosfml2"
)

type Player struct {
	Shape *sf.RectangleShape
}

func NewPlayer(WorldWidth int, WorldHeight int) *Player {
	Player := new(Player)
	Player.Shape, _ = sf.NewRectangleShape()
	Player.Shape.SetSize(sf.Vector2f{32, 32})
	Player.Shape.SetFillColor(sf.Color{255,128,0,255})
	x := rand.Intn(WorldWidth)
	if x > 32 {
		x -= 32
	}
	Player.Shape.SetPosition(sf.Vector2f{float32(x), float32(WorldHeight-32)})
	return Player
}

func (this *Player) MoveLeft() {
	this.Shape.Move(sf.Vector2f{-10.0, 0.0})
}

func (this *Player) MoveUp() {
	this.Shape.Move(sf.Vector2f{0.0, -10.0})
}

func (this *Player) MoveRight() {
	this.Shape.Move(sf.Vector2f{10.0, 0.0})
}

func (this *Player) MoveDown() {
	this.Shape.Move(sf.Vector2f{0.0, 10.0})
}

func (this *Player) Draw(RenderWindow *sf.RenderWindow, RenderStates sf.RenderStates) {
	RenderWindow.Draw(this.Shape, RenderStates)
}
