package main

import (
	"APP/internal/game"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

func StartGame(w fyne.Window, game game.Game, diceCount int) {
	// TURN TEXT
	labelTurn := widget.NewLabel("")
	labelTurn.Alignment = fyne.TextAlignCenter

	// LOADING TEXTURES
	diceTextures := map[int]fyne.Resource{}
	for i := 1; i <= 6; i++ {
		path := fmt.Sprintf("assets/die%d.png", i)
		res, err := fyne.LoadResourceFromPath(path)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
		diceTextures[i] = res
	}

	// CREATING CONTAINERS
	enemyHand := container.NewGridWithColumns(diceCount)
	playerHand := container.NewGridWithColumns(diceCount)
	
	// INIT VALUES
	game.Enemy.Reroll()
	game.Player.Reroll()

	// GETTING VALUES
	enemyValues := game.Enemy.GetDiceValues()
	playerValues := game.Player.GetDiceValues()
	labelTurn.SetText(game.GetTurn() + "'S MOVE")

	fmt.Println("ENEMY DICE VALUES:", enemyValues)
	fmt.Println("PLAYER DICE VALUES:", playerValues)
	
	// ADDING DICE TO CONTAINERS
	for i := 0; i < diceCount; i++ {
		if enemyValues[i] <= 0 || enemyValues[i] > 6 {
			fmt.Println("ERROR: INVALID ENEMY DIE VALUE AT INDEX", i)
		}
		if playerValues[i] <= 0 || playerValues[i] > 6 {
			fmt.Println("ERROR: INVALID PLAYER DIE VALUE AT INDEX", i)
		}

		die := canvas.NewImageFromResource(diceTextures[enemyValues[i]])
		if die != nil {
			die.SetMinSize(fyne.NewSize(64, 64))
			die.FillMode = canvas.ImageFillContain
			enemyHand.Add(die)
		}
		die = canvas.NewImageFromResource(diceTextures[playerValues[i]])
		if die != nil {
			die.SetMinSize(fyne.NewSize(64, 64))
			die.FillMode = canvas.ImageFillContain
			playerHand.Add(die)
		}
	}
	gameContent := container.NewVBox(labelTurn, enemyHand, playerHand)
	w.SetContent(gameContent)
}

func main() {
	//WINDOW INIT
	a := app.New()
	w := a.NewWindow("StAy_A_dDiE")
	w.Resize(fyne.NewSize(500, 500))

	labelTitle := widget.NewLabel("select dice count")

	//SELECT DICE COUNT
	var diceCount int
	diceCountRadio := widget.NewRadioGroup([]string{
		"3 dice",
		"4 dice",
		"5 dice",
	}, func(option string) {
		switch option {
		case "3 dice":
			diceCount = 3
		case "4 dice":
			diceCount = 4
		case "5 dice":
			diceCount = 5
		}
	})

	startButton := widget.NewButton("START GAME", func() {
		if(diceCount != 0) {
			var game game.Game = game.NewGame(diceCount)
			StartGame(w, game, diceCount)
		}
	})

	startContent := container.NewVBox(labelTitle, diceCountRadio, startButton)

	w.SetContent(startContent)

	w.ShowAndRun()
}