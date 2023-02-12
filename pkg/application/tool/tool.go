package tool

func MapMatrix[T any, M any](inputMatrix [][]T, transformer func(x int, y int, obj T) (M, error)) ([][]M, error) {
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

func RangeMatrix(width int, height int, callback func(x int, y int)) {
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			callback(x, y)
		}
	}
}
