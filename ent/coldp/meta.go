package coldp

type Meta struct {
	ID              int
	DOI             string
	Title           string
	Alias           string
	Description     string
	Issued          string
	Version         string
	Keywords        []string
	GeographicScope string
	TaxonomicScope  string
	TemporalScope   string
	Confidence      int
	Completeness    int
	License         string
	URL             string
	Logo            string
	Label           string
	Citation        string
	Private         bool
	*Contact
	Editors    []Editor
	Publishers []Publisher
}

type Contact struct {
	ID int
	*Person
	*Organization
	Note string
}

type Editor struct {
	*Person
	*Organization
	Note string
}

type Contributor struct {
	*Person
	*Organization
	Note string
}

type Publisher struct {
	*Person
	*Organization
	Note string
}

type Person struct {
	ID     int
	Orcid  string
	Given  string
	Family string
	email  string
}

type Organization struct {
	ID    int
	RorID string
	Name  string
	Email string
	URL   string
}
