package coldp

type TaxonConceptRelation struct {
	// TaxonID is a reference to a taxon.
	TaxonID string

	// RelateTaxonID is a reference to the second taxon.
	RelatedTaxonID string

	// SourceID is a reference to source from metadata.
	SourceID string

	Type TaxonConceptRelType

	// ReferenceID to description of TaxonConceptRelation.
	ReferenceID string

	// Remarks about the specimen.
	Remarks string

	// Modified timestamp.
	Modified string

	// ModifiedBy contains name of a person who added/updated the record.
	ModifiedBy string
}

func (t TaxonConceptRelation) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	t.TaxonID = row["taxon_id"]
	t.RelatedTaxonID = row["related_taxon_id"]
	t.SourceID = row["source_id"]
	t.Type = NewTaxonConceptRelType(row["type"])
	t.ReferenceID = row["reference_id"]
	t.Remarks = row["remarks"]
	return t, warning
}
