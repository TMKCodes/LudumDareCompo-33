package main

import (
	"math/rand"
	sf "bitbucket.org/krepa098/gosfml2"
)

type Player struct {
	Shape *sf.RectangleShape
	WorldWidth int
	WorldHeight int
}

func NewPlayer(WorldWidth int, WorldHeight int, Textures []Texture) *Player {
	Player := new(Player)
	Player.WorldWidth = WorldWidth
	Player.WorldHeight = WorldHeight
	Player.Shape, _ = sf.NewRectangleShape()
	Player.Shape.SetSize(sf.Vector2f{32, 32})
	Player.Shape.SetTexture((GetTexture("Player.png", Textures)).Data, false)
	//Player.Shape.SetFillColor(sf.Color{255,128,0,255})
	x := rand.Intn(WorldWidth)
	if x > 32 {
		x -= 32
	}
	Player.Shape.SetPosition(sf.Vector2f{float32(x), float32(WorldHeight-64)})
	return Player
}

func (this *Player) MoveLeft() {
	PlayerPosition := this.Shape.GetPosition()
	if PlayerPosition.X > 0 {
		this.Shape.Move(sf.Vector2f{-10.0, 0.0})
	}
}

func (this *Player) MoveUp() {
	PlayerPosition := this.Shape.GetPosition()
	if PlayerPosition.Y > 0 {
		this.Shape.Move(sf.Vector2f{0.0, -10.0})
	}
}

func (this *Player) MoveRight() {
	PlayerPosition := this.Shape.GetPosition()
	if PlayerPosition.X < float32(this.WorldWidth - 32) {
		this.Shape.Move(sf.Vector2f{10.0, 0.0})
	}
}

func (this *Player) MoveDown() {
	PlayerPosition := this.Shape.GetPosition()
	if PlayerPosition.Y < float32(this.WorldHeight - 32) {
		this.Shape.Move(sf.Vector2f{0.0, 10.0})
	}
}

func (this *Player) Draw(RenderWindow *sf.RenderWindow, RenderStates sf.RenderStates) {
	RenderWindow.Draw(this.Shape, RenderStates)
}
