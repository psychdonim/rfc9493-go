package internal

import (
	"bytes"
	"encoding/json"
)

type SubjectId struct {
	Id
	Aliases []*Id
}

func (s *SubjectId) UnmarshallJSON(data []byte) {
	dec := json.NewDecoder(bytes.NewReader(data))

	assertDelim(dec, '{')

	for !isNextDelim(dec, '}') {
		fieldName := assertString(dec)
		switch fieldName {
		case "aliases":
			s.Aliases = make([]*Id, 0)
			assertDelim(dec, '[')
			for !isNextDelim(dec, ']') {
				newId := &Id{}
				newId.UnmarshallJSON(dec)

				s.Aliases = append(s.Aliases, newId)
			}
		case "format":
			s.Format = assertString(dec)
		default:
			s.Fields[fieldName] = assertString(dec)

		}
	}
}
