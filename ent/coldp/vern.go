package coldp

// Vernacular name of a taxon.
type Vernacular struct {
	// TaxonID of the corresponding taxon.
	TaxonID string

	// SourceID to the source from metadata.
	SourceID string

	// Name (vernacular).
	Name string

	// Transliteration to latin alphabet.
	Transliteration string

	// Langugage of the name (three-letter ISO 639-3 code).
	Language string

	// Preferred is true when this vernacular names is considered the most
	// used for the language.
	Preferred bool

	// Country name (two-letter ISO 3166-2  code)
	Country string

	// Area (optional) describing the geographic use of the vernacular name in
	// free text within the given country.
	Area string

	// Sex (ptional) of the organism this vernacular name is restricted to.
	Sex Sex

	// ReferenceID where the name is supported.
	ReferenceID string

	// Remarks for the name.
	Remarks string

	// Modified is a timestamp.
	Modified string

	// ModifiedBy a person who last modified the record.
	ModifiedBy string
}

// Load populates a Vernacular object from a row of data.
func (v Vernacular) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	v.TaxonID = row["taxonid"]
	v.SourceID = row["sourceid"]
	v.Name = row["name"]
	v.Transliteration = row["transliteration"]
	v.Language = row["language"]
	v.Preferred = row["preferred"] == "true"
	v.Country = row["country"]
	v.Area = row["area"]
	v.Sex = NewSex(row["sex"])
	v.ReferenceID = row["referenceid"]
	v.Remarks = row["remarks"]
	v.Modified = row["modified"]
	v.ModifiedBy = row["modifiedby"]
	return v, warning
}
