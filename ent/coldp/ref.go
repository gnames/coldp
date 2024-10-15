package coldp

type Reference struct {
	ID       string
	Citation string
	Link     string
	DOI      string
	Remarks  string
}

func (r Reference) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	r.ID = row["id"]
	r.Citation = row["citation"]
	r.Link = row["link"]
	r.DOI = row["doi"]
	r.Remarks = row["remarks"]
	return r, warning
}
