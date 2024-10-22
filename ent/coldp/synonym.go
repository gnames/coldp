package coldp

type Synonym struct {
	ID            string
	TaxonID       string
	SourceID      string
	NameID        string
	NamePhrase    string
	AccordingToID string
	Status        NameStatus
	ReferenceID   string
	Link          string
	Remarks       string
	Modified      string
	ModifiedBy    string
}

func (s Synonym) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	s.ID = row["id"]
	s.TaxonID = row["taxonid"]
	s.SourceID = row["source_id"]
	s.NameID = row["nameid"]
	s.NamePhrase = row["name_phrase"]
	s.AccordingToID = row["accordingtoid"]
	s.Status = NewNameStatus(row["status"])
	s.ReferenceID = row["referenceid"]
	s.Link = row["link"]
	s.Remarks = row["remarks"]
	s.Modified = row["modified"]
	s.ModifiedBy = row["modified_by"]
	return s, warning
}
