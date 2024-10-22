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

type NameStatus int

const (
	UnknownNS NameStatus = iota
	EstablishedNS
	AcceptableNS
	UnacceptableNS
	ConservedNS
	RejectedNS
	DoubtfulNS
	ManuscriptNS
	ChresonymNS
)

func NewNameStatus(s string) NameStatus {
	s = strings.ToLower(s)
	switch s {
	case "1", "established":
		return EstablishedNS
	case "2", "acceptable":
		return AcceptableNS
	case "3", "unacceptable":
		return UnacceptableNS
	case "4", "conserved":
		return ConservedNS
	case "5", "rejected":
		return RejectedNS
	case "6", "doubtful":
		return DoubtfulNS
	case "7", "manuscript":
		return ManuscriptNS
	case "8", "chresonym":
		return ChresonymNS
	default:
		return UnknownNS
	}
}

type NamePart int

const (
	UnknownNP NamePart = iota
	GenericNP
	InfragenericNP
	SpecificNP
	InfraspecificNP
)

func NewNamePart(s string) NamePart {
	s = strings.ToLower(s)
	switch s {
	case "1", "generic":
		return GenericNP
	case "2", "infrageneric":
		return InfraspecificNP
	case "3", "specific":
		return SpecificNP
	case "4", "infraspecific":
		return InfraspecificNP
	default:
		return UnknownNP
	}
}

type Rank int

const (
	Unranked Rank = iota
	Aberration
	Biovar
	Chemoform
	Chemovar
	Class
	Cohort
	Convariety
	Cultivar
	CultivarGroup
	Division
	Domain
	Epifamily
	Falanx
	Family
	Form
	FormaSpecialis
	Genus
	Gigaclass
	Gigaorder
	Grandfamily
	Grandorder
	Grex
	Hypoorder
	Infraclass
	Infracohort
	Infradivision
	Infrafamily
	InfragenericName
	Infragenus
	Infrakingdom
	Infralegion
	Infraorder
	Infraphylum
	InfraspecificName
	InfrasubspecificName
	Infratribe
	Kingdom
	Klepton
	Legion
	Lusus
	Magnorder
	Megaclass
	Megacohort
	Megafamily
	Microphylum
	Minorder
	Mirorder
	Morph
	Morphovar
	Mutatio
	Nanophylum
	Nanorder
	Natio
	Order
	Other
	Parvclass
	Parvorder
	Parvphylum
	Pathovar
	Phagovar
	Phylum
	Proles
	Realm
	Section
	Series
	Serovar
	Species
	SpeciesAggregate
	Strain
	Subclass
	Subcohort
	Subdivision
	Subfamily
	Subform
	Subgenus
	Subkingdom
	Sublegion
	Suborder
	Subphylum
	Subrealm
	Subsection
	Subseries
	Subspecies
	Subterclass
	Subtribe
	Subvariety
	Superclass
	Supercohort
	Superdivision
	Superdomain
	Superfamily
	Superform
	Supergenus
	Superkingdom
	Superlegion
	Superorder
	Superphylum
	Supersection
	Superseries
	Supertribe
	Supervariety
	SupragenericName
	Tribe
	Variety
)

func (r Rank) String() string {
	if res, ok := rankToString[r]; ok {
		return res
	}
	return "unknown?"
}

func NewRank(s string) Rank {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "_-", " ")
	s = strings.TrimSpace(s)
	switch s {
	case "dom":
		return Domain
	case "div":
		return Division
	case "f":
		return Form
	case "var":
		return Variety
	}
	if res, ok := stringToRank[s]; ok {
		return res
	}
	return Unranked
}

var rankToString = map[Rank]string{
	Unranked:             "unranked",
	Aberration:           "aberration",
	Biovar:               "biovar",
	Chemoform:            "chemoform",
	Chemovar:             "chemovar",
	Class:                "class",
	Cohort:               "cohort",
	Convariety:           "convariety",
	Cultivar:             "cultivar",
	CultivarGroup:        "cultivar group",
	Division:             "division",
	Domain:               "domain",
	Epifamily:            "epifamily",
	Falanx:               "falanx",
	Family:               "family",
	Form:                 "form",
	FormaSpecialis:       "forma specialis",
	Genus:                "genus",
	Gigaclass:            "gigaclass",
	Gigaorder:            "gigaorder",
	Grandfamily:          "grandfamily",
	Grandorder:           "grandorder",
	Grex:                 "grex",
	Hypoorder:            "hypoorder",
	Infraclass:           "infraclass",
	Infracohort:          "infracohort",
	Infradivision:        "infradivision",
	Infrafamily:          "infrafamily",
	InfragenericName:     "infrageneric name",
	Infragenus:           "infragenus",
	Infrakingdom:         "infrakingdom",
	Infralegion:          "infralegion",
	Infraorder:           "infraorder",
	Infraphylum:          "infraphylum",
	InfraspecificName:    "infraspecific name",
	InfrasubspecificName: "infrasubspecific name",
	Infratribe:           "infratribe",
	Kingdom:              "kingdom",
	Klepton:              "klepton",
	Legion:               "legion",
	Lusus:                "lusus",
	Magnorder:            "magnorder",
	Megaclass:            "megaclass",
	Megacohort:           "megacohort",
	Megafamily:           "megafamily",
	Microphylum:          "microphylum",
	Minorder:             "minorder",
	Mirorder:             "mirorder",
	Morph:                "morph",
	Morphovar:            "morphovar",
	Mutatio:              "mutatio",
	Nanophylum:           "nanophylum",
	Nanorder:             "nanorder",
	Natio:                "natio",
	Order:                "order",
	Other:                "other",
	Parvclass:            "parvclass",
	Parvorder:            "parvorder",
	Parvphylum:           "parvphylum",
	Pathovar:             "pathovar",
	Phagovar:             "phagovar",
	Phylum:               "phylum",
	Proles:               "proles",
	Realm:                "realm",
	Section:              "section",
	Series:               "series",
	Serovar:              "serovar",
	Species:              "species",
	SpeciesAggregate:     "species aggregate",
	Strain:               "strain",
	Subclass:             "subclass",
	Subcohort:            "subcohort",
	Subdivision:          "subdivision",
	Subfamily:            "subfamily",
	Subform:              "subform",
	Subgenus:             "subgenus",
	Subkingdom:           "subkingdom",
	Sublegion:            "sublegion",
	Suborder:             "suborder",
	Subphylum:            "subphylum",
	Subrealm:             "subrealm",
	Subsection:           "subsection",
	Subseries:            "subseries",
	Subspecies:           "subspecies",
	Subterclass:          "subterclass",
	Subtribe:             "subtribe",
	Subvariety:           "subvariety",
	Superclass:           "superclass",
	Supercohort:          "supercohort",
	Superdivision:        "superdivision",
	Superdomain:          "superdomain",
	Superfamily:          "superfamily",
	Superform:            "superform",
	Supergenus:           "supergenus",
	Superkingdom:         "superkingdom",
	Superlegion:          "superlegion",
	Superorder:           "superorder",
	Superphylum:          "superphylum",
	Supersection:         "supersection",
	Superseries:          "superseries",
	Supertribe:           "supertribe",
	Supervariety:         "supervariety",
	SupragenericName:     "suprageneric name",
	Tribe:                "tribe",
	Variety:              "variety",
}

var stringToRank = func() map[string]Rank {
	res := make(map[string]Rank)
	for k, v := range rankToString {
		res[v] = k
	}
	return res
}()

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
