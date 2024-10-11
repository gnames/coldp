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
	referenceID          string
	publishedInYear      string
	link                 string
}
