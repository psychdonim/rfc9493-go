package subject

import (
	"errors"

	"github.com/psychdonim/rfc9493-go/internal"
)

const EMAIL_FORMAT = "email"
const EMAIL_ID_NAME = "email"

type EmailIdentifier struct {
	Email string
}

func (e *EmailIdentifier) HandleWith(h Handlers) {
	h.HandleEmailIdentifier(e)
}

func (e *EmailIdentifier) Format() string {
	return EMAIL_FORMAT
}

func (e *EmailIdentifier) fromSubjectId(s *internal.SubjectId) error {
	if s.Format != EMAIL_FORMAT {
		return errors.New("subject id mismatched format")
	}

	email, ok := s.Fields[EMAIL_ID_NAME]
	if !ok {
		return errors.New("field not presented")
	}

	e.Email = email

	return nil
}
