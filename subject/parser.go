package subject

import (
	"github.com/psychdonim/rfc9493-go/internal"
)

func Parse(data []byte) (SubjectIdentifier, error) {
	subject := &internal.SubjectId{}
	subject.UnmarshallJSON(data)

	switch subject.Format {
	case ACCOUNT_FORMAT:
		acc := &AccountIdentifier{}
		return acc, acc.FromSubjectId(subject)

	case EMAIL_FORMAT:
		email := &EmailIdentifier{}
		return email, email.fromSubjectId(subject)

	case ISS_SUB_FORMAT:
		is := &IssuerSubjectIdentifier{}
		return is, is.fromSubjectId(subject)

	case OPAQUE_FORMAT:
		opaque := &OpaqueIdentifier{}
		return opaque, opaque.fromSubjectId(subject)

	case PHONE_FORMAT:
		phone := &PhoneNumberIdentifier{}
		return phone, phone.fromSubjectId(subject)

	case DID_DORMAT:
		did := &DecentralizedIdentifier{}
		return did, did.fromSubjectId(subject)

	case URI_FORMAT:
		uri := &UriIdentifier{}
		return uri, uri.fromSubjectId(subject)
	}

	return nil, nil
}
