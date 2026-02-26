package subject

import (
	"github.com/psychdonim/rfc9493-go/internal"
	"github.com/psychdonim/rfc9493-go/internal/errors"
)

const DID_DORMAT = "did"
const DID_ID_NAME = "url"

type DecentralizedIdentifier struct {
	Url string
}

func (d *DecentralizedIdentifier) HandleWith(h Handlers) {
	h.HandleDecentralizedIdentifier(d)
}

func (d *DecentralizedIdentifier) Format() string {
	return DID_DORMAT
}

func (d *DecentralizedIdentifier) fromSubjectId(s *internal.SubjectId) error {
	if s.Format != DID_DORMAT {
		return errors.NewMismatchedFormatError(DID_DORMAT, s.Format)
	}

	url, ok := s.Fields[DID_ID_NAME]
	if !ok {
		return errors.NewMissedFieldError(DID_ID_NAME)
	}

	d.Url = url

	return nil
}
