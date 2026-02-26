package internal

import (
	"bytes"
	"encoding/json"
)

type SubjectId struct {
	id
	Aliases []*id
}

func (s *SubjectId) UnmarshallJSON(data []byte) {
	dec := json.NewDecoder(bytes.NewReader(data))

	assertDelim(dec, '{')

	for !isNextDelim(dec, '}') {
		fieldName := assertString(dec)
		switch fieldName {
		case "aliases":
			s.Aliases = make([]*id, 0)
			assertDelim(dec, '[')
			for !isNextDelim(dec, ']') {
				newId := &id{}
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
