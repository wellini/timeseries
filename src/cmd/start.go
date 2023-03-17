package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
	"timeseries/timeseries"
)

var Start = &cobra.Command{
	Use:  "start",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var currentPath, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("could not obtain working directory: %w", err)
		}
		var splitPath = strings.Split(currentPath, string(os.PathSeparator))
		var currentDir = splitPath[len(splitPath)-1]

		var ts = &timeseries.TimeSeries{
			Type:        timeseries.Start,
			Time:        time.Now(),
			ProjectName: currentDir,
			TaskName:    args[0],
		}

		err = timeseries.TsRepository.Store(cmd.Context(), ts)
		if err != nil {
			return fmt.Errorf("could not store timeseries: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(Start)
}
