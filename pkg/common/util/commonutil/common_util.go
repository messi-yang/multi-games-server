package commonutil

func MapMatrix[T any, M any](inputMatrix [][]T, transformer func(i int, j int, obj T) (M, error)) ([][]M, error) {
	outputMatrix := make([][]M, 0)

	for i := 0; i < len(inputMatrix); i += 1 {
		outputMatrix = append(outputMatrix, make([]M, 0))
		for j := 0; j < len(inputMatrix[0]); j += 1 {
			transformedObj, err := transformer(i, j, inputMatrix[i][j])
			if err != nil {
				return nil, err
			}
			outputMatrix[i] = append(outputMatrix[i], transformedObj)
		}
	}
	return outputMatrix, nil
}

func RangeMatrix(width int, height int, callback func(i int, j int)) {
	for i := 0; i < width; i += 1 {
		for j := 0; j < height; j += 1 {
			callback(i, j)
		}
	}
}
