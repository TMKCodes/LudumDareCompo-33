package main

import (
	"time"
	"math/rand"
	sf "bitbucket.org/krepa098/gosfml2"
)

type Enemy struct {
	HasChild bool
	HasFather bool
	HasMother bool
	FatherHasWeapon bool
	Shape *sf.RectangleShape
}


func NewEnemy(WorldWidth int) *Enemy {
	Enemy := new(Enemy)
	rand.Seed(time.Now().UnixNano())
	Enemy.HasChild = true
	if rand.Intn(100) < 80 {
		Enemy.HasMother = true
	} else {
		Enemy.HasMother = false
	}
	if Enemy.HasMother == true {
		if rand.Intn(1) == 1 {
			Enemy.HasFather = true
		} else {
			Enemy.HasFather = false
		}	
	} else {
		Enemy.HasFather = true
	}
	if rand.Intn(1) == 1 {
		Enemy.FatherHasWeapon = true
	} else {
		Enemy.FatherHasWeapon = false
	}
	Enemy.Shape, _ = sf.NewRectangleShape()
	if Enemy.HasMother == true && Enemy.HasFather == true && Enemy.HasChild == true {
		Enemy.Shape.SetSize(sf.Vector2f{96, 32})	
		Enemy.Shape.SetFillColor(sf.Color{128,128,128,255})
	} else if Enemy.HasMother == true && Enemy.HasFather == false && Enemy.HasChild == true {
		Enemy.Shape.SetSize(sf.Vector2f{64, 32})
		Enemy.Shape.SetFillColor(sf.Color{128,0,128,255})
	} else if Enemy.HasMother == false && Enemy.HasFather == true && Enemy.HasFather == true {
		Enemy.Shape.SetSize(sf.Vector2f{64, 32})
		Enemy.Shape.SetFillColor(sf.Color{0,128,128,255})
	}
	x := rand.Intn(WorldWidth);
	if x > 96 {
		x -= 96
	}
	Enemy.Shape.SetPosition(sf.Vector2f{float32(x), 0.0})
	return Enemy
}

func (this *Enemy) MoveDown() {
	this.Shape.Move(sf.Vector2f{0.0, 5.0})
}

func (this *Enemy) Destroy(WorldHeight int) bool {
	position := this.Shape.GetPosition()
	if position.Y > float32(WorldHeight) {
		return true
	}
	return false
}

func (this *Enemy) Draw(RenderWindow *sf.RenderWindow, RenderStates sf.RenderStates) {
	RenderWindow.Draw(this.Shape, RenderStates)
}