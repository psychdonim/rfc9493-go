package subject

import (
	"github.com/psychdonim/rfc9493-go/internal"
	"github.com/psychdonim/rfc9493-go/internal/errors"
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
	if s.Format != ISS_SUB_FORMAT {
		return errors.NewMismatchedFormatError(ISS_SUB_FORMAT, s.Format)
	}

	iss, ok := s.Fields[ISS_ID_NAME]
	if !ok {
		return errors.NewMissedFieldError(ISS_ID_NAME)
	}

	sub, ok := s.Fields[SUB_ID_NAME]
	if !ok {
		return errors.NewMissedFieldError(SUB_ID_NAME)
	}

	is.Iss = iss
	is.Sub = sub

	return nil
}
