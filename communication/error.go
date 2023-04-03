package communication

import (
	"fmt"
)

const (
	PROGRAM_FAILED_ERROR_MESSAGE = "was not able to get the output of the program {%v}, it return the error\n{%v}"

	UNABLE_TO_DECODE_JSON_ERROR_MESSAGE = "was not able decode the JSON return error {%v}"
)

type ProgramFailedError struct {
	Program string
	Message string
}

func (p ProgramFailedError) Error() string {
	return fmt.Sprintf(PROGRAM_FAILED_ERROR_MESSAGE, p.Program, p.Message)
}

type UnableToDecodeJsonError struct {
	Message string
}

func (u UnableToDecodeJsonError) Error() string {
	return fmt.Sprintf(UNABLE_TO_DECODE_JSON_ERROR_MESSAGE, u.Message)
}
