package subject

import (
	"github.com/psychdonim/rfc9493-go/internal"
	"github.com/psychdonim/rfc9493-go/internal/errors"
)

const ACCOUNT_FORMAT = "account"
const ACCOUNT_ID_NAME = "uri"

type AccountIdentifier struct {
	Uri string
}

func (a *AccountIdentifier) HandleWith(handlers Handlers) {
	handlers.HandleAccountIdentifier(a)
}

func (a *AccountIdentifier) Format() string {
	return ACCOUNT_FORMAT
}

func (a *AccountIdentifier) fromSubjectId(s *internal.SubjectId) error {
	if s.Format != ACCOUNT_FORMAT {
		return errors.NewMismatchedFormatError(ACCOUNT_FORMAT, s.Format)
	}

	uri, ok := s.Fields[ACCOUNT_ID_NAME]
	if !ok {
		return errors.NewMissedFieldError(ACCOUNT_ID_NAME)
	}

	a.Uri = uri
	return nil
}
