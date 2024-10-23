package coldp

// Autor represents an author of a scientific name or publication.
type Autor struct {
	// ID is the unique identifier for this author.
	ID string
	// SourceID is the identifier of the source database for this author.
	SourceID string
	// AlternativeID is an alternative identifier for this author.
	AlternativeID string
	// Given is the author's given name(s).
	Given string
	// Family is the author's family name.
	Family string
	// Suffix is any suffix associated with the author's name (e.g., Jr., III).
	Suffix string
	// AbbreviationBotany is the standard botanical abbreviation for this author.
	AbbreviationBotany string
	// AlternativeNames lists any alternative names or spellings for this author.
	AlternativeNames string
	// Sex is the author's sex.
	Sex Sex
	// Country is the author's country of origin or residence.
	Country string
	// Birth is the author's date of birth.
	Birth string
	// BirthPlace is the author's place of birth.
	BirthPlace string
	// Death is the author's date of death.
	Death string
	// Affiliation is the author's institutional affiliation.
	Affiliation string
	// Interest describes the author's research interests or areas of expertise.
	Interest string
	// ReferenceID is the identifier of a reference related to this author.
	ReferenceID string
	// Link is a link to more information about this author.
	Link string
	// Remarks are any remarks about this author.
	Remarks string
	// Modified is the date when this author's information was last modified.
	Modified string
	// ModifiedBy is the user who last modified this author's information.
	ModifiedBy string
}

// Load populates the Author data from a row of data.
func (a Autor) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	a.ID = row["id"]
	a.SourceID = row["source_id"]
	a.AlternativeID = row["alternative_id"]
	a.Given = row["given"]
	a.Family = row["family"]
	a.Suffix = row["suffix"]
	a.AbbreviationBotany = row["abbreviation_botany"]
	a.AlternativeNames = row["alternative_names"]
	a.Sex = NewSex(row["sex"])
	a.Country = row["country"]
	a.Birth = row["birth"]
	a.BirthPlace = row["birth_place"]
	a.Death = row["death"]
	a.Affiliation = row["affiliation"]
	a.Interest = row["interest"]
	a.ReferenceID = row["reference_id"]
	a.Link = row["link"]
	a.Remarks = row["remarks"]
	a.Modified = row["modified"]
	a.ModifiedBy = row["modified_by"]
	return a, warning
}