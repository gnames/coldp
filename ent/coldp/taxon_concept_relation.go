package coldp

type TaxonConceptRelation struct {
	TaxonID        string
	RelatedTaxonID string
	SourceID       string
	Type           string
	ReferenceID    string
	Remarks        string
}

func (t TaxonConceptRelation) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	t.TaxonID = row["taxon_id"]
	t.RelatedTaxonID = row["related_taxon_id"]
	t.SourceID = row["source_id"]
	t.Type = row["type"]
	t.ReferenceID = row["reference_id"]
	t.Remarks = row["remarks"]
	return t, warning
}
