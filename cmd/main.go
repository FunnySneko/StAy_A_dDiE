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

func NewEnemyRollButtonsContainer(g game.Game) *fyne.Container {
	count := len(g.Enemy.GetDiceValues())
	enemyRollButtonContainer := container.NewGridWithColumns(count)
	for i := 0; i < count; i++ {
		enemyRollButton := widget.NewButton("ROLL", func() {})
		enemyRollButtonContainer.Add(enemyRollButton)
	}
	return enemyRollButtonContainer
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
	enemyHand := NewDiceHandContainer(g.Enemy.GetDiceValues())
	playerHand := NewDiceHandContainer(g.Player.GetDiceValues())
	// HEALTH
	enemyHealth := widget.NewLabel(fmt.Sprint("[ ENEMY'S HEALTH: ", g.Enemy.Health, " ]"))
	playerHealth := widget.NewLabel(fmt.Sprint("[ YOUR HEALTH: ", g.Player.Health, " ]"))
	playerHealth.Alignment = fyne.TextAlignTrailing
	// VALUES
	enemyDiceValue := widget.NewLabel(fmt.Sprint(g.Enemy.GetDiceTotalValue()))
	enemyDiceValue.Alignment = fyne.TextAlignCenter
	enemyDiceValue.TextStyle.Bold = true
	playerDiceValue := widget.NewLabel(fmt.Sprint(g.Player.GetDiceTotalValue()))
	playerDiceValue.Alignment = fyne.TextAlignCenter
	playerDiceValue.TextStyle.Bold = true
	// REROLL BUTTON
	buttonPlayerReroll := widget.NewButton("REROLL YOUR DICE", func() {})

	content := container.NewVBox(NewStageDisplayer(g),
		NewHeaderText("FIGHT"),
		layout.NewSpacer(),
		NewEnemyRollButtonsContainer(g),
		enemyHand,
		container.NewGridWithColumns(3, enemyHealth, enemyDiceValue),
		widget.NewSeparator(),
		container.NewGridWithColumns(3, layout.NewSpacer(), playerDiceValue, playerHealth),
		playerHand,
		buttonPlayerReroll,
		layout.NewSpacer())
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
	w.Resize(fyne.NewSize(530, 650))
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