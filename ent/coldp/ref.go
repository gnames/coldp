package coldp

type Reference struct {
	ID                  string
	AlternativeID       string
	SourceID            string
	Citation            string
	Type                string
	Author              string
	AuthorID            string
	Editor              string
	EditorID            string
	Title               string
	TitleShort          string
	ContainerAuthor     string
	ContainerTitle      string
	ContainerTitleShort string
	Issued              string
	Accessed            string
	CollectionTitle     string
	CollectionEditor    string
	Volume              string
	Issue               string
	Edition             string
	Page                string
	Publisher           string
	PublisherPlace      string
	Version             string
	ISBN                string
	ISSN                string
	DOI                 string
	Link                string
	Remarks             string
}

func (r Reference) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	r.ID = row["id"]
	r.AlternativeID = row["alternative_id"]
	r.SourceID = row["source_id"]
	r.Citation = row["citation"]
	r.Type = row["type"]
	r.Author = row["author"]
	r.AuthorID = row["author_id"]
	r.Editor = row["editor"]
	r.EditorID = row["editor_id"]
	r.Title = row["title"]
	r.TitleShort = row["title_short"]
	r.ContainerAuthor = row["container_author"]
	r.ContainerTitle = row["container_title"]
	r.ContainerTitleShort = row["container_title_short"]
	r.Issued = row["issued"]
	r.Accessed = row["accessed"]
	r.CollectionTitle = row["collection_title"]
	r.CollectionEditor = row["collection_editor"]
	r.Volume = row["volume"]
	r.Issue = row["issue"]
	r.Edition = row["edition"]
	r.Page = row["page"]
	r.Publisher = row["publisher"]
	r.PublisherPlace = row["publisher_place"]
	r.Version = row["version"]
	r.ISBN = row["isbn"]
	r.ISSN = row["issn"]
	r.DOI = row["doi"]
	r.Link = row["link"]
	r.Remarks = row["remarks"]
	return r, warning
}
