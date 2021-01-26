package errors

var (
	UnableToResolve = New("unable to resolve")
)

// WrongType return an error for the expected and actual types
func WrongType(expected, actual interface{}) error {
	return Errorf("wrong type: wanted %T, got %T", expected, actual)
}
