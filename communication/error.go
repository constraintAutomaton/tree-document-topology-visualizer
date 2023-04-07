package communication

import (
	"fmt"
)

const (
	PROGRAM_FAILED_ERROR_MESSAGE = "was not able to get the output of the program {%v}, it return the error\n{%v}"

	UNABLE_TO_DECODE_JSON_ERROR_MESSAGE = "was not able decode the JSON return error {%v}"
)

// ProgramFailedError is an error that describes that an external program called by a process has failed.
type ProgramFailedError struct {
	Program string
	Message string
}

func (p ProgramFailedError) Error() string {
	return fmt.Sprintf(PROGRAM_FAILED_ERROR_MESSAGE, p.Program, p.Message)
}

// UnableToDecodeJsonError is an error that describes that that the current program was not able to decode a serialized JSON.
type UnableToDecodeJsonError struct {
	Message string
}

func (u UnableToDecodeJsonError) Error() string {
	return fmt.Sprintf(UNABLE_TO_DECODE_JSON_ERROR_MESSAGE, u.Message)
}
