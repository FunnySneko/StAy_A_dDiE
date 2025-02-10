package main

import (
	"APP/internal/game"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢃⣬⡉⣿⣿⣿⢟⣭⡿⣭⢻⣿⣿⣿⣿⣿⣿⣿⡿⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣴⣶⣿⣿⣯⣿⣎⠛⣿⣷⣦⡻⣿⣿⣿⢸⣿⠀⣿⣿⣿⠟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⢹⣿⣿⣿⣿⡿⣱⠿⣿⣿⠿⡟⢁⡟⡌⣿⣿⠈⡟⠀⣿⢹⡟⢸⣿⠏⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣷⡙⢿⣿⠀⣿⢿⣿⢻⡳⣏⢦⣥⣤⣴⣸⣞⢧⠇⣙⠻⡄⠀⠀⢿⢸⠃⢸⣿⢸⡿⢋⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⡌⠻⣿⣿⣧⠘⣿⠀⠁⡾⢃⠮⠂⠙⠉⠈⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠨⠀⢸⠃⠏⢠⣾⣿⡿⠟⣫⣿⣿⣿⣿⣿
⣿⣿⣿⡿⢿⣿⡀⠈⠻⢿⣇⠈⠀⠐⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⠋⠁⣀⣾⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣮⣟⠷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣾⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣛⣯⠽⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣴⣾⣿⣿⣿⣿⣿
⣿⣿⣿⢟⣥⣶⠖⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⣤⣶⣾⣿⣿⢰⣿⣶⣶⡄⠀⠀⠀⠀⠀⠀⠀⠀⢀⡉⠛⢿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣛⣛⣋⣉⡠⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⣧⠹⣿⣿⣿⡟⣸⣿⣿⣿⡏⠀⣄⠀⠀⠀⠀⠀⠀⠀⢨⣙⣛⣛⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⡀⣏⣿⣿⣿⣿⣿⣿⢣⣿⣿⡿⠿⠿⠼⠹⠀⣀⠀⠀⠀⠀⠀⠠⢩⣿⣿⣿⣿⣿⡿⠟⢻
⣿⣿⣿⣿⣿⠃⣰⠃⠀⠀⠀⠀⠀⠀⠀⢸⣷⢹⠼⠟⠛⠻⣿⣿⣿⣯⢤⣴⣶⣶⣾⡷⣸⡿⡆⠀⠀⠀⢤⢷⣶⣶⣿⣿⣿⣿⡂⠛⢠
⣿⣿⣿⣿⡏⣴⣟⣴⠀⠁⠀⠀⠀⠀⠀⡄⣵⣶⣶⢿⠷⡵⣞⣿⣿⣿⣞⣽⢾⣿⣿⣵⣿⣳⠁⠀⠀⠀⠈⣊⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣯⣾⣟⡄⣠⠁⠀⠀⠀⢵⣷⣜⣿⣿⣻⠟⣛⣶⣿⣿⣿⣿⣮⣥⣆⣠⣤⣤⢏⣼⢀⠇⣀⡠⣬⡟⣏⣙⣯⣯⠻⢿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⡿⢋⣐⠀⠀⠀⠘⣟⣉⣁⣔⣦⣴⣶⣿⠟⠛⠛⠛⠿⠟⢻⣯⡽⣮⣿⡿⣼⠄⡗⣥⡅⣲⣿⣿⣿⣿⣿⣞⡽⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣷⡿⡋⣀⡀⢄⠈⣿⢷⡻⣲⣼⠋⢉⣠⣴⡸⣿⣷⣶⣶⣿⣿⣿⣿⣿⡇⣷⠟⣼⣿⡮⣿⢻⣿⡯⣿⣿⣿⠊⣿
⣿⣿⡿⢿⣿⣿⣿⣿⣿⣾⣾⣯⡘⢸⡞⡘⣿⣾⣿⣿⣿⣿⣿⣿⣽⣿⣽⣿⣿⣿⣿⣿⣿⣿⠋⠀⠀⠈⠉⡅⣮⡟⣿⣇⣽⣯⡅⣼⣿
⣿⠂⢶⢊⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠉⠻⠘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣛⣛⣿⣿⣿⣿⠏⠀⠀⠀⠀⠀⠀⠉⡴⣫⠶⠿⠷⣸⣿⣿
⣿⣿⣶⣿⣿⣿⣿⠿⠛⠿⢿⣿⡿⠃⠀⠀⠀⠈⠻⣿⣿⣿⣶⣿⣶⣽⣶⣿⣿⣿⣿⡿⢁⢄⠀⠀⠀⠀⠀⠀⢄⢥⡷⢟⡟⣿⣿⣿⣿
⣿⣿⣿⣿⣿⢯⡎⠀⢠⣾⣷⣮⠁⠀⠀⠀⠀⠀⠀⣈⡙⠿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠀⠀⠁⠀⠀⠀⠀⠀⠀⠈⠛⠿⠟⣼⣿⣿⣿⣿
⣿⣿⣿⣿⣿⠋⠽⢼⡟⠿⠽⣿⠀⠀⠀⠀⠀⠀⠀⢿⣿⣷⣄⣉⢛⡿⠿⠛⠫⠂⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⢤⣤⣾⢳⣤⣴⡿⠗⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠀⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⢿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⠇⠍⠙⠭⣭⢭⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠿⠿⣿⣿
⣿⣿⣿⣿⡿⠒⠓⢒⣰⣷⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉
⣿⣿⣿⣿⠷⠀⠀⠀⢉⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣿⣿⣿⣿⠀⠀⠐⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣿⣿⠛⠁⠀⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
*/

func StartGame(w fyne.Window, game game.Game) {
	labelTurn := widget.NewLabel("TURN")
	enemyHand := container.NewHBox()
	playerHand := container.NewHBox()
	gameContent := container.NewVBox(labelTurn, enemyHand, playerHand)
	w.SetContent(gameContent)

	go func() {
		for {
			game.Update()
			labelTurn.SetText(game.GetTurn() + "'S TURN")
		}
	}()
}

func main() {
	a := app.New()
	w := a.NewWindow("StAy_A_dDiE")
	w.Resize(fyne.NewSize(500, 500))

	labelTitle := widget.NewLabel("select dice count")

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
			StartGame(w, game)
		}
	})

	startContent := container.NewVBox(labelTitle, diceCountRadio, startButton)

	w.SetContent(startContent)

	w.ShowAndRun()
}
