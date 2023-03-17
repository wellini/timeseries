package timeseries

import (
	"context"
)

type CSVRepository struct {
	filename string
}

var _ Repository = (*CSVRepository)(nil)

func NewCSVRepository(filename string) *CSVRepository {
	ret := &CSVRepository{
		filename: filename,
	}

	return ret
}

func (C *CSVRepository) Store(ctx context.Context, t *TimeSeries) error {
	//TODO implement me
	panic("implement me")
}

func (C *CSVRepository) GetAllSortByTime(ctx context.Context) ([]TimeSeries, error) {
	//TODO implement me
	panic("implement me")
}
