package timeseries

import "time"

type Type int

const (
	Unset Type = iota
	Start
	End
)

type TimeSeries struct {
	Type        Type
	Time        time.Time
	ProjectName string
	JobName     string
}
