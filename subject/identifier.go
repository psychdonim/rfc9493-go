package subject

type SubjectIdentifier interface {
	HandleWith(Handlers)
	Format() string
}
