package subject

import (
	"testing"

	"github.com/psychdonim/rfc9493-go/internal"
)

var acc = &AccountIdentifier{
	Uri: "test_uri",
}

const testUri = "uri_test"

func TestAccount_Format(t *testing.T) {
	if acc.Format() != ACCOUNT_FORMAT {
		t.Error("wrong account format")
	}
}

func TestAccount_FromSubjectId_ValidSubjectId(t *testing.T) {
	sub := &internal.SubjectId{
		Id: internal.Id{
			Format: ACCOUNT_FORMAT,
			Fields: map[string]string{
				ACCOUNT_ID_NAME: testUri,
			},
		},
	}

	acc := &AccountIdentifier{}
	err := acc.fromSubjectId(sub)

	if err != nil {
		t.Fail()
	}

	if acc.Uri != testUri {
		t.Error("wrong uri")
	}
}

func TestAccount_FromSubjectId_InvalidFormat(t *testing.T) {
	sub := &internal.SubjectId{
		Id: internal.Id{
			Format: EMAIL_FORMAT,
			Fields: map[string]string{
				ACCOUNT_ID_NAME: testUri,
			},
		},
	}

	acc := &AccountIdentifier{}
	err := acc.fromSubjectId(sub)

	if err == nil {
		t.Error("error expected cause of wrong format")
	}
}

func TestAccount_FromSubjectId_MissingField(t *testing.T) {
	sub := &internal.SubjectId{
		Id: internal.Id{
			Format: ACCOUNT_FORMAT,
			Fields: map[string]string{
				EMAIL_FORMAT: testUri,
			},
		},
	}

	acc := &AccountIdentifier{}
	err := acc.fromSubjectId(sub)

	if err == nil {
		t.Error("error expected cause of missing field")
	}
}
