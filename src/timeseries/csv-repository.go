package timeseries

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type CSVRepository struct {
	filename string
}

var _ Repository = (*CSVRepository)(nil)

const timeSeriesDirectory = ".timeseries"
const timeSeriesCSVFilename = "timeseries.csv"

func NewCSVRepository() *CSVRepository {

	ret := &CSVRepository{
		filename: filepath.Join(timeSeriesDirectory, string(os.PathSeparator), timeSeriesCSVFilename),
	}

	ret.initIfCSVDoesNotExist()

	return ret
}

func (C *CSVRepository) initIfCSVDoesNotExist() error {

	var err = (error)(nil)

	if _, err = os.Stat(timeSeriesDirectory); os.IsNotExist(err) {
		err = os.Mkdir(timeSeriesDirectory, fs.ModePerm)
	}

	if _, err = os.Stat(C.filename); os.IsNotExist(err) {
		f := (*os.File)(nil)
		f, err = os.Create(C.filename)
		err = f.Close()
	}

	return err
}

func (C *CSVRepository) Store(ctx context.Context, t *TimeSeries) error {
	var err = (error)(nil)

	err = C.initIfCSVDoesNotExist()
	if err != nil {
		return err
	}

	var csvLine = fmt.Sprintf("'%s', '%s', '%s', %d\n", t.ProjectName, t.TaskName, t.Time.String(), t.Type)

	f, err := os.OpenFile(C.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	if _, err = f.Write([]byte(csvLine)); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}

	return err
}

func (C *CSVRepository) GetAllSortByTime(ctx context.Context) ([]TimeSeries, error) {
	//TODO implement me
	panic("implement me")
}

var TsRepository = NewCSVRepository()
