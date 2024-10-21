package coldp

type NameUsage struct {
	ID                        string // t
	AlternativeID             string // t
	NameAlternativeID         string
	SourceID                  string // t
	ParentID                  string // t
	BasionymID                string // t
	Status                    string // t
	ScientificName            string
	Authorship                string
	Rank                      string
	Notho                     string
	OriginalSpelling          bool
	Uninomial                 string
	GenericName               string
	InfragenericEpithet       string
	SpecificEpithet           string
	InfraspecificEpithet      string
	CultivarEpithet           string
	CombinationAuthorship     string
	CombinationAuthorshipID   string
	CombinationExAuthorship   string
	CombinationExAuthorshipID string
	CombinationAuthorshipYear string
	BasionymAuthorship        string
	BasionymAuthorshipID      string
	BasionymExAuthorship      string
	BasionymExAuthorshipID    string
	BasionymAuthorshipYear    string
	NamePhrase                string
	NameReferenceID           string
	PublishedInYear           string
	PublishedInPage           string
	PublishedInPageLink       string
	Gender                    string
	GenderAgreement           string
	Etymology                 string
	Code                      NomCode
	NameStatus                NameStatus
	AccordingToID             string // t
	AccordingToPage           string // t
	AccordingToPageLink       string // t
	ReferenceID               string // t
	Scrutinizer               string // t
	ScrutinizerID             string // t
	ScrutinizerDate           string // t
	Extinct                   bool   // t
	TemporalRangeStart        string // t
	TemporalRangeEnd          string // t
	Environment               string // t
	Species                   string // t
	Section                   string // t
	Subgenus                  string // t
	Genus                     string // t
	Subtribe                  string // t
	Tribe                     string // t
	Subfamily                 string // t
	Family                    string // t
	Superfamily               string // t
	Suborder                  string // t
	Order                     string // t
	Subclass                  string // t
	Class                     string // t
	Subphylum                 string // t
	Phylum                    string // t
	Kingdom                   string // t
	Ordinal                   string // t
	BranchLength              string // t
	Link                      string // t
	NameRemarks               string
	Remarks                   string // t
	Modified                  string // t
	ModifiedBy                string // t
}

func (n NameUsage) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	n.ID = row[""]
	n.AlternativeID = row["alternativeid"]
	n.NameAlternativeID = row["namealternativeid"]
	n.SourceID = row["sourceid"]
	n.ParentID = row["parentid"]
	n.BasionymID = row["basionymid"]
	n.Status = row["status"]
	n.ScientificName = row["scientificname"]
	n.Authorship = row["authorship"]
	n.Rank = row["rank"]
	n.Notho = row["notho"]
	n.OriginalSpelling = ToBool(row["originalspelling"])
	n.Uninomial = row["uninomial"]
	n.GenericName = row["genericname"]
	n.InfragenericEpithet = row["infragenericepithet"]
	n.SpecificEpithet = row["specificepithet"]
	n.InfraspecificEpithet = row["infraspecificepithet"]
	n.CultivarEpithet = row["cultivarepithet"]
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
	n.NamePhrase = row["namephrase"]
	n.NameReferenceID = row["namereferenceid"]
	n.PublishedInYear = row["publishedinyear"]
	n.PublishedInPage = row["publishedinpage"]
	n.PublishedInPageLink = row["publishedinpagelink"]
	n.Gender = row["gender"]
	n.GenderAgreement = row["genderagreement"]
	n.Etymology = row["etymology"]
	n.Code = NewNomCode(row["code"])
	n.NameStatus = NewNameStatus(row["namestatus"])
	n.AccordingToID = row["accordingtoid"]
	n.AccordingToPage = row["accordingtopage"]
	n.AccordingToPageLink = row["accordingtopagelink"]
	n.ReferenceID = row["referenceid"]
	n.Scrutinizer = row["scrutinizer"]
	n.ScrutinizerID = row["scrutinizerid"]
	n.ScrutinizerDate = row["scrutinizerdate"]
	n.Extinct = ToBool(row["extinct"])
	n.TemporalRangeStart = row["temporalrangestart"]
	n.TemporalRangeEnd = row["temporalrangeend"]
	n.Environment = row["environment"]
	n.Species = row["species"]
	n.Section = row["section"]
	n.Subgenus = row["subgenus"]
	n.Genus = row["genus"]
	n.Subtribe = row["subtribe"]
	n.Tribe = row["tribe"]
	n.Subfamily = row["subfamily"]
	n.Family = row["family"]
	n.Superfamily = row["superfamily"]
	n.Suborder = row["suborder"]
	n.Order = row["order"]
	n.Subclass = row["subclass"]
	n.Class = row["class"]
	n.Subphylum = row["subphylum"]
	n.Phylum = row["phylum"]
	n.Kingdom = row["kingdom"]
	n.Ordinal = row["ordinal"]
	n.BranchLength = row["branchlength"]
	n.Link = row["link"]
	n.NameRemarks = row["nameremarks"]
	n.Remarks = row["remarks"]
	n.Modified = row["modified"]
	n.ModifiedBy = row["modifiedby"]
	return n, warning
}
