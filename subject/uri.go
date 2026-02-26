package subject

import (
	"github.com/psychdonim/rfc9493-go/internal"
	"github.com/psychdonim/rfc9493-go/internal/errors"
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
		return errors.NewMismatchedFormatError(URI_FORMAT, s.Format)
	}

	uri, ok := s.Fields[URI_ID_NAME]
	if !ok {
		return errors.NewMissedFieldError(URI_ID_NAME)
	}

	u.Uri = uri

	return nil
}
