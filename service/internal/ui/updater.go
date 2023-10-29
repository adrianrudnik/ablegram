package ui

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"time"
)

type UiUpdater struct {
	ticker      *time.Ticker
	statusLast  bool
	statusText  *canvas.Text
	progressBar *widget.ProgressBarInfinite
}

func NewUiUpdater(text *canvas.Text, infinite *widget.ProgressBarInfinite) *UiUpdater {
	return &UiUpdater{
		ticker:      time.NewTicker(500 * time.Millisecond),
		statusText:  text,
		progressBar: infinite,
	}
}

func (u *UiUpdater) Run(progress *stats.ProcessProgress) {
	for {
		select {
		case <-u.ticker.C:
			u.updateStatusText(progress)
		}
	}
}

func (u *UiUpdater) updateStatusText(progress *stats.ProcessProgress) {
	if u.statusLast == progress.IsInProgress() {
		return
	}

	if progress.IsInProgress() {
		u.statusText.Text = "The service is processing files."
		u.progressBar.Start()
	} else {
		u.statusText.Text = "The service has completed all tasks."
		u.progressBar.Hide()
	}

	u.statusLast = progress.IsInProgress()
	u.statusText.Refresh()
}
