package coldp

type Vernacular struct {
	TaxonID         string
	Name            string
	Transliteration string
	Language        string
	Preferred       bool
	Country         string
	Area            string
	Sex             Gender
	ReferenceID     string
	Remarks         string
	Modified        string
	ModifiedBy      string
}

func (v Vernacular) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	v.TaxonID = row["taxonid"]
	v.Name = row["name"]
	v.Transliteration = row["transliteration"]
	v.Language = row["language"]
	v.Preferred = row["preferred"] == "true"
	v.Country = row["country"]
	v.Area = row["area"]
	v.Sex = NewGender(row["sex"])
	v.ReferenceID = row["referenceid"]
	v.Remarks = row["remarks"]
	v.Modified = row["modified"]
	v.ModifiedBy = row["modifiedby"]
	return v, warning
}
