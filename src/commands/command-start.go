package commands

import (
	"context"
	"time"
	"timeseries/timeseries"
	"timeseries/utils"
)

func start(ctx context.Context, repository timeseries.Repository, rawTaskId string) error {

	taskId, parsingError := utils.ParseJobId(rawTaskId)

	if parsingError != nil {
		return parsingError
	}

	ts := &timeseries.TimeSeries{
		Type:        timeseries.Start,
		Time:        time.Now(),
		ProjectName: taskId.ProjectName,
		TaskName:    taskId.TaskName,
	}

	return repository.Store(ctx, ts)
}
