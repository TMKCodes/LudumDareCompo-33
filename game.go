package main

import (
	"fmt"
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
	ChildsCollected int
	ChildsCollectedText *sf.Text
	WelcomeText *sf.Text
	PlayText *sf.Text
	StartTime int64
	EndTime int64
	TimeText *sf.Text
	EndingText *sf.Text
	EndingPointsText *sf.Text
	RestartText *sf.Text
	CreditsText *sf.Text
	State int
	Font *sf.Font
	Music *sf.Music
	Textures []Texture
	ChildStolen *sf.Music
}

func NewGame(title string, width uint, height uint, bpp uint, vsync bool) *Game {
	Game := new(Game)
	Game.Width = width
	Game.Height = height
	Game.RenderWindow = sf.NewRenderWindow(sf.VideoMode{width, height, bpp}, title, sf.StyleDefault, sf.DefaultContextSettings())
	Game.RenderWindow.SetVSyncEnabled(vsync)
	Game.Font, _ = sf.NewFontFromFile("res/fonts/UbuntuMono-R.ttf")
	Game.State = 0;
	Game.StartStateInit()
	Game.Textures = NewTextures("res/images/")
	return Game
}

func (this *Game) IsOpen() bool {
	return this.RenderWindow.IsOpen()
}

func (this *Game) GameStateInit() {
	this.Player = NewPlayer(int(this.Width), int(this.Height), this.Textures)
	this.EnemyList = make([]*Enemy, 0)
	this.EnemyList = append(this.EnemyList, NewEnemy(int(this.Width), this.Textures))
	this.LastEnemyGenerated = time.Now().UnixNano()
	this.EnemyGenerationSpeed = 500000000
	this.ChildsCollected = 0
	this.ChildsCollectedText, _ = sf.NewText(this.Font);
	this.ChildsCollectedText.SetString(fmt.Sprintf("Childs Collected: %v", this.ChildsCollected))
	this.ChildsCollectedText.SetCharacterSize(24);
	this.ChildsCollectedText.SetColor(sf.Color{0,0,0,255})
	this.ChildsCollectedText.SetPosition(sf.Vector2f{20, 20})
	this.StartTime = time.Now().Unix()
	this.TimeText, _ = sf.NewText(this.Font);
	this.TimeText.SetString(fmt.Sprintf("Collection time: %v", time.Now().Unix() - this.StartTime))
	TimeTextRect := this.TimeText.GetLocalBounds()
	this.TimeText.SetOrigin(sf.Vector2f{float32(TimeTextRect.Width), 0.0})
	this.TimeText.SetCharacterSize(24);
	this.TimeText.SetColor(sf.Color{0,0,0,255})
	this.TimeText.SetPosition(sf.Vector2f{float32(int(this.Width) - 20), 20})
	this.Music, _ = sf.NewMusicFromFile("res/audio/Child_Service_-_Going_to_steal_your_Kids.ogg")
	this.Music.SetLoop(true)
	this.Music.Play()
	this.ChildStolen, _ = sf.NewMusicFromFile("res/audio/childstolen.wav")
}

func (this *Game) StartStateInit() {
	this.WelcomeText, _ = sf.NewText(this.Font)
	this.WelcomeText.SetString(fmt.Sprintf("You work now in Child Service and you must collect every child who has no mother!"))
	this.WelcomeText.SetCharacterSize(36)
	this.WelcomeText.SetColor(sf.Color{0,0,0,255})
	WelcomeTextRect := this.WelcomeText.GetLocalBounds()
	x := (float32(this.Width) - WelcomeTextRect.Width) / 2
	this.WelcomeText.SetPosition(sf.Vector2f{x, 250})

	this.PlayText, _ = sf.NewText(this.Font)
	this.PlayText.SetString(fmt.Sprintf("Press Enter to play!"))
	this.PlayText.SetCharacterSize(42)
	this.PlayText.SetColor(sf.Color{0,0,0,255})
	PlayTextRect := this.PlayText.GetLocalBounds()
	x = (float32(this.Width) - PlayTextRect.Width) / 2
	this.PlayText.SetPosition(sf.Vector2f{x, 350})
}

func (this *Game) KilledStateInit() {
	this.EndTime = time.Now().Unix()
	this.EndingText, _ = sf.NewText(this.Font)
	this.EndingText.SetString(fmt.Sprintf("You were killed by a Father who had a weapon."))
	this.EndingText.SetCharacterSize(36)
	this.EndingText.SetColor(sf.Color{0,0,0,255})
	EndingTextRect := this.EndingText.GetLocalBounds()
	x := (float32(this.Width) - EndingTextRect.Width) / 2
	this.EndingText.SetPosition(sf.Vector2f{x, 250})

	this.EndingPointsText, _ = sf.NewText(this.Font)
	this.EndingPointsText.SetString(fmt.Sprintf("You collected %v children in %v seconds.", this.ChildsCollected, this.EndTime - this.StartTime))
	this.EndingPointsText.SetCharacterSize(36)
	this.EndingPointsText.SetColor(sf.Color{0,0,0,255})
	EndingPointsTextRect := this.EndingPointsText.GetLocalBounds()
	x = (float32(this.Width) - EndingPointsTextRect.Width) / 2
	this.EndingPointsText.SetPosition(sf.Vector2f{x, 350})

	this.RestartText, _ = sf.NewText(this.Font)
	this.RestartText.SetString(fmt.Sprintf("Press Enter to restart the game!"))
	this.RestartText.SetCharacterSize(42)
	this.RestartText.SetColor(sf.Color{0,0,0,255})
	RestartTextRect := this.RestartText.GetLocalBounds()
	x = (float32(this.Width) - RestartTextRect.Width) / 2
	this.RestartText.SetPosition(sf.Vector2f{x, 450})

	this.CreditsText, _ = sf.NewText(this.Font)
	this.CreditsText.SetString(fmt.Sprintf("Created for Ludum Dare 33 Compo by Toni Korpela!"))
	this.CreditsText.SetCharacterSize(42)
	this.CreditsText.SetColor(sf.Color{0,0,0,255})
	CreditsTextRect := this.CreditsText.GetLocalBounds()
	x = (float32(this.Width) - CreditsTextRect.Width) / 2
	this.CreditsText.SetPosition(sf.Vector2f{x, 550})
}

func (this *Game) MonsterStateInit() {
	this.EndTime = time.Now().Unix()
	this.EndingText, _ = sf.NewText(this.Font)
	this.EndingText.SetString(fmt.Sprintf("You were such a monster to try and steal child from a mother."))
	this.EndingText.SetCharacterSize(36)
	this.EndingText.SetColor(sf.Color{0,0,0,255})
	EndingTextRect := this.EndingText.GetLocalBounds()
	x := (float32(this.Width) - EndingTextRect.Width) / 2
	this.EndingText.SetPosition(sf.Vector2f{x, 250})

	this.EndingPointsText, _ = sf.NewText(this.Font)
	this.EndingPointsText.SetString(fmt.Sprintf("You collected %v children in %v seconds.", this.ChildsCollected, this.EndTime - this.StartTime))
	this.EndingPointsText.SetCharacterSize(36)
	this.EndingPointsText.SetColor(sf.Color{0,0,0,255})
	EndingPointsTextRect := this.EndingPointsText.GetLocalBounds()
	x = (float32(this.Width) - EndingPointsTextRect.Width) / 2
	this.EndingPointsText.SetPosition(sf.Vector2f{x, 350})
	
	this.RestartText, _ = sf.NewText(this.Font)
	this.RestartText.SetString(fmt.Sprintf("Press Enter to restart the game!"))
	this.RestartText.SetCharacterSize(42)
	this.RestartText.SetColor(sf.Color{255,0,0,255})
	RestartTextRect := this.RestartText.GetLocalBounds()
	x = (float32(this.Width) - RestartTextRect.Width) / 2
	this.RestartText.SetPosition(sf.Vector2f{x, 450})

	this.CreditsText, _ = sf.NewText(this.Font)
	this.CreditsText.SetString(fmt.Sprintf("Created for Ludum Dare 33 Compo by Toni Korpela!"))
	this.CreditsText.SetCharacterSize(42)
	this.CreditsText.SetColor(sf.Color{0,0,0,255})
	CreditsTextRect := this.CreditsText.GetLocalBounds()
	x = (float32(this.Width) - CreditsTextRect.Width) / 2
	this.CreditsText.SetPosition(sf.Vector2f{x, 550})
}

func (this *Game) Update() {
	if this.State == 0 {
		// Opening state
		for event := this.RenderWindow.PollEvent(); event != nil; event = this.RenderWindow.PollEvent() {
			switch ev := event.(type) {
			case sf.EventKeyReleased:
				if ev.Code == sf.KeyEscape {
					this.RenderWindow.Close()
				}
				if ev.Code == sf.KeyReturn {
					this.State = 1
					this.GameStateInit()
				}
			case sf.EventClosed:
				this.RenderWindow.Close()
			}
		}
	} else if this.State == 1 {
		// Game state
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
		t := time.Now().UnixNano()
		this.TimeText.SetString(fmt.Sprintf("Collection time: %v", time.Now().Unix() - this.StartTime))
		if t > this.LastEnemyGenerated + this.EnemyGenerationSpeed {
			this.LastEnemyGenerated = t 
			this.EnemyList = append(this.EnemyList, NewEnemy(int(this.Width), this.Textures))
		}
		if sf.KeyboardIsKeyPressed(sf.KeyLeft) == true || sf.KeyboardIsKeyPressed(sf.KeyA) == true {
			this.Player.MoveLeft()
		}
		if sf.KeyboardIsKeyPressed(sf.KeyUp) == true || sf.KeyboardIsKeyPressed(sf.KeyW) == true  {
			this.Player.MoveUp()
		}
		if sf.KeyboardIsKeyPressed(sf.KeyRight) == true || sf.KeyboardIsKeyPressed(sf.KeyD) == true {
			this.Player.MoveRight()
		}
		if sf.KeyboardIsKeyPressed(sf.KeyDown) == true || sf.KeyboardIsKeyPressed(sf.KeyS) == true {
			this.Player.MoveDown()
		}
		for i := 0; i < len(this.EnemyList); i++  {
			if this.EnemyList[i].Collision(this.Player) == true {
				if this.EnemyList[i].GetHasFather() == true && this.EnemyList[i].GetHasMother() == false {
					if this.EnemyList[i].GetFatherHasWeapon() == false {
						this.ChildStolen.Play()
						this.ChildsCollected += 1;
						this.ChildsCollectedText.SetString(fmt.Sprintf("Childs Collected: %v", this.ChildsCollected))
						this.EnemyList = append(this.EnemyList[:i], this.EnemyList[i+1:]...)
						i--;
					} else if this.EnemyList[i].GetFatherHasWeapon() == true {
						// End the game, because the Child Service worker was killed.
						this.KilledStateInit()
						this.State = 2
					}
				} else {
					// End the game and call the Child Service worker a Monster.
					this.MonsterStateInit()
					this.State = 3
				}
			} else {
				this.EnemyList[i].MoveDown()
				if this.EnemyList[i].Destroy(int(this.Height)) == true {
					this.EnemyList = append(this.EnemyList[:i], this.EnemyList[i+1:]...)
					i--;
				}
			}
		}
		if this.ChildsCollected > 20 && this.ChildsCollected < 40 {
			this.EnemyGenerationSpeed = 500000000
		} else if this.ChildsCollected > 40 && this.ChildsCollected < 60 {
			this.EnemyGenerationSpeed = 400000000
		} else if this.ChildsCollected > 60 && this.ChildsCollected < 80 {
			this.EnemyGenerationSpeed = 300000000
		} else if this.ChildsCollected > 80 && this.ChildsCollected < 100 {
			this.EnemyGenerationSpeed = 200000000
		} else if this.ChildsCollected > 100 && this.ChildsCollected < 150 {
			this.EnemyGenerationSpeed = 100000000
		} else if this.ChildsCollected > 150 && this.ChildsCollected < 300 {
			this.EnemyGenerationSpeed = 50000000
		} else if this.ChildsCollected > 300 {
			this.EnemyGenerationSpeed = 25000000
		}
	} else if this.State == 2 {
		// ending when killed state
		this.Music.Stop()
		for event := this.RenderWindow.PollEvent(); event != nil; event = this.RenderWindow.PollEvent() {
			switch ev := event.(type) {
			case sf.EventKeyReleased:
				if ev.Code == sf.KeyEscape {
					this.RenderWindow.Close()
				}
				if ev.Code == sf.KeyReturn {
					this.GameStateInit()
					this.State = 0
				}
			case sf.EventClosed:
				this.RenderWindow.Close()
			}
		}
	} else if this.State == 3 {
		// ending when called a monster state
		this.Music.Stop()
		for event := this.RenderWindow.PollEvent(); event != nil; event = this.RenderWindow.PollEvent() {
			switch ev := event.(type) {
			case sf.EventKeyReleased:
				if ev.Code == sf.KeyEscape {
					this.RenderWindow.Close()
				}
				if ev.Code == sf.KeyReturn {
					this.GameStateInit()
					this.State = 0
				}
			case sf.EventClosed:
				this.RenderWindow.Close()
			}
		}
	}
}

func (this *Game) Draw() {
	this.RenderWindow.Clear(sf.Color{44, 176, 55, 255})
	if this.State == 0 {
		this.WelcomeText.Draw(this.RenderWindow, sf.DefaultRenderStates())
		this.PlayText.Draw(this.RenderWindow, sf.DefaultRenderStates())
	} else if this.State == 1 {
		for Enemy, _ := range this.EnemyList {
			this.EnemyList[Enemy].Draw(this.RenderWindow, sf.DefaultRenderStates())
		}
		this.Player.Draw(this.RenderWindow, sf.DefaultRenderStates())
		this.ChildsCollectedText.Draw(this.RenderWindow, sf.DefaultRenderStates())
		this.TimeText.Draw(this.RenderWindow, sf.DefaultRenderStates())
	} else if this.State == 2 || this.State == 3 {
		this.EndingText.Draw(this.RenderWindow, sf.DefaultRenderStates())
		this.EndingPointsText.Draw(this.RenderWindow, sf.DefaultRenderStates())
		this.RestartText.Draw(this.RenderWindow, sf.DefaultRenderStates())
		this.CreditsText.Draw(this.RenderWindow, sf.DefaultRenderStates())
	}
	this.RenderWindow.Display()
}

