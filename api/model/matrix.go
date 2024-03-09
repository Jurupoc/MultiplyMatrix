package model

type Matrix struct {
	Rows    int
	Cloumns int
	Data    []int
}

func NewMatrix(row, column int) *Matrix{
	return &Matrix{
		Rows: row,
		Cloumns: column,
		Data: make([]int, row*column),
	}
}

func getIndex(matrix *Matrix, row, column int) int {
	return (row * matrix.Cloumns) + column
}

func At(matrix *Matrix, row, column int) int {
	return matrix.Data[getIndex(matrix, row, column)]
}

func Set(matrix *Matrix, row, column, value int) {
	matrix.Data[getIndex(matrix, row, column)] = value
}
