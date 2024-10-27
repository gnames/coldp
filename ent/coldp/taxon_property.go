package coldp

// TaxonProperty allows to add arbitrary properties that further describe
// a taxon.
type TaxonProperty struct {
	// TaxonID is a reference to a taxon.
	TaxonID string

	// SourceID is a reference to source from metadata.
	SourceID string

	// Property type according to some controlled vocabulary.
	Property string

	// Value of the property.
	Value string

	// ReferenceID points to the publication about TaxonProperty.
	ReferenceID string

	// Page where this property is given.
	Page string

	// Ordinal for sorting.
	Ordinal int

	// Remarks about the specimen.
	Remarks string

	// Modified timestamp.
	Modified string

	// ModifiedBy contains name of a person who added/updated the record.
	ModifiedBy string
}

func (t TaxonProperty) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	t.TaxonID = row["taxon_id"]
	t.SourceID = row["source_id"]
	t.Property = row["property"]
	t.Value = row["value"]
	t.ReferenceID = row["reference_id"]
	t.Page = row["page"]
	t.Ordinal = ToInt(row["ordinal"])
	t.Remarks = row["remarks"]
	t.Modified = row["modified"]
	t.ModifiedBy = row["modified_by"]
	return t, warning
}
