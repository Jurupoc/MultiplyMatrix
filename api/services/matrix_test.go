package services

import (
	"testing"
	"projectm/model"
)

func fillMatrix(matrix model.Matrix) {
	for row_index := range matrix.Rows {
		for column_index := range matrix.Cloumns{
			matrix.Set(row_index, column_index, 1)
		}
	}
}

func TestMatrixMultiplication(t *testing.T) {
	testMatrixA := model.NewMatrix(2, 2)
	testMatrixB := model.NewMatrix(2, 2)

	fillMatrix(*testMatrixA)
	fillMatrix(*testMatrixB)

	result := MultiplyMatrix(testMatrixA, testMatrixB)
	for row_index := range testMatrixA.Rows{
		for column_index := range testMatrixA.Cloumns{
			if result.At(row_index, column_index) != 2 {
				t.Errorf("erro on multiply matrix expected value %d got: %v", 2, result.Data)
			}
		}
	}

}
