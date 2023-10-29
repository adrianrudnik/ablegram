package ui

import (
	"fyne.io/fyne/v2/canvas"
	"github.com/adrianrudnik/ablegram/internal/stats"
	"time"
)

type UiUpdater struct {
	ticker     *time.Ticker
	statusText *canvas.Text
	statusLast bool
}

func NewUiUpdater(text *canvas.Text) *UiUpdater {
	return &UiUpdater{
		ticker:     time.NewTicker(500 * time.Millisecond),
		statusText: text,
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
	} else {
		u.statusText.Text = "The service has completed all tasks."
	}

	u.statusLast = progress.IsInProgress()
	u.statusText.Refresh()
}
