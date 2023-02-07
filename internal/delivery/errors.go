package delivery

var (
	ErrBadParams        = NewError("bad params width/height")
	ErrBadParsedAddress = NewError("incorrect url for source file")
	ErrReadImage        = NewError("reading image error")
	ErrImageProcessing  = NewError("image processing is not possible")
)

type Error struct {
	msg string
}

func NewError(msg string) *Error {
	return &Error{msg: msg}
}

func (err *Error) Error() string {
	return err.msg
}
