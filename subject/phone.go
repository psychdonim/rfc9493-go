package subject

import (
	"errors"

	"github.com/psychdonim/rfc9493-go/internal"
)

const PHONE_FORMAT = "phone_number"
const PHONE_ID_NAME = "phone_number"

type PhoneNumberIdentifier struct {
	PhoneNumber string
}

func (p *PhoneNumberIdentifier) HandleWith(h Handlers) {
	h.HandlePhoneNumberIdentifier(p)
}

func (p *PhoneNumberIdentifier) Format() string {
	return PHONE_FORMAT
}

func (p *PhoneNumberIdentifier) fromSubjectId(s *internal.SubjectId) error {
	if s.Format != PHONE_FORMAT {
		return errors.New("subject id mismatched format")
	}

	phone, ok := s.Fields[PHONE_ID_NAME]
	if !ok {
		return errors.New("field iss not presented")
	}

	p.PhoneNumber = phone

	return nil
}
