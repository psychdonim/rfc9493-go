package subject

import (
	"github.com/psychdonim/rfc9493-go/internal"
	"github.com/psychdonim/rfc9493-go/internal/errors"
)

const OPAQUE_FORMAT = "opaque"
const OPAQUE_ID_NAME = "id"

type OpaqueIdentifier struct {
	Id string
}

func (o *OpaqueIdentifier) HandleWith(h Handlers) {
	h.HandleOpaqueIdentifier(o)
}

func (o *OpaqueIdentifier) Format() string {
	return OPAQUE_FORMAT
}

func (o *OpaqueIdentifier) fromSubjectId(s *internal.SubjectId) error {
	if s.Format != OPAQUE_FORMAT {
		return errors.NewMismatchedFormatError(OPAQUE_FORMAT, s.Format)
	}

	id, ok := s.Fields[OPAQUE_ID_NAME]
	if !ok {
		return errors.NewMissedFieldError(OPAQUE_ID_NAME)
	}

	o.Id = id

	return nil
}
