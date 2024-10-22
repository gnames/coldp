package coldp

type NameRelation struct {
	NameID        string
	RelatedNameID string
	SourceID      string
	Type          string
	ReferenceID   string
	Remarks       string
}

func (n NameRelation) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	n.NameID = row["name_id"]
	n.RelatedNameID = row["related_name_id"]
	n.SourceID = row["source_id"]
	n.Type = row["type"]
	n.ReferenceID = row["reference_id"]
	n.Remarks = row["remarks"]
	return n, warning
}
