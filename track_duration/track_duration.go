package trackduration

import (
	"log"
	"time"
)

type TrackDuration struct {
	label          string
	periodDuration time.Duration
	lastTime       time.Time
}

func New(label string, periodDuration time.Duration) *TrackDuration {
	td := &TrackDuration{
		label:          label,
		periodDuration: periodDuration}
	return td
}

func (td *TrackDuration) Track(do func()) {
	start := time.Now()
	do()
	if start.Sub(td.lastTime) >= td.periodDuration {
		log.Println(td.label, ":时间损耗:", time.Since(start))
		td.lastTime = start
	}
}

func Track(label string, do func()) {
	start := time.Now()
	do()
	log.Println(label, " 时间损耗:", time.Since(start))
}
