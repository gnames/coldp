package coldp

type Name struct {
	ID                   string
	AlternativeID        string
	BasionymID           string
	ScientificName       string
	Authorship           string
	Rank                 string
	Uninomial            string
	Genus                string
	InfragenericEpithet  string
	SpecificEpithet      string
	InfraspecificEpithet string
	Code                 NomCode
	ReferenceID          string
	PublishedInYear      string
	Link                 string
}

func (n *Name) Load(headers, data []string) error {
	row, warning := RowToMap(headers, data)
	n.ID = row["id"]
	n.AlternativeID = row["alternativeid"]
	n.BasionymID = row["basionymid"]
	n.ScientificName = row["scientificname"]
	n.Authorship = row["authorship"]
	n.Rank = row["rank"]
	n.Uninomial = row["uninomial"]
	n.Genus = row["genus"]
	n.InfragenericEpithet = row["infraspecificepithet"]
	n.SpecificEpithet = row["specificepithet"]
	n.InfraspecificEpithet = row["infraspecificepithet"]
	n.Code = NewNomCode(row["code"])
	n.ReferenceID = row["referenceid"]
	n.PublishedInYear = row["publishedinyear"]
	n.Link = row["link"]
	return warning
}
