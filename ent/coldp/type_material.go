package coldp

type TypeMaterial struct {
	ID                  string
	SourceID            string
	NameID              string
	Citation            string
	Status              string
	InstitutionCode     string
	CatalogNumber       string
	ReferenceID         string
	Locality            string
	Country             string
	Latitude            string
	Longitude           string
	Altitude            string
	Host                string
	Sex                 Gender
	Date                string
	Collector           string
	AssociatedSequences string
	Link                string
	Remarks             string
}

func (t TypeMaterial) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	t.ID = row["id"]
	t.SourceID = row["source_id"]
	t.NameID = row["name_id"]
	t.Citation = row["citation"]
	t.Status = row["status"]
	t.InstitutionCode = row["institution_code"]
	t.CatalogNumber = row["catalog_number"]
	t.ReferenceID = row["reference_id"]
	t.Locality = row["locality"]
	t.Country = row["country"]
	t.Latitude = row["latitude"]
	t.Longitude = row["longitude"]
	t.Altitude = row["altitude"]
	t.Host = row["host"]
	t.Sex = NewGender(row["sex"])
	t.Date = row["date"]
	t.Collector = row["collector"]
	t.AssociatedSequences = row["associated_sequences"]
	t.Link = row["link"]
	t.Remarks = row["remarks"]
	return t, warning
}
