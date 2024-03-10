package handlers

import (
	"fmt"
	"net/http"
	"projectm/services"
)

var expectedFields = []string{"csvA", "csvB"}

func getCsvName(request *http.Request, names []string) ([]string, error) {
	csvNames := make([]string, len(names))

	for index, name := range names {
		_, header, err := request.FormFile(name)
		if err != nil {
			return nil, fmt.Errorf("%v - %v", err, names)
		}
		csvNames[index] = header.Filename
	}

	return csvNames, nil
}

func CsvHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadGateway)
		return
	}

	csvNames, err := getCsvName(request, expectedFields)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	resultMatrix, err := services.MultiplyFromCsv(csvNames)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(resultMatrix.Data)
}

