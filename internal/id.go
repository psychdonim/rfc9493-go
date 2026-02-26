package internal

import "encoding/json"

type Id struct {
	Format string
	Fields map[string]string
}

func (i *Id) UnmarshallJSON(dec *json.Decoder) {
	assertDelim(dec, '{')

	for !isNextDelim(dec, '}') {
		fieldName := assertString(dec)
		switch fieldName {
		case "format":
			i.Format = assertString(dec)
		case "aliases":
			// return err
		default:
			i.Fields[fieldName] = assertString(dec)
		}
	}
}
