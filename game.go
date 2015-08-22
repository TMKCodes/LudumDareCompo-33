package main

import (
	"time"
	sf "bitbucket.org/krepa098/gosfml2"
)

type Game struct {
	RenderWindow *sf.RenderWindow
	EnemyList []*Enemy
	EnemyGenerationSpeed int64
	LastEnemyGenerated int64
	Player *Player
	Width uint
	Height uint
}

func NewGame(title string, width uint, height uint, bpp uint, vsync bool) *Game {
	Game := new(Game)
	Game.Width = width
	Game.Height = height
	Game.RenderWindow = sf.NewRenderWindow(sf.VideoMode{width, height, bpp}, title, sf.StyleDefault, sf.DefaultContextSettings())
	Game.RenderWindow.SetVSyncEnabled(vsync)
	Game.Player = NewPlayer(int(width), int(height))
	Game.EnemyList = append(Game.EnemyList, NewEnemy(int(width)))
	Game.LastEnemyGenerated = time.Now().UnixNano()
	Game.EnemyGenerationSpeed = 500000000
	return Game
}

func (this *Game) IsOpen() bool {
	return this.RenderWindow.IsOpen()
}

func (this *Game) Update() {
	t := time.Now().UnixNano()
	if t > this.LastEnemyGenerated + this.EnemyGenerationSpeed {
		this.LastEnemyGenerated = t 
		this.EnemyList = append(this.EnemyList, NewEnemy(int(this.Width)))
	}
	for event := this.RenderWindow.PollEvent(); event != nil; event = this.RenderWindow.PollEvent() {
		switch ev := event.(type) {
		case sf.EventKeyReleased:
			if ev.Code == sf.KeyEscape {
				this.RenderWindow.Close()
			}
		case sf.EventClosed:
			this.RenderWindow.Close()
		}
	}
	if sf.KeyboardIsKeyPressed(sf.KeyLeft) == true {
		this.Player.MoveLeft()
	}
	if sf.KeyboardIsKeyPressed(sf.KeyUp) == true {
		this.Player.MoveUp()
	}
	if sf.KeyboardIsKeyPressed(sf.KeyRight) == true {
		this.Player.MoveRight()
	}
	if sf.KeyboardIsKeyPressed(sf.KeyDown) == true {
		this.Player.MoveDown()
	}
	for Enemy, _ := range this.EnemyList {
		this.EnemyList[Enemy].MoveDown()
		if this.EnemyList[Enemy].Destroy(int(this.Height)) == true {
			this.EnemyList = append(this.EnemyList[:Enemy], this.EnemyList[Enemy+1:]...)
			break;
		}
	}
}

func (this *Game) Draw() {
	this.RenderWindow.Clear(sf.Color{255, 255, 255, 255})
	for Enemy, _ := range this.EnemyList {
		this.EnemyList[Enemy].Draw(this.RenderWindow, sf.DefaultRenderStates())
	}
	this.Player.Draw(this.RenderWindow, sf.DefaultRenderStates())
	this.RenderWindow.Display()
}

