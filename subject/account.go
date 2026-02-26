package subject

import (
	"errors"

	"github.com/psychdonim/rfc9493-go/internal"
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

func (a *AccountIdentifier) FromSubjectId(s *internal.SubjectId) error {
	if s.Format != ACCOUNT_FORMAT {
		return errors.New("subject id mismatched format")
	}

	uri, ok := s.Fields[ACCOUNT_ID_NAME]
	if !ok {
		return errors.New("field not presented")
	}

	a.Uri = uri
	return nil
}
