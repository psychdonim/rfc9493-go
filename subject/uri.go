package subject

import (
	"github.com/psychdonim/rfc9493-go/internal"
	"github.com/psychdonim/rfc9493-go/internal/errors"
)

// URI format identifier as defined in [RFC 9493 §3.2.7].
//
// [RFC 9493 §3.2.7]: https://www.rfc-editor.org/rfc/rfc9493.html#name-uniform-resource-identifier
const URI_FORMAT = "uri"

// Field name used for URI-based subject identification as defined in [RFC 9493 §3.2.7].
//
// [RFC 9493 §3.2.7]: https://www.rfc-editor.org/rfc/rfc9493.html#name-uniform-resource-identifier
const URI_ID_NAME = "uri"

// UriIdentifier represents a subject identified by a URI
// as defined in [RFC 9493 §3.2.7].
//
// Implements the [SubjectIdentifier] interface.
//
// [RFC 9493 §3.2.7]: https://www.rfc-editor.org/rfc/rfc9493.html#name-uniform-resource-identifier
type UriIdentifier struct {
	Uri string
}

// HandleWith dispatches the identifier to the provided [Handlers].
func (u *UriIdentifier) HandleWith(h Handlers) {
	h.HandleUriIdentifier(u)
}

// Format returns the identifier format ("uri").
func (u *UriIdentifier) Format() string {
	return URI_FORMAT
}

// FromSubjectID populates the identifier from [internal.SubjectID].
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
