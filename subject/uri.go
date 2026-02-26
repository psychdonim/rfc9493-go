package subject

import (
	"errors"

	"github.com/psychdonim/rfc9493-go/internal"
)

const URI_FORMAT = "uri"
const URI_ID_NAME = "uri"

type UriIdentifier struct {
	Uri string
}

func (u *UriIdentifier) HandleWith(h Handlers) {
	h.HandleUriIdentifier(u)
}

func (u *UriIdentifier) Format() string {
	return URI_FORMAT
}

func (u *UriIdentifier) fromSubjectId(s *internal.SubjectId) error {
	if s.Format != URI_FORMAT {
		return errors.New("subject id mismatched format")
	}

	uri, ok := s.Fields[URI_ID_NAME]
	if !ok {
		return errors.New("field iss not presented")
	}

	u.Uri = uri

	return nil
}
