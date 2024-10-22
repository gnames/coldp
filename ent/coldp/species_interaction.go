package coldp

type SpeciesInteraction struct {
	TaxonID                    string
	Related_taxonID            string
	SourceID                   string
	RelatedTaxonScientificName string
	Type                       string
	ReferenceID                string
	Remarks                    string
}

func (s SpeciesInteraction) Load(headers, data []string) (DataLoader, error) {
	row, warning := RowToMap(headers, data)
	s.TaxonID = row["taxon_id"]
	s.Related_taxonID = row["related_taxon_id"]
	s.SourceID = row["source_id"]
	s.RelatedTaxonScientificName = row["related_taxon_scientific_name"]
	s.Type = row["type"]
	s.ReferenceID = row["reference_id"]
	s.Remarks = row["remarks"]
	return s, warning
}
