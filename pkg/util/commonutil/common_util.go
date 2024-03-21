package commonutil

func MapWithError[T any, M any](inputArray []T, transformer func(i int, obj T) (M, error)) ([]M, error) {
	outputArray := make([]M, 0)

	for i := 0; i < len(inputArray); i += 1 {
		transformedObj, err := transformer(i, inputArray[i])
		if err != nil {
			return nil, err
		}
		outputArray = append(outputArray, transformedObj)
	}
	return outputArray, nil
}

func RangeMatrix(width int, height int, callback func(x int, z int) error) error {
	for x := 0; x < width; x += 1 {
		for z := 0; z < height; z += 1 {
			err := callback(x, z)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ToPointer[T any](v T) *T {
	return &v
}
