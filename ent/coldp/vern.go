package coldp

type Vernacular struct {
	TaxonID         string
	Name            string
	Transliteration string
	Language        string
	Preferred       bool
	Country         string
	Area            string
	Sex             Gender
	ReferenceID     string
	Remarks         string
	Modified        string
	ModifiedBy      string
}
