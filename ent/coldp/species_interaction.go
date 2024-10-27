package coldp

type SpeciesInteraction struct {
	// TaxonID is a reference to a taxon.
	TaxonID string

	// RelateTaxonID is a reference to the second taxon.
	RelatedTaxonID string

	// SourceID is a reference to source from metadata.
	SourceID string

	// RelatedTaxonScientificName is the name of the related taxon.
	// TODO: find if it is deprecated.
	RelatedTaxonScientificName string

	// Type of interaction.
	Type SpInteractionType

	// ReferenceID describing species interaction.
	ReferenceID string

	// Remarks about the specimen.
	Remarks string

	// Modified timestamp.
	Modified string

	// ModifiedBy contains name of a person who added/updated the record.
	ModifiedBy string
}

func (s SpeciesInteraction) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	s.TaxonID = row["taxon_id"]
	s.RelatedTaxonID = row["related_taxon_id"]
	s.SourceID = row["source_id"]
	s.RelatedTaxonScientificName = row["related_taxon_scientific_name"]
	s.Type = NewSpInteractionType(row["type"])
	s.ReferenceID = row["reference_id"]
	s.Remarks = row["remarks"]
	return s, warning
}
