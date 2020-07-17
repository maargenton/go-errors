package errors

// Sentinel is sentinel error type that can be set as a constant
type Sentinel string

func (e Sentinel) Error() string {
	return string(e)
}
