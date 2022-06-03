package gameworker

type errMissingGameServiceDependency struct {
}

func (e *errMissingGameServiceDependency) Error() string {
	return "The dependency gameservice.GameService is missing in \"gameworker\"."
}

type errMissingMessageServiceDependency struct {
}

func (e *errMissingMessageServiceDependency) Error() string {
	return "The dependency messageservice.MessageService is missing in \"gameworker\"."
}
