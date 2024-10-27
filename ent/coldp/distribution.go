package coldp

// Distribution of a taxon in a given area.
type Distribution struct {
	// TaxonID for the distribution.
	TaxonID string

	// Source (from metadata) for distrubution data.
	SourceID string

	// Area is the geographic area this distribution record is about. The value
	// provides a human label for the area specified by areaID. Free text values
	// can be provided here when the gazetteer is set to TEXT
	Area string

	// AreaID is the identifier/code for the geographic area this distribution
	// record is about. The value must be taken from the gazetteer this record
	// declares. E.g. country codes, TDWG codes or TEOW identifiers. If the TEXT
	// gazetteer is used only the free text area should be given with no areaID.
	AreaID string

	// Gezetteer is the geographic gazetteer the area is defined in. If none is
	// given defaults to free TEXT.
	Gazetteer GazetteerEnt

	// Status of the distribution.
	Status DistrStatus

	// ReferenceID to the corresponding reference.
	ReferenceID string

	// Remarks are notes to the data.
	Remarks string

	// Modified timestamp.
	Modified string

	// ModifiedBy contains name of a person who added/updated the record.
	ModifiedBy string
}

func (d Distribution) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	d.TaxonID = row["taxon_id"]
	d.SourceID = row["source_id"]
	d.Area = row["area"]
	d.AreaID = row["area_id"]
	d.Gazetteer = NewGazetteerEnt(row["gazetteer"])
	d.Status = NewDistrStatus(row["status"])
	d.ReferenceID = row["reference_id"]
	d.Remarks = row["remarks"]
	d.Modified = row["modified"]
	d.ModifiedBy = row["modified_by"]
	return d, warning
}
