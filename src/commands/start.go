package commands

import (
	"context"
	"errors"
	"fmt"
	"time"
	"timeseries/timeseries"
)

var repo timeseries.Repository

func init() {
	const filename = "./asdadd.csv"
	repo = timeseries.NewCSVRepository(filename)
}

func Start(ctx context.Context, opts CommandLineOptions, args []string) error {
	if len(args) == 0 {
		return errors.New("job name missing")
	}

	projectName := opts.Flag("project")
	if len(projectName) == 0 {
		projectName = "Global"
	}

	ts := &timeseries.TimeSeries{
		Type:        timeseries.Start,
		Time:        time.Now(),
		ProjectName: projectName,
		TaskName:    args[0],
	}

	err := repo.Store(ctx, ts)
	if err != nil {
		return fmt.Errorf("could not store timeseries: %w", err)
	}

	return nil
}
