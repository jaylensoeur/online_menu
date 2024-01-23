package http_errors

import (
	"errors"
)

type BadRequestError struct {
}

func (b *BadRequestError) Errors() error {
	return errors.New("bad request")
}
