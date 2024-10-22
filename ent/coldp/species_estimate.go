package coldp

type SpeciesEstimate struct {
	TaxonID     string
	SourceID    string
	Estimate    string
	Type        string
	ReferenceID string
	Remarks     string
}

func (s SpeciesEstimate) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	s.TaxonID = row["taxon_id"]
	s.SourceID = row["source_id"]
	s.Estimate = row["estimate"]
	s.Type = row["type"]
	s.ReferenceID = row["reference_id"]
	s.Remarks = row["remarks"]
	return s, warning
}
