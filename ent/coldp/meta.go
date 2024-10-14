package coldp

type Meta struct {
	Key               string   `yaml:"key,omitempty"`
	Title             string   `yaml:"title"`
	Alias             string   `yaml:"alias,omitempty"`
	Description       string   `yaml:"description,omitempty"`
	DOI               string   `yaml:"doi,omitempty"`
	Issued            string   `yaml:"issued,omitempty"`
	Version           string   `yaml:"version,omitempty"`
	GeographicScope   string   `yaml:"geographicScope,omitempty"`
	TaxonomicScope    string   `yaml:"taxonomicScope,omitempty"`
	TemporalScope     string   `yaml:"temporalScope,omitempty"`
	Keywords          []string `yaml:"keywords,omitempty"`
	Confidence        int      `yaml:"confidence,omitempty"`
	Completeness      int      `yaml:"completeness,omitempty"`
	License           string   `yaml:"license,omitempty"`
	URL               string   `yaml:"url,omitempty"`
	Logo              string   `yaml:"logo,omitempty"`
	Label             string   `yaml:"label,omitempty"`
	Citation          string   `yaml:"citation,omitempty"`
	LastImportAttempt string   `yaml:"lastImportAttempt,omitempty"`
	LastImportState   string   `yaml:"lastImportState,omitempty"`
	Private           bool     `yaml:"private,omitempty"`
	Contact           *Actor   `yaml:"contact,omitempty"`
	Publisher         *Actor   `yaml:"publisher,omitempty"`
	Editors           []Actor  `yaml:"editor,omitempty"`
	Creators          []Actor  `yaml:"creator,omitempty"`
	Contributors      []Actor  `yaml:"contributor,omitempty"`
	Sources           []Source `yaml:"source,omitempty"`
}

type Actor struct {
	Orcid        string `yaml:"orcid,omitempty"`
	Given        string `yaml:"given,omitempty"`
	Family       string `yaml:"family,omitempty"`
	Email        string `yaml:"email,omitempty"`
	State        string `yaml:"state,omitempty"`
	Country      string `yaml:"country,omitempty"`
	RorID        string `yaml:"rorid,omitempty"`
	Organization string `yaml:"organisation,omitempty"`
	URL          string `yaml:"url,omitempty"`
	Note         string `yaml:"note,omitempty"`
}

type Source struct {
	ID      string `yaml:"id,omitempty"`
	Type    string `yaml:"type,omitempty"`
	Title   string `yaml:"title,omitempty"`
	Authors []any  `yaml:"author,omitempty"`
	Issued  string `yaml:"issued,omitempty"`
	ISBN    string `yaml:"isbn,omitempty"`
}
