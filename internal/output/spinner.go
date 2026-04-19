package output

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

type Spinner struct {
	label string
	stop  chan struct{}
	done  chan struct{}
}

func NewSpinner(label string) *Spinner {
	return &Spinner{
		label: label,
		stop:  make(chan struct{}),
		done:  make(chan struct{}),
	}
}

func (s *Spinner) Start() {
	go func() {
		defer close(s.done)
		i := 0
		for {
			select {
			case <-s.stop:
				return
			default:
				fmt.Printf("\r%s %s", color.New(color.FgCyan).Sprint(spinnerFrames[i%len(spinnerFrames)]), s.label)
				i++
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()
}

func (s *Spinner) Success() {
	close(s.stop)
	<-s.done
	fmt.Printf("\r%s %s\n", color.New(color.FgGreen).Sprint("✓"), s.label)
}

func (s *Spinner) Fail() {
	close(s.stop)
	<-s.done
	fmt.Printf("\r%s %s\n", color.New(color.FgRed).Sprint("✗"), s.label)
}
