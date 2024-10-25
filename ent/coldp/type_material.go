package coldp

// Type material designated to names. Type material should only be
// associated with the original name, not with a recombination.
type TypeMaterial struct {
	// ID is optional
	ID string
	// SourceID where type material is mentioned. It comes from metadata.
	SourceID string
	// NameID is the id of a name that attached to this type.
	NameID string
	// Citation is a reference where type specimen is attached to name.
	Citation string
	// Status of the type specimen.
	Status TypeStatus
	// InstitutionCode is the name or acronym in use by the institution having
	// custody of the material.
	InstitutionCode string
	// The identifier for the specimen in a collection.
	CatalogNumber string
	// ReferenceID indicating the publication of the type designation.
	ReferenceID string
	// Locality of the type. Ideally from largest area to smallest.
	Locality string
	// Country of the type locality. Preferably as ISO codes.
	Country string
	// Latitude is a decimal latitude of the type locality given in WGS84.
	Latitude string
	// Longitute is a decimal longitude of the type locality given in WGS84.
	Longitude string
	// Altitue is a decimal longitude of the type locality given in WGS84
	Altitude string
	// Host is the host organism from which the type specimen was obtained
	// (symbiotype).
	Host string
	// Sex of the type specimen.
	Sex Sex
	// Date the type material was gathered. Recommended to be given as ISO 8601
	// dates.
	Date string
	// Collector of the type specimen.
	Collector string
	// AssociatedSequences to the specimen.
	AssociatedSequences string
	// Link to further information about the specimen, e.g. as provided by
	// the institute holding the collection.
	Link string
	// Remarks about the specimen.
	Remarks string
}

// Load populates the TypeMaterial object from a row of data.
func (t TypeMaterial) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	t.ID = row["id"]
	t.SourceID = row["source_id"]
	t.NameID = row["name_id"]
	t.Citation = row["citation"]
	t.Status = NewTypeStatus(row["status"])
	t.InstitutionCode = row["institution_code"]
	t.CatalogNumber = row["catalog_number"]
	t.ReferenceID = row["reference_id"]
	t.Locality = row["locality"]
	t.Country = row["country"]
	t.Latitude = row["latitude"]
	t.Longitude = row["longitude"]
	t.Altitude = row["altitude"]
	t.Host = row["host"]
	t.Sex = NewSex(row["sex"])
	t.Date = row["date"]
	t.Collector = row["collector"]
	t.AssociatedSequences = row["associated_sequences"]
	t.Link = row["link"]
	t.Remarks = row["remarks"]
	return t, warning
}
