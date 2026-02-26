package subject

type Handlers interface {
	HandleAccountIdentifier(*AccountIdentifier)
	HandleEmailIdentifier(*EmailIdentifier)
	HandleIssuerSubjectIdentifier(*IssuerSubjectIdentifier)
	HandleOpaqueIdentifier(*OpaqueIdentifier)
	HandlePhoneNumberIdentifier(*PhoneNumberIdentifier)
	HandleDecentralizedIdentifier(*DecentralizedIdentifier)
	HandleUriIdentifier(*UriIdentifier)
}
