package coldp

type Taxon struct {
	ID                  string
	AlternativeID       string
	SourceID            string
	ParentID            string
	Ordinal             string
	BranchLength        string
	NameID              string
	NamePhrase          string
	AccordingToID       string
	AccordingToPage     string
	AccordingToPageLink string
	Scrutinizer         string
	ScrutinizerID       string
	ScrutinizerDate     string
	BasionymID          string
	Provisional         string
	ReferenceID         string
	Extinct             string
	TemporalRangeStart  string
	TemporalRangeEnd    string
	Environment         string
	Species             string
	Section             string
	Subgenus            string
	Genus               string
	Subtribe            string
	Tribe               string
	Subfamily           string
	Family              string
	Superfamily         string
	Suborder            string
	Order               string
	Subclass            string
	Class               string
	Subphylum           string
	Phylum              string
	Kingdom             string
	Link                string
	Remarks             string
	Modified            string
	ModifiedBy          string
}

func (t Taxon) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	t.ID = row["id"]
	t.AlternativeID = row["alternativeid"]
	t.SourceID = row["sourceid"]
	t.ParentID = row["parentid"]
	t.Ordinal = row["ordinal"]
	t.BranchLength = row["branchlength"]
	t.NameID = row["nameid"]
	t.NamePhrase = row["namephrase"]
	t.AccordingToID = row["accordingtoid"]
	t.AccordingToPage = row["accordingtopage"]
	t.AccordingToPageLink = row["accordingtopagelink"]
	t.Scrutinizer = row["scrutinizer"]
	t.ScrutinizerID = row["scrutinizerid"]
	t.BasionymID = row["basionymid"]
	t.Provisional = row["Provisional"]
	t.ReferenceID = row["referenceid"]
	t.Extinct = row["extinct"]
	t.TemporalRangeStart = row["temporalrangestart"]
	t.TemporalRangeEnd = row["temporalrangeend"]
	t.Environment = row["environment"]
	t.Species = row["species"]
	t.Section = row["section"]
	t.Subgenus = row["subgenus"]
	t.Genus = row["genus"]
	t.Subtribe = row["subtribe"]
	t.Tribe = row["tribe"]
	t.Subfamily = row["subfamily"]
	t.Family = row["family"]
	t.Superfamily = row["superfamily"]
	t.Suborder = row["suborder"]
	t.Order = row["order"]
	t.Subclass = row["subclass"]
	t.Class = row["class"]
	t.Subphylum = row["subphylum"]
	t.Phylum = row["phylum"]
	t.Kingdom = row["kingdom"]
	t.ReferenceID = row["referenceid"]
	t.Link = row["link"]
	t.Remarks = row["remarks"]
	t.Modified = row["modified"]
	t.ModifiedBy = row["modifiedby"]
	return t, warning
}
