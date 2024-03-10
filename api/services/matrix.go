package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"projectm/model"
	"strconv"
)

type TempData struct {
	RowIndex 	int
	ColumnIndex int
	Data		int
}

func isValidCase(matrixA, matrixB *model.Matrix) error {
	if matrixA.Columns != matrixB.Rows {
		return fmt.Errorf("number of matrixA Columns NOT EQUAL to matrixB Rows")
	}
	return nil
}

func isValidCsv(csvNames []string) error {
	if len(csvNames) != 2 {
		return fmt.Errorf("expected 2 csv names got %d (%v)", len(csvNames), csvNames)
	}
	return nil
}

func FromCsv(csvFileName string) (*model.Matrix, error){
	file, err := os.Open(csvFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	Matrix := model.NewMatrix(len(data), len(data[0]))
	for row_index, row := range data {
		for column_index, value := range row {
			int_value, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}

			Matrix.Set(row_index, column_index, int_value)
		}
	}

	return Matrix, nil
}

func MultiplyMatrix(matrixA, matrixB *model.Matrix) (*model.Matrix, error) {
	err := isValidCase(matrixA, matrixB)
	if err != nil {
		return nil, err
	}
	resultMatrix := model.NewMatrix(matrixA.Rows, matrixB.Columns)

	channel := make(chan TempData)
	for row_index := range matrixA.Rows{
		for column_index := range matrixB.Columns{
			go func(row_index, column_index int){
				sum := 0
				for columnA := range matrixA.Columns{
					sum += matrixA.At(row_index, columnA) * matrixB.At(columnA, column_index)
				}

				channel <- TempData{
					RowIndex: row_index,
					ColumnIndex: column_index,
					Data: sum,
				}

			}(row_index, column_index)
		}
	}

	for range matrixA.Rows{
		for range matrixB.Columns{
			result := <- channel
			resultMatrix.Set(result.RowIndex, result.ColumnIndex, result.Data)
		}
	}

	return resultMatrix, nil
}

func MultiplyFromCsv(csvNames []string) (*model.Matrix, error){
	err := isValidCsv(csvNames)
	if err != nil {
		return nil, err
	}

	matrixA, err := FromCsv(csvNames[0])
	if err != nil {
		return nil, err
	}

	matrixB, err := FromCsv(csvNames[1])
	if err != nil {
		return nil, err
	}

	resultMatrix, err := MultiplyMatrix(matrixA, matrixB)
	if err != nil {
		return nil, err
	}

	return resultMatrix, nil
}
