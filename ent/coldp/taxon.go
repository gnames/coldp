package coldp

type Taxon struct {
	ID              string
	NameID          string
	ParentID        string
	AccordingToID   string
	Scrutinizer     string
	ScrutinizerID   string
	ScrutinizerDate string
	ReferenceID     string
	Link            string
}

func (t Taxon) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	t.ID = row["id"]
	t.NameID = row["nameid"]
	t.ParentID = row["parentid"]
	t.AccordingToID = row["accordingtoid"]
	t.Scrutinizer = row["scrutinizer"]
	t.ScrutinizerID = row["scrutinizerid"]
	t.ReferenceID = row["referenceid"]
	t.Link = row["link"]
	return t, warning
}
