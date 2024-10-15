package coldp

type Synonym struct {
	ID            string
	TaxonID       string
	NameID        string
	AccordingToID string
	ReferenceID   string
	Link          string
}

func (s Synonym) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	s.ID = row["id"]
	s.TaxonID = row["taxonid"]
	s.NameID = row["nameid"]
	s.AccordingToID = row["accordingtoid"]
	s.ReferenceID = row["referenceid"]
	s.Link = row["link"]
	return s, warning
}
