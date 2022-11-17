package mediators

import (
	"errors"
	"os"

	"api-bootcamp/dto"

	"github.com/gocarina/gocsv"
)

type ApiMediator interface {
	GetCSVElementByID(id int) (dto.GetCsvDTO, error)
}

type apiMediator struct {
	CsvFile string
}

func NewApiMediator(csvFile string) ApiMediator {
	return &apiMediator{
		CsvFile: csvFile,
	}
}

func (ap *apiMediator) GetCSVElementByID(id int) (dto.GetCsvDTO, error) {
	csvDTo := []*dto.GetCsvDTO{}
	in, err := os.Open(ap.CsvFile)
	if err != nil {
		return dto.GetCsvDTO{}, err
	}
	defer in.Close()

	if err := gocsv.UnmarshalFile(in, &csvDTo); err != nil {
		return dto.GetCsvDTO{}, err
	}
	for _, client := range csvDTo {
		if client.ID == id {
			return *client, nil
		}
	}

	return dto.GetCsvDTO{}, errors.New("element not found")
}
