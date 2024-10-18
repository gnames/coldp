package coldp

type NameUsage struct {
	ID                        string
	AlternativeID             string
	NameAlternativeID         string
	SourceID                  string
	ParentID                  string
	BasionymID                string
	Status                    string
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
	AccordingToID             string
	AccordingToPage           string
	AccordingToPageLink       string
	ReferenceID               string
	Scrutinizer               string
	ScrutinizerID             string
	ScrutinizerDate           string
	Extinct                   bool
	TemporalRangeStart        string
	TemporalRangeEnd          string
	Environment               string
	Species                   string
	Section                   string
	Subgenus                  string
	Genus                     string
	Subtribe                  string
	Tribe                     string
	Subfamily                 string
	Family                    string
	Superfamily               string
	Suborder                  string
	Order                     string
	Subclass                  string
	Class                     string
	Subphylum                 string
	Phylum                    string
	Kingdom                   string
	Ordinal                   string
	BranchLength              string
	Link                      string
	NameRemarks               string
	Remarks                   string
	Modified                  string
	ModifiedBy                string
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
