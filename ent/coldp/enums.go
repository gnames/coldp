package coldp

import "strings"

// NomCode provides types of nomenclatural Codes.
type NomCode int

// Constants for different nomenclatural codes.
const (
	UnknownCode       NomCode = iota
	Bacterial                 // Bacteriological Code
	Botanical                 // Botanical Code
	Cultivars                 // Cultivated Plant Code
	PhytoSociological         // Phytosociological Code
	Virus                     // Virus Code
	Zoological                // Zoological Code
)

// NewNomCode converts a string (number or word) to NomCode.
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

// NamePart represents the part of a scientific name.
type NamePart int

// Constants for different name parts.
const (
	UnknownNP       NamePart = iota
	GenericNP                // The genus part of the name.
	InfragenericNP           // The infrageneric epithet (e.g., subgenus, section) of the name.
	SpecificNP               // The specific epithet of the name.
	InfraspecificNP          // The infraspecific epithet (e.g., subspecies, variety) of the name.
)

// NewNamePart creates a new NamePart from a string representation.
func NewNamePart(s string) NamePart {
	s = strings.ToLower(s)
	switch s {
	case "1", "generic":
		return GenericNP
	case "2", "infrageneric":
		return InfraspecificNP // Note: This seems like a typo, should be InfragenericNP
	case "3", "specific":
		return SpecificNP
	case "4", "infraspecific":
		return InfraspecificNP
	default:
		return UnknownNP
	}
}

// ArchiveType provides type of CoLDP. Can be 'flat' or NameUsage, and
// 'original' or Name.
type ArchiveType int

// Constants for different archive types.
const (
	UnknownAT   ArchiveType = iota
	NameAT                  // Represents an archive with Name data.
	NameUsageAT             // Represents an archive with NameUsage data.
)

// Gender represents the grammatical gender of a scientific name.
type Gender int

// Constants for different grammatical genders.
const (
	UnknownG Gender = iota
	Masculine
	Feminine
	Neutral
)

// NewGender creates a new Gender from a string representation.
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

// String returns the string representation of the gender.
func (g Gender) String() string {
	switch g {
	case Masculine:
		return "m"
	case Feminine:
		return "f"
	case Neutral:
		return "neut"
	default:
		return ""
	}
}

// FileType represents the type of file.
type FileType int

// Constants for different file types.
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
// It is used to convert a file name to the type of information
// provided in that file.
type DataType int

// Constants for different data types.
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

// FileFormats returns a list of supported file formats for a given DataType.
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

// String returns the string representation of the DataType.
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

// NewDataType creates a new DataType from a string representation.
func NewDataType(s string) DataType {
	s = strings.ToLower(s)
	if dt, ok := LowCaseToDataType[s]; ok {
		return dt
	}
	return UnkownDT
}
