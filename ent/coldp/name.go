package coldp

// Name contains information about a scientific name.
type Name struct {
	// ID is the unique identifier for this name.
	ID string

	// AlternativeID can include several alternative identifiers for this name.
	// IDs can be with or without scope: col:123, gbif:uuid, plazi:int.
	AlternativeID string

	// SourceID is the identifier of the source for this name (taken from
	// metadata).
	SourceID string

	// BasionymID is the identifier of the basionym for this name.
	// Basionym meaning is broad, and can include any terminal name (name without
	// parentage).
	BasionymID string

	// ScientificName is a full canonical form (without authorship) of
	// the scientific name.
	ScientificName string

	// Authorship is the verbatim authorship of the scientific name.
	Authorship string

	// FullScientificName contains the most complete version of scientific name
	// available (with authorship, subgenus, intermediate authors, hybrid
	// signs etc).
	// It is not part of CoLDP standard.
	FullScientificName string

	// Rank is the taxonomic rank of the name.
	Rank Rank

	// Uninomial is the uninomial for names at the rank of genus or higher.
	Uninomial string

	// Genus is the genus name.
	Genus string

	// InfragenericEpithet of the name.
	InfragenericEpithet string

	// SpecificEpithet of the name.
	SpecificEpithet string

	// InfraspecificEpithet of the name.
	InfraspecificEpithet string

	// CultivarEpithet is the cultivar epithet.
	CultivarEpithet string

	// Notho is the nothotaxon part of the name.
	Notho NamePart

	// OriginalSpelling indicates whether the name is spelled as originally
	// published.
	OriginalSpelling bool

	// CombinationAuthorship is the authorship of the combination.
	// Authors are separated by '|'.
	CombinationAuthorship string

	// CombinationAuthorshipID is the identifier of the combination authorship.
	// Authors are separated by '|'.
	CombinationAuthorshipID string

	// CombinationExAuthorship is the ex authorship of the combination.
	// Authors are separated by '|'.
	CombinationExAuthorship string

	// CombinationExAuthorshipID is the identifier of the combination ex
	// authorship. Authors are separated by '|'.
	CombinationExAuthorshipID string

	// CombinationAuthorshipYear is the year of the combination authorship.
	// Authors are separated by '|'.
	CombinationAuthorshipYear string

	// BasionymAuthorship is the authorship of the basionym.
	// Authors are separated by '|'.
	BasionymAuthorship string

	// BasionymAuthorshipID is the identifier of the basionym authorship.
	// Authors are separated by '|'.
	BasionymAuthorshipID string

	// BasionymExAuthorship is the ex authorship of the basionym.
	// Authors are separated by '|'.
	BasionymExAuthorship string

	// BasionymExAuthorshipID is the identifier of the basionym ex authorship.
	// Authors are separated by '|'.
	BasionymExAuthorshipID string

	// BasionymAuthorshipYear is the year of the basionym authorship.
	// Authors are separated by '|'.
	BasionymAuthorshipYear string

	// Code is the nomenclatural code that applies to this name.
	Code NomCode

	// Status is the nomenclatural status of the name.
	Status NomStatus

	// ReferenceID is the identifier of the reference in which the name was
	// published.
	ReferenceID string

	// PublishedInYear is the year in which the name was published.
	PublishedInYear string

	// PublishedInPage is the page in which the name was published.
	PublishedInPage string

	// PublishedInPageLink is the link to the page in which the name was
	// published.
	PublishedInPageLink string

	// Gender is the grammatical gender of the name.
	Gender Gender

	// GenderAgreement indicates whether the gender of the name agrees with the
	// gender of the genus.
	GenderAgreement bool

	// Etymology is the etymology of the name.
	Etymology string

	// Link is a URL to more information about the name.
	Link string

	// Remarks are any remarks/notes about the name.
	Remarks string

	// Modified is the date when the name was last modified.
	Modified string

	// ModifiedBy is the user who last modified the name.
	ModifiedBy string
}

// Load processes a slice of strings into Name object using
// field names from headers.
func (n Name) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	n.ID = row["id"]
	n.AlternativeID = row["alternativeid"]
	n.SourceID = row["sourceid"]
	n.BasionymID = row["basionymid"] // becomes NameRelation
	n.ScientificName = row["scientificname"]
	n.Authorship = row["authorship"] // verbatim author string
	n.Rank = NewRank(row["rank"])
	n.Uninomial = row["uninomial"]
	n.Genus = row["genus"]
	n.InfragenericEpithet = row["infragenericepithet"]
	n.SpecificEpithet = row["specificepithet"]
	n.InfraspecificEpithet = row["infraspecificepithet"]
	n.CultivarEpithet = row["cultivarepithet"]
	n.Notho = NewNamePart(row["notho"])
	n.OriginalSpelling = ToBool(row["originalspelling"])
	n.CombinationAuthorship = row["combinationauthorship"]
	n.CombinationAuthorshipID = row["combinationauthorshipid"]
	n.CombinationExAuthorship = row["combinationexauthorship"]
	n.CombinationExAuthorshipID = row["combinationexauthorshipid"]
	n.CombinationAuthorshipYear = row["combinationauthorshipyear"]
	n.BasionymAuthorship = row["basionymauthorship"]
	n.BasionymAuthorshipID = row["basionymauthorshipid"]
	n.BasionymExAuthorship = row["basionymexauthorship"]
	n.BasionymExAuthorshipID = row["basionymexauthorshipid"]
	n.BasionymAuthorshipYear = row["basionymauthorshipyear"]
	n.Code = NewNomCode(row["code"])
	n.Status = NewNomStatus(row["status"])
	n.ReferenceID = row["referenceid"]
	n.PublishedInYear = row["publishedinyear"]
	n.PublishedInPage = row["publishedinpage"]
	n.PublishedInPageLink = row["publishedinpagelink"]
	n.Gender = NewGender(row["gender"])
	n.GenderAgreement = ToBool(row["genderagreement"])
	n.Etymology = row["etymology"]
	n.Link = row["link"]
	n.Remarks = row["remarks"]
	n.Modified = row["modified"]
	n.ModifiedBy = row["modifiedby"]
	return n, warning
}
