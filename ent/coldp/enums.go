package coldp

import "strings"

// NomCode provides types of nomenclatural Codes
type NomCode int

const (
	UnknownCode NomCode = iota
	Bacterial
	Botanical
	Cultivars
	PhytoSociological
	Virus
	Zoological
)

// NewNomCode converts a string (number or word) to NomCode
func NewNomCode(s string) NomCode {
	s = strings.ToLower(s)
	switch s {
	case "1", "bacterial", "icnp":
		return Bacterial
	case "2", "botanical", "icn", "icnafp", "icbn":
		return Botanical
	case "3", "cultivars", "icncp":
		return Cultivars
	case "4", "phytosociological", "icpn":
		return PhytoSociological
	case "5", "virus", "icvcn":
		return Virus
	case "6", "zoological", "iczn":
		return Zoological
	default:
		return UnknownCode
	}
}

// ArchiveType provides type of CoLDP. Can be 'flat' or NameUsage, and
// 'original' or Name.
type ArchiveType int

const (
	UnknownAT ArchiveType = iota
	NameAT
	NameUsageAT
)

type Gender int

const (
	UnknownG Gender = iota
	Masculine
	Feminine
	Neutral
)

func NewGender(s string) Gender {
	s = strings.ToLower(s)
	switch s {
	case "1", "m", "masculine":
		return Masculine
	case "2", "f", "feminine":
		return Feminine
	case "3", "n", "neutral":
		return Neutral
	default:
		return UnknownG
	}
}

type FileType int

const (
	UnknownFileType FileType = iota
	JSON
	YAML
	CSV
	TSV
	// pipe-separated file
	PSV
	// json-based Citation Style Language
	JSONCSL
	BIBTEX
)

// DataType provides types of data files in CoLDP archive.
type DataType int

const (
	UnkownDT DataType = iota
	AuthorDT
	NameRelationDT
	TaxonDT
	SynonymDT
	NameUsageDT
	NameDT
	TaxonPropertyDT
	TaxonConceptRelationDT
	SpeciesInteractionDT
	SpeciesEstimateDT
	ReferenceDT
	TypeMaterialDT
	DistributionDT
	MediaDT
	VernacularNameDT
	TreatmentDT
)

func (dt DataType) FileFormats() []FileType {
	switch dt {
	case AuthorDT, NameRelationDT, TaxonDT, SynonymDT, NameUsageDT,
		NameDT, TaxonPropertyDT, TaxonConceptRelationDT, SpeciesInteractionDT,
		SpeciesEstimateDT, TypeMaterialDT, DistributionDT, MediaDT,
		VernacularNameDT, TreatmentDT:
		return []FileType{CSV, PSV, TSV}
	case ReferenceDT:
		return []FileType{BIBTEX, JSONCSL, CSV, PSV, TSV}
	default:
		return []FileType{UnknownFileType}
	}
}

func (dt DataType) String() string {
	return DataTypeToString[dt]
}

// StringToDataType maps strings to DataTypes.
var StringToDataType = map[string]DataType{
	"Author":               AuthorDT,
	"NameRelation":         NameRelationDT,
	"Taxon":                TaxonDT,
	"Synonym":              SynonymDT,
	"NameUsage":            NameUsageDT,
	"Name":                 NameDT,
	"TaxonProperty":        TaxonPropertyDT,
	"TaxonConceptRelation": TaxonConceptRelationDT,
	"SpeciesInteraction":   SpeciesInteractionDT,
	"SpeciesEstimate":      SpeciesEstimateDT,
	"Reference":            ReferenceDT,
	"TypeMaterial":         TypeMaterialDT,
	"Distribution":         DistributionDT,
	"Media":                MediaDT,
	"VernacularName":       VernacularNameDT,
	"Treatment":            TreatmentDT,
}

// DataTypeToString maps DataType to string.
var DataTypeToString = func() map[DataType]string {
	res := make(map[DataType]string)
	for k, v := range StringToDataType {
		res[v] = k
	}
	res[UnkownDT] = "Unknown"
	return res
}()

// LowCaseToDataType provides map of low-case strings to DataType.
var LowCaseToDataType = func() map[string]DataType {
	res := make(map[string]DataType)
	for k, v := range StringToDataType {
		res[strings.ToLower(k)] = v
	}
	return res
}()

func NewDataType(s string) DataType {
	s = strings.ToLower(s)
	if dt, ok := LowCaseToDataType[s]; ok {
		return dt
	}
	return UnkownDT
}
