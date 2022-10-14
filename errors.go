package libunzeep

type CanNotReadZipError struct {
	err error
}

func (e CanNotReadZipError) Error() string {
	msg := "can not read zip file"
	if e.err == nil {
		return "can not read zip file"
	}
	return msg + ": " + e.err.Error()
}
func (err CanNotReadZipError) Unwrap() error { return err.err }
