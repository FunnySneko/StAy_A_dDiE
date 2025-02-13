package main

import (
	"APP/internal/game"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func StartGame(w fyne.Window, game game.Game, diceCount int) {
	const diceSize = 100

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

	// INIT VALUES
	game.Enemy.Reroll()
	game.Player.Reroll()

	// CREATING CONTAINERS
	enemyHand := container.NewGridWithColumns(diceCount)
	playerHand := container.NewGridWithColumns(diceCount)

	updateUI := func() {
		// GETTING VALUES
		enemyValues := game.Enemy.GetDiceValues()
		playerValues := game.Player.GetDiceValues()
		labelTurn.SetText(game.GetTurn() + "'S TURN")

		fmt.Println("ENEMY DICE VALUES:", enemyValues)
		fmt.Println("PLAYER DICE VALUES:", playerValues)

		// CLEANING CONTAINERS
		enemyHand.Objects = nil
		playerHand.Objects = nil

		for i := 0; i < diceCount; i++ {
			die := canvas.NewImageFromResource(diceTextures[enemyValues[i]])
			die.SetMinSize(fyne.NewSize(diceSize, diceSize))
			die.FillMode = canvas.ImageFillContain
			enemyHand.Add(die)

			die = canvas.NewImageFromResource(diceTextures[playerValues[i]])
			die.SetMinSize(fyne.NewSize(diceSize, diceSize))
			die.FillMode = canvas.ImageFillContain
			playerHand.Add(die)
		}

		// REFRESHING CONTAINERS
		enemyHand.Refresh()
		playerHand.Refresh()
	}

	// CREATING PLAYER BUTTON
	playerRerollButton := widget.NewButton("REROLL", func() {
		game.Player.Reroll()
		updateUI()
	})

	enemyRerollButtons := container.NewGridWithColumns(diceCount)
	for i := 0; i < diceCount; i++ {
		enemyRerollButton := widget.NewButton("REROLL", func(id int) func() {
			return func() {
				game.Enemy.RollDie(id)
				updateUI()
			}
		}(i))
		enemyRerollButtons.Add(enemyRerollButton)
	}

	gameContent := container.NewVBox(
		labelTurn,
		layout.NewSpacer(),
		enemyRerollButtons,
		enemyHand,
		layout.NewSpacer(),
		playerHand,
		playerRerollButton,
		layout.NewSpacer(),
	)

	w.SetContent(gameContent)

	updateUI()
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