package coldp

type Synonym struct {
	// Optional unique identifier for the synonym. If given it should not clash
	// with the taxon ids.
	ID string
	// TaxonID is the unique identifier for the related taxon.
	TaxonID string
	// SourceID is the identifier of the source from metadata.
	SourceID string
	// NameID is the identifier of the name associated with this taxon.
	NameID string
	// NamePhrase is an optional annotation attached to the name in this
	// context (eg `sensu lato` etc).
	NamePhrase string
	// AccordingToID is ReferenceID of the source that this taxon is based on.
	AccordingToID string

	Status      TaxonomicStatus
	ReferenceID string
	Link        string
	Remarks     string
	Modified    string
	ModifiedBy  string
}

func (s Synonym) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	s.ID = row["id"]
	s.TaxonID = row["taxonid"]
	s.SourceID = row["source_id"]
	s.NameID = row["nameid"]
	s.NamePhrase = row["name_phrase"]
	s.AccordingToID = row["accordingtoid"]
	s.Status = NewNomStatus(row["status"])
	s.ReferenceID = row["referenceid"]
	s.Link = row["link"]
	s.Remarks = row["remarks"]
	s.Modified = row["modified"]
	s.ModifiedBy = row["modified_by"]
	return s, warning
}
