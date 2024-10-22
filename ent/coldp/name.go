package coldp

type Name struct {
	ID                        string
	AlternativeID             string
	SourceID                  string
	BasionymID                string
	ScientificName            string
	Authorship                string
	Rank                      string
	Uninomial                 string
	Genus                     string
	InfragenericEpithet       string
	SpecificEpithet           string
	InfraspecificEpithet      string
	Cultivar_epithet          string
	Notho                     NamePart
	OriginalSpelling          string
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
	Code                      NomCode
	Status                    NameStatus
	ReferenceID               string
	PublishedInYear           string
	PublishedInPage           string
	PublishedInPageLink       string
	Gender                    Gender
	GenderAgreement           bool
	Etymology                 string
	Link                      string
	Remarks                   string
	Modified                  string
	Modified_by               string
}

func (n Name) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	n.ID = row["id"]
	n.AlternativeID = row["alternative_id"]
	n.SourceID = row["source_id"]
	n.BasionymID = row["basionym_id"]
	n.ScientificName = row["scientific_name"]
	n.Authorship = row["authorship"]
	n.Rank = row["rank"]
	n.Uninomial = row["uninomial"]
	n.Genus = row["genus"]
	n.InfragenericEpithet = row["infrageneric_epithet"]
	n.SpecificEpithet = row["specific_epithet"]
	n.InfraspecificEpithet = row["infraspecific_epithet"]
	n.Cultivar_epithet = row["cultivar_epithet"]
	n.Notho = NewNamePart(row["notho"])
	n.OriginalSpelling = row["original_spelling"]
	n.CombinationAuthorship = row["combination_authorship"]
	n.CombinationAuthorshipID = row["combination_authorship_id"]
	n.CombinationExAuthorship = row["combination_ex_authorship"]
	n.CombinationExAuthorshipID = row["combination_ex_authorship_id"]
	n.CombinationAuthorshipYear = row["combination_authorship_year"]
	n.BasionymAuthorship = row["basionym_authorship"]
	n.BasionymAuthorshipID = row["basionym_authorship_id"]
	n.BasionymExAuthorship = row["basionym_ex_authorship"]
	n.BasionymExAuthorshipID = row["basionym_ex_authorship_id"]
	n.BasionymAuthorshipYear = row["basionym_authorship_year"]
	n.Code = NewNomCode(row["code"])
	n.Status = NewNameStatus(row["status"])
	n.ReferenceID = row["reference_id"]
	n.PublishedInYear = row["published_in_year"]
	n.PublishedInPage = row["published_in_page"]
	n.PublishedInPageLink = row["published_in_page_link"]
	n.Gender = NewGender(row["gender"])
	n.GenderAgreement = row["gender_agreement"] == "true"
	n.Etymology = row["etymology"]
	n.Link = row["link"]
	n.Remarks = row["remarks"]
	n.Modified = row["modified"]
	n.Modified_by = row["modified_by"]
	return n, warning
}
