package coldp

type Media struct {
	TaxonID  string
	SourceID string
	Url      string
	Type     string
	Format   string
	Title    string
	Created  string
	Creator  string
	License  string
	Link     string
	Remarks  string
}

func (m Media) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	m.TaxonID = row["taxon_id"]
	m.SourceID = row["source_id"]
	m.Url = row["url"]
	m.Type = row["type"]
	m.Format = row["format"]
	m.Title = row["title"]
	m.Created = row["created"]
	m.Creator = row["creator"]
	m.License = row["license"]
	m.Link = row["link"]
	m.Remarks = row["remarks"]
	return m, warning
}
