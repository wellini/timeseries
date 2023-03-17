package timeseries

import "context"

type Repository interface {
	Store(ctx context.Context, t *TimeSeries) error
	GetAllSortByTime(ctx context.Context) ([]TimeSeries, error)
}
