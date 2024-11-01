package custom

type Custom struct {
	DomainName string
}

func New(domainName string) Custom {
	return Custom{
		DomainName: domainName,
	}
}
