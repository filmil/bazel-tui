package main

import (
	"bytes"
	"flag"
	"os/exec"
	"strings"

	"github.com/epiclabs-io/winman"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"fmt"
	"time"
)

func main() {

	var cmdline string
	var pause time.Duration

	flag.StringVar(&cmdline, "cmdline", "", "command line")
	flag.DurationVar(&pause, "pause", 3*time.Second*10, "Delay")
	flag.Parse()

	app := tview.NewApplication()
	wm := winman.NewWindowManager()

	content := tview.NewTextView().
		SetText("loading...").
		SetTextAlign(tview.AlignCenter)

	go func() {
		if cmdline == "" {
			return
		}
		args := strings.Split(cmdline, " ")
		for {
			cmd := exec.Command(args[0], args[1:]...)
			var stdout bytes.Buffer
			cmd.Stdout = &stdout
			var stderr bytes.Buffer
			cmd.Stderr = &stderr

			err := cmd.Run()
			if err != nil {
				app.QueueUpdateDraw(func() {
					content.SetText(fmt.Sprintf("error: %v", err))
				})
			} else {
				app.QueueUpdateDraw(func() {
					stderrText := stderr.String()
					if stderrText != "" {
						content.SetText(stderrText)
					} else {
						content.SetText(stdout.String())
					}
				})
			}
			<-time.After(pause)
		}
	}()

	mouseLog := tview.NewTextView().SetText("mouse log:")

	mouseWin := wm.NewWindow().
		Show().
		SetRoot(mouseLog).
		SetTitle("Mouse").
		SetDraggable(true).
		SetResizable(true)

	window := wm.NewWindow(). // create new window and add it to the window manager
					Show().             // make window visible
					SetRoot(content).   // have the text view above be the content of the window
					SetDraggable(true). // make window draggable around the screen
					SetResizable(true). // make the window resizable
					SetTitle("Hi!").    // set the window title
					AddButton(
			&winman.Button{ // create a button with an X to close the application
				Symbol:  'X',
				OnClick: func() { app.Stop() }, // close the application
			}).
		AddButton(
			&winman.Button{ // create a button with an X to close the application
				Symbol:  '+',
				OnClick: func() {},
			}).
		SetMouseCapture(
			func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
				fmt.Fprintf(mouseLog, "action: %+v event: %+v\n", action, event)
				return action, event
			},
		)

	window.SetRect(5, 5, 30, 10) // place the window

	wm.AddWindow(mouseWin)

	// now, execute the application:
	if err := app.SetRoot(wm, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
