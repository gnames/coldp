package coldp

type Taxon struct {
	ID              string
	NameID          string
	ParentID        string
	AccordingToID   string
	Scrutinizer     string
	ScrutinizerID   string
	ScrutinizerDate string
	ReferenceID     string
	Link            string
}

func (t *Taxon) Load(headers, row []string) error {
	return nil
}
