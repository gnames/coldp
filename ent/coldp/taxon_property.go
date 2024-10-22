package coldp

type TaxonProperty struct {
	TaxonID     string
	SourceID    string
	Property    string
	Value       string
	ReferenceID string
	Page        string
	Ordinal     int
	Remarks     string
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
	return t, warning
}
