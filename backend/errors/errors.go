package errors

import "fmt"

type parseError struct {
	cursor int
	expect string
	actual byte
}

// NewParseError returns the parse error message
func NewParseError(cursor int, expect string, actual byte) *parseError {
	return &parseError{cursor, expect, actual}
}

func (pe *parseError) Error() string {
	if pe.expect != "" {
		return fmt.Sprintf("invalid char at pos %d, expect '%s' got '%s'", pe.cursor, pe.expect, []byte{pe.actual})
	}

	return fmt.Sprintf("invalid char at pos %d, got '%s'", pe.cursor, []byte{pe.actual})
}

type eofError struct {
	cursor int
	expect string
}

// NewEOFError returns the input error message
func NewEOFError(cursor int, expect string) *eofError {
	return &eofError{cursor, expect}
}

func (ie *eofError) Error() string {
	if ie.expect == "" {
		return fmt.Sprintf("invalid char at pos %d, got EOF", ie.cursor)
	}

	return fmt.Sprintf("invalid char at pos %d, expect '%s' got EOF", ie.cursor, ie.expect)
}

type overflowError struct {
	cursor int
	actual []byte
	rtype  string
}

// NNewOverflowError returns the overflow error message
func NewOverflowError(cursor int, actual []byte, rtype string) *overflowError {
	return &overflowError{cursor, actual, rtype}
}

func (oe *overflowError) Error() string {
	return fmt.Sprintf("invalid number length get at pos %d when decode to %s, got %s", oe.cursor, oe.rtype, oe.actual)
}
