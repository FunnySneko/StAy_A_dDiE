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
	"image/color"
)

func NewHeaderText(text string) fyne.CanvasObject {
	labelHeader := canvas.NewText(text, color.Black)
	labelHeader.Alignment = fyne.TextAlignCenter
	labelHeader.TextStyle.Bold = true
	labelHeader.TextSize = 20
	return labelHeader
}

func NewEnemyRollButtons(g game.Game, onRoll func(index int)) []*widget.Button {
	count := len(g.Enemy.GetDiceValues())
	enemyRollButtons := make([]*widget.Button, 0, count)
	for i := 0; i < count; i++ {
		index := i
		enemyRollButton := widget.NewButton("ROLL", func() {
			onRoll(index)
		})
		enemyRollButtons = append(enemyRollButtons, enemyRollButton)
	}
	return enemyRollButtons
}

func NewDiceHandContainer(values []int) *fyne.Container {
	diceHandContainer := container.NewGridWithColumns(len(values))
	for _, value := range values {
		dieImage := canvas.NewImageFromResource(diceTextures[value])
		dieImage.SetMinSize(fyne.NewSize(100, 100))
		dieImage.FillMode = canvas.ImageFillContain
		diceHandContainer.Add(dieImage)
	}
	return diceHandContainer
}

func NewStageDisplayer(g game.Game) fyne.Widget {
	labelStage := widget.NewLabel(fmt.Sprint("STAGE ", g.Stage))
	labelStage.Alignment = fyne.TextAlignCenter
	return labelStage
}

var diceTextures = map[int]fyne.Resource{}

func ChangeScreen(w fyne.Window, g game.Game, e game.Event) {
	switch e {
	case game.Fight:
		FightScreen(w, g)
	}
}

func FightScreen(w fyne.Window, g game.Game) {
	// HAND CONTAINERS
	enemyHand := container.NewGridWithColumns(len(g.Enemy.Dice))
	playerHand := container.NewGridWithColumns(len(g.Player.Dice))

	// VALUES
	enemyDiceValue := widget.NewLabel(fmt.Sprint(g.Enemy.GetDiceTotalValue()))
	enemyDiceValue.Alignment = fyne.TextAlignCenter
	enemyDiceValue.TextStyle.Bold = true
	playerDiceValue := widget.NewLabel(fmt.Sprint(g.Player.GetDiceTotalValue()))
	playerDiceValue.Alignment = fyne.TextAlignCenter
	playerDiceValue.TextStyle.Bold = true

	//HEALTH
	enemyHealth := widget.NewLabel("")
	playerHealth := widget.NewLabel("")
	playerHealth.Alignment = fyne.TextAlignTrailing

	// TURN
	turn := widget.NewLabel("TURN")
	turn.Alignment = fyne.TextAlignCenter

	// ENEMY ROLL BUTTONS
	enemyRollButtons := make([]*widget.Button, len(g.Enemy.Dice))

	// REROLL BUTTON
	var buttonPlayerReroll *widget.Button

	updateUI := func() {
		for i := range g.Enemy.Dice {
			if g.Enemy.RollOpportunities[i] == 0 || g.Turn == game.EnemyTurn {
				enemyRollButtons[i].Disable()
			} else {
				enemyRollButtons[i].Enable()
			}
			enemyRollButtons[i].Refresh()
		}
		enemyHand.Objects = NewDiceHandContainer(g.Enemy.GetDiceValues()).Objects
		playerHand.Objects = NewDiceHandContainer(g.Player.GetDiceValues()).Objects
		enemyDiceValue.SetText(fmt.Sprint(g.Enemy.GetDiceTotalValue()))
		playerDiceValue.SetText(fmt.Sprint(g.Player.GetDiceTotalValue()))
		enemyHealth.SetText(fmt.Sprint("[ ENEMY'S HEALTH: ", g.Enemy.Health, " ]"))
		playerHealth.SetText(fmt.Sprint("[ YOUR HEALTH: ", g.Player.Health, " ]"))
		if g.Turn == game.EnemyTurn {
			turn.SetText("ENEMY'S TURN")
			buttonPlayerReroll.Disable()
		} else {
			turn.SetText("YOUR TURN")
			buttonPlayerReroll.Enable()
		}
		enemyHand.Refresh()
		playerHand.Refresh()
		enemyDiceValue.Refresh()
		playerDiceValue.Refresh()
		buttonPlayerReroll.Refresh()
	}

	enemyTurn := func() {
		g.NextTurn()
		updateUI()
		g.EnemyMove()
		g.NextTurn()
		updateUI()
	}

	rollEnemyDie := func(index int) {
		g.Enemy.RollDie(index)
		enemyTurn()
	}

	rerollPlayerDice := func() {
		g.Player.Reroll()
		enemyTurn()
	}

	buttonPlayerReroll = widget.NewButton("REROLL YOUR DICE", func() {
		rerollPlayerDice()
	})

	enemyRollButtons = NewEnemyRollButtons(g, rollEnemyDie)
	enemyRollButtonsContainer := container.NewGridWithColumns(len(g.Enemy.Dice))
	for i := range g.Enemy.Dice {
		enemyRollButtonsContainer.Add(enemyRollButtons[i])
	}

	content := container.NewVBox(
		NewStageDisplayer(g),
		NewHeaderText("FIGHT"),
		layout.NewSpacer(),
		turn,
		enemyRollButtonsContainer,
		enemyHand,
		container.NewGridWithColumns(3, enemyHealth, enemyDiceValue),
		widget.NewSeparator(),
		container.NewGridWithColumns(3, layout.NewSpacer(), playerDiceValue, playerHealth),
		playerHand,
		buttonPlayerReroll,
		layout.NewSpacer(),
	)

	updateUI()

	w.SetContent(content)
}

func StartScreen(w fyne.Window, g game.Game) {

	labelUnderDice := widget.NewLabel(fmt.Sprint("YOUR STARTING HAND: ", g.Player.GetDiceTotalValue()))
	labelUnderDice.Alignment = fyne.TextAlignCenter
	buttonStart := widget.NewButton("START", func() {
		ChangeScreen(w, g, g.NextStage())
	})
	content := container.NewVBox(NewStageDisplayer(g),
		NewHeaderText("YOUR JOURNEY STARTS HERE"),
		layout.NewSpacer(),
		NewDiceHandContainer(g.Player.GetDiceValues()),
		labelUnderDice, buttonStart,
		layout.NewSpacer())
	w.SetContent(content)
}

func main() {
	a := app.New()
	w := a.NewWindow("StAy_A_dDiE")
	w.Resize(fyne.NewSize(550, 600))
	g := game.NewGame()
	for i := 1; i <= 6; i++ {
		path := fmt.Sprintf("assets/die%d.png", i)
		res, err := fyne.LoadResourceFromPath(path)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
		diceTextures[i] = res
	}
	StartScreen(w, g)
	w.ShowAndRun()
}
