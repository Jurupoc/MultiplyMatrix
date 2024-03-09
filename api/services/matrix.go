package services

import (
	"errors"
	"log"
	"projectm/model"
)

type TempData struct {
	RowIndex 	int;
	ColumnIndex int;
	Data		int;
}

func isValidCase(matrixA, matrixB *model.Matrix) error {
	if matrixA.Cloumns != matrixB.Rows {
		return errors.New("matrix A Columns NOT EQUAL to matrix B Rows")
	}
	return nil
}

func MultiplyMatrix(matrixA, matrixB *model.Matrix) *model.Matrix {
	err := isValidCase(matrixA, matrixB)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	resultMatrix := model.NewMatrix(matrixA.Rows, matrixB.Cloumns)

	channel := make(chan TempData)
	for row_index := range matrixA.Rows{
		for column_index := range matrixB.Cloumns{
			go func(row_index, column_index int){
				sum := 0
				for columnA := range matrixA.Cloumns{
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
		for range matrixB.Cloumns{
			result := <- channel
			resultMatrix.Set(result.RowIndex, result.ColumnIndex, result.Data)
		}
	}

	return resultMatrix
}
