package app

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

const (
	OneHour     = time.Hour
	FiveMinutes = time.Minute * 5

	TwentyMinutes = time.Minute * 20
	TwentySeconds = time.Second * 20
)

type T struct{}

type MainApp struct {
	MainApp fyne.App
	pause   chan T
	resume  chan T
}

func NewMainApp() MainApp {
	return MainApp{
		MainApp: app.New(),
		pause:   make(chan T),
		resume:  make(chan T),
	}
}

func (m *MainApp) BreakInitializer(timeLeftLabel *widget.Label, progress *widget.ProgressBar) {
	for i := 1.0; i < OneHour.Seconds(); i++ {
		select {
		case <-m.pause:
			<-m.resume
		default:

		}

		switch {
		case i == OneHour.Seconds(): // 1 hour work - 5 minutes break
			i = 0.0 // reset timer
			go m.createRestPopUp(FiveMinutes)
			continue

		case int(i)%int(TwentyMinutes.Seconds()) == 0: // 20 - 20 - 20 rule
			go m.createRestPopUp(TwentySeconds)
		}

		timeLeft, err := time.ParseDuration(fmt.Sprintf("%0.1fs", OneHour.Seconds()-i))
		if err != nil {
			panic(err)
		}

		timeLeftLabel.SetText(fmt.Sprintf("Time to break: %s", timeLeft.Abs().String()))

		percentage := i / OneHour.Seconds()
		progress.SetValue(percentage)

		time.Sleep(time.Second) // Sleep at the end of calculation.
	}
}

func (m *MainApp) createRestPopUp(timeToRest time.Duration) {
	m.pause <- struct{}{} // send signal to pause main counter

	popUpWindow := m.MainApp.NewWindow("Go Rest Eyes")
	popUpWindow.Resize(fyne.Size{Width: 400, Height: 150})

	popUpBar := widget.NewProgressBar()
	popUpTimeLeft := widget.NewLabel(fmt.Sprintf("Time until the break ends: %s", timeToRest.Abs()))

	breakText := fmt.Sprintf("Start break - %s", timeToRest.Abs().String())

	popUpStartBreakButton := widget.NewButton(breakText, func() {
		popUpWindow.SetContent(container.NewVBox(popUpBar, popUpTimeLeft))
		m.timeBreak(timeToRest, popUpBar, popUpTimeLeft, popUpWindow)
	})

	popUpWindow.SetContent(popUpStartBreakButton)
	popUpWindow.RequestFocus()
	popUpWindow.Show()
}

func (m *MainApp) timeBreak(timeToRest time.Duration, bar *widget.ProgressBar, label *widget.Label, breakWindows fyne.Window) {
	estimatedEndTime := time.Now().Add(timeToRest).Format("15:04:05")

	for i := 0.0; i < timeToRest.Seconds(); i++ {
		timeLeft, err := time.ParseDuration(fmt.Sprintf("%0.1fs", timeToRest.Seconds()-i))
		if err != nil {
			panic(err)
		}

		percentage := i / timeToRest.Seconds()
		bar.SetValue(percentage)

		label.SetText(fmt.Sprintf("Time until the break ends: %s\n~%s", timeLeft.Abs().String(), estimatedEndTime))
		time.Sleep(time.Second)
	}

	breakEndLabel := widget.NewLabel("Break has ended")
	breakEndLabel.Alignment = fyne.TextAlignCenter

	closeBreak := widget.NewButton("Start working again", func() {
		breakWindows.Close()
		m.resume <- struct{}{} // on click send signal to resume main counter
	})

	breakWindows.SetContent(container.NewVBox(breakEndLabel, closeBreak))
}
