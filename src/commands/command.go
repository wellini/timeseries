package commands

import (
	"github.com/spf13/cobra"
	"timeseries/timeseries"
)

type TsCommand struct {
	TimeSeriesRepository timeseries.Repository
	Command              cobra.Command
}
