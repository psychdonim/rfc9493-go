package subject

import (
	"errors"

	"github.com/psychdonim/rfc9493-go/internal"
)

const ISS_SUB_FORMAT = "iss_sub"
const ISS_ID_NAME = "iss"
const SUB_ID_NAME = "sub"

type IssuerSubjectIdentifier struct {
	Iss string
	Sub string
}

func (is *IssuerSubjectIdentifier) HandleWith(h Handlers) {
	h.HandleIssuerSubjectIdentifier(is)
}

func (is *IssuerSubjectIdentifier) Format() string {
	return ISS_SUB_FORMAT
}

func (is *IssuerSubjectIdentifier) fromSubjectId(s *internal.SubjectId) error {
	if s.Format != EMAIL_FORMAT {
		return errors.New("subject id mismatched format")
	}

	iss, ok := s.Fields[ISS_ID_NAME]
	if !ok {
		return errors.New("field iss not presented")
	}

	sub, ok := s.Fields[SUB_ID_NAME]
	if !ok {
		return errors.New("field sub is not presented")
	}

	is.Iss = iss
	is.Sub = sub

	return nil
}
