package internal

import (
	"encoding/json"
	"errors"
	"io"
)

func assertDelim(dec *json.Decoder, delim json.Delim) {
	if !isNextDelim(dec, delim) {
		panic("expected delim")
	}
}

func isNextDelim(dec *json.Decoder, delim json.Delim) bool {
	t, eof := assertNextToken(dec)
	if eof {
		panic(io.EOF)
	}

	return t.(json.Delim) == delim
}

func assertString(dec *json.Decoder) string {
	t, eof := assertNextToken(dec)
	if eof {
		panic(io.EOF)
	}

	fieldName, ok := t.(string)
	if !ok {
		panic(errors.New("expected field name"))
	}

	return fieldName
}

func assertNextToken(dec *json.Decoder) (json.Token, bool) {
	t, err := dec.Token()

	if err == io.EOF {
		return nil, false
	}

	if err != nil {
		panic(err)
	}

	return t, true
}
