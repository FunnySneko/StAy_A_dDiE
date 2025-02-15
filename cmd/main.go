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

func StartGame(w fyne.Window, g game.Game, diceCount int) {
	const diceSize = 100
	const playerIndexEnemy = 0
	const playerIndexPlayer = 1

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

	// TOTAL VALUE LABELS
	labelEnemyValue := widget.NewLabel("0")
	labelEnemyValue.Alignment = fyne.TextAlignCenter
	labelEnemyValue.TextStyle.Bold = true
	labelPlayerValue := widget.NewLabel("0")
	labelPlayerValue.Alignment = fyne.TextAlignCenter
	labelPlayerValue.TextStyle.Bold = true

	// INIT VALUES
	g.Enemy.Reroll()
	g.Player.Reroll()
	for i := range g.Enemy.RollOpportunities {
		g.Enemy.RollOpportunities[i] = 1
	}
	for i := range g.Player.RollOpportunities {
		g.Player.RollOpportunities[i] = 1
	}

	// CREATING HAND CONTAINERS
	enemyHand := container.NewGridWithColumns(diceCount)
	playerHand := container.NewGridWithColumns(diceCount)

	// CREATING PLAYER REROLL POINTER
	var playerRerollButton *widget.Button

	// CREATING ENEMY REROLL BUTTON POINTERS
	var enemyRerollButtons []widget.Button = make([]widget.Button, diceCount)
	// CREATING ENEMY REROLL BUTTON CONTAINER
	enemyRerollButtonsContainer := container.NewGridWithColumns(diceCount)

	updateUI := func() {
		// GETTING VALUES
		enemyValues := g.Enemy.GetDiceValues()
		playerValues := g.Player.GetDiceValues()
		labelEnemyValue.SetText(fmt.Sprint(g.Enemy.GetDiceTotalValue()))
		labelPlayerValue.SetText(fmt.Sprint(g.Player.GetDiceTotalValue()))

		fmt.Println("ENEMY DICE VALUES:", enemyValues)
		fmt.Println("PLAYER DICE VALUES:", playerValues)

		fmt.Println("ENEMY ROLL OPPORTUNITIES:", g.Enemy.RollOpportunities)
		fmt.Println("PLAYER ROLL OPPORTUNITIES:", g.Player.RollOpportunities)

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

		for i := range g.Enemy.RollOpportunities {
			if g.Enemy.RollOpportunities[i] == 0 {
				enemyRerollButtons[i].Disable()
			} else {
				enemyRerollButtons[i].Enable()
				if g.GetTurn() == game.EnemyTurn {
					enemyRerollButtons[i].Disable()
				}
			}
			enemyRerollButtons[i].Refresh()
		}

		if g.GetTurn() == game.EnemyTurn {
			labelTurn.SetText("ENEMY'S TURN")
			playerRerollButton.Disable()
		} else {
			labelTurn.SetText("PLAYER'S TURN")
			playerRerollButton.Enable()
		}

		// REFRESHING CONTAINERS
		enemyHand.Refresh()
		playerHand.Refresh()
		enemyRerollButtonsContainer.Refresh()
	}

	rerollDice := func(playerIndex int) {
		if playerIndex == playerIndexEnemy {
			g.Enemy.Reroll()
		} else {
			g.Player.Reroll()
		}
		g.Update()
		updateUI()
	}

	rollDie := func(playerIndex, index int) {
		if playerIndex == playerIndexEnemy {
			g.Enemy.RollDie(index)
		} else {
			g.Player.RollDie(index)
		}
		g.Update()
		updateUI()
	}

	// CREATING PLAYER REROLL BUTTON
	playerRerollButton = widget.NewButton("REROLL", func() {
		rerollDice(playerIndexPlayer)
	})

	// CREATING ENEMY ROLL BUTTONS
	for i := 0; i < diceCount; i++ {
		enemyRerollButtons[i] = *widget.NewButton("REROLL", func() {
			rollDie(playerIndexEnemy, i)
		})
		enemyRerollButtonsContainer.Add(&enemyRerollButtons[i])
	}

	gameContent := container.NewVBox(
		labelTurn,
		layout.NewSpacer(),
		enemyRerollButtonsContainer,
		enemyHand,
		layout.NewSpacer(),
		labelEnemyValue,
		labelPlayerValue,
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
	w.Resize(fyne.NewSize(530, 420))

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
		if diceCount != 0 {
			var game game.Game = game.NewGame(diceCount)
			StartGame(w, game, diceCount)
		}
	})

	startContent := container.NewVBox(labelTitle, diceCountRadio, startButton)

	w.SetContent(startContent)

	w.ShowAndRun()
}