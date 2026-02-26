package subject

import (
	"errors"

	"github.com/psychdonim/rfc9493-go/internal"
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
		return errors.New("subject id mismatched format")
	}

	id, ok := s.Fields[OPAQUE_ID_NAME]
	if !ok {
		return errors.New("field iss not presented")
	}

	o.Id = id

	return nil
}
