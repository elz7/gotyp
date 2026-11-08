package game

import (
	"fmt"
	"time"
)

func formatDuration(d time.Duration) string {
	if d.Minutes() == 0 {
		return fmt.Sprintf("00:%02d", int(d.Seconds()))
	}
	return fmt.Sprintf("%02d:%02d", int(d.Minutes()), int(d.Seconds())%60)
}
