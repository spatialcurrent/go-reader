package reader

type ErrorNotImplemented struct {
	Function string
	Object   string
}

func (e *ErrorNotImplemented) Error() string {
	return e.Function + " is not implemented by " + e.Object
}
