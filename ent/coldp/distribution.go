package coldp

type Distribution struct {
	TaxonID     string
	SourceID    string
	Area        string
	AreaID      string
	Gazetteer   string
	Status      string
	ReferenceID string
	Remarks     string
}

func (d Distribution) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	d.TaxonID = row["taxon_id"]
	d.SourceID = row["source_id"]
	d.Area = row["area"]
	d.AreaID = row["area_id"]
	d.Gazetteer = row["gazetteer"]
	d.Status = row["status"]
	d.ReferenceID = row["reference_id"]
	d.Remarks = row["remarks"]
	return d, warning
}
