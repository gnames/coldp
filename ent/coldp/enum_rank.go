package coldp

import "strings"

// Rank represents the taxonomic rank of a scientific name.
type Rank int

// String returns the string representation of the rank.
// It prioritizes abbreviations if available, then falls back to full names.
func (r Rank) String() string {
	var res string
	var ok bool

	switch r {
	case UnknownRank, Unranked:
		return ""
	}
	if res, ok = rankToAbbr[r]; ok {
		return res + "."
	}
	if res, ok = rankToString[r]; ok {
		return res
	}
	// something has to be added if you see this return
	return "missing, fix me"
}

// NewRank creates a new Rank from a string representation.
// It handles trimming and normalization to ensure consistent matching.
func NewRank(s string) Rank {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, ".")
	s = strings.ToLower(s)
	if s == "" {
		return Unranked
	}

	var res Rank
	var ok bool
	if res, ok = abbrToRank[s]; ok {
		return res
	}
	if res, ok = stringToRank[s]; ok {
		return res
	}
	return UnknownRank
}

// Constants for different taxonomic ranks.
const (
	UnknownRank Rank = iota
	Superdomain
	Domain
	Subdomain
	Infradomain
	Empire
	Realm
	Subrealm
	Superkingdom
	Kingdom
	Subkingdom
	Infrakingdom
	Superphylum
	Phylum
	Subphylum
	Infraphylum
	Parvphylum
	Microphylum
	Nanophylum
	Claudius
	Gigaclass
	Megaclass
	Superclass
	Class
	subclass
	Infraclass
	Subterclass
	Parvclass
	Superdivision
	Division
	Subdivision
	Infradivision
	Superlegion
	Legion
	Sublegion
	Infralegion
	Megacohort
	Supercohort
	Cohort
	Subcohort
	Infracohort
	Gigaorder
	Magnorder
	Grandorder
	Mirorder
	Superorder
	Order
	Nanorder
	Hypoorder
	Minorder
	Suborder
	Infraorder
	Parvorder
	SupersectionZoology
	SectionZoology
	SubsectionZoology
	Falanx
	Gigafamily
	Megafamily
	Grandfamily
	Superfamily
	Epifamily
	Family
	Subfamily
	Infrafamily
	Supertribe
	Tribe
	Subtribe
	Infratribe
	SupragenericName
	Supergenus
	Genus
	Subgenus
	Infragenus
	SupersectionBotany
	SectionBotany
	SubsectionBotany
	Superseries
	Series
	Subseries
	InfragenericName
	SpeciesAggregate
	Species
	InfraspecificName
	Grex
	Klepton
	Subspecies
	CultivarGroup
	Convariety
	InfrasubspecificName
	Proles
	Natio
	Aberration
	Morph
	Supervariety
	Variety
	Subvariety
	Superform
	Form
	Subform
	Pathovar
	Biovar
	Chemovar
	Morphovar
	Phagovar
	Serovar
	Chemoform
	FormaSpecialis
	Lusus
	Cultivar
	Mutatio
	Strain
	Other
	Unranked
)

// abbrToRank maps rank abbreviations to their corresponding Rank values.
var abbrToRank = map[string]Rank{
	"superdom":   Superdomain,
	"dom":        Domain,
	"superreg":   Superkingdom,
	"reg":        Kingdom,
	"subreg":     Subkingdom,
	"infrareg":   Infrakingdom,
	"superphyl":  Superphylum,
	"phyl":       Phylum,
	"subphyl":    Subphylum,
	"infraphyl":  Infraphylum,
	"parvphyl":   Parvphylum,
	"microphyl":  Microphylum,
	"gigacl":     Gigaclass,
	"megacl":     Megaclass,
	"supercl":    Superclass,
	"cl":         Class,
	"subcl":      Subclass,
	"subtercl":   Subterclass,
	"parvcl":     Parvclass,
	"superdiv":   Superdivision,
	"div":        Division,
	"subdiv":     Subdivision,
	"infradiv":   Infradivision,
	"superleg":   Superlegion,
	"legion":     Legion,
	"subleg":     Sublegion,
	"infraleg":   Infralegion,
	"gigaord":    Gigaorder,
	"grandord":   Grandorder,
	"mirord":     Mirorder,
	"superord":   Superorder,
	"ord":        Order,
	"nanord":     Nanorder,
	"hypord":     Hypoorder,
	"minord":     Minorder,
	"subord":     Suborder,
	"infraord":   Infraorder,
	"parvord":    Parvorder,
	"supersect":  Supersection,
	"sect":       Section,
	"subsect":    Subsection,
	"megafam":    Megafamily,
	"grandfam":   Grandfamily,
	"superfam":   Superfamily,
	"fam":        Family,
	"subfam":     Subfamily,
	"infrafam":   Infrafamily,
	"supertrib":  Supertribe,
	"trib":       Tribe,
	"subtrib":    Subtribe,
	"infratrib":  Infratribe,
	"supergen":   Supergenus,
	"gen":        Genus,
	"subgen":     Subgenus,
	"infragen":   Infragenus,
	"superser":   Superseries,
	"ser":        Series,
	"subser":     Subseries,
	"sp":         Species,
	"infrasp":    InfraspecificName,
	"gx":         Grex,
	"subsp":      Subspecies,
	"convar":     Convariety,
	"infrasubsp": InfrasubspecificName,
	"ab":         Aberration,
	"supervar":   Supervariety,
	"var":        Variety,
	"subvar":     Subvariety,
	"superf":     Superform,
	"f":          Form,
	"subf":       Subform,
	"pv":         Pathovar,
	"f.sp":       FormaSpecialis,
	"cv":         Cultivar,
	"mut":        Mutatio,
}

// rankToAbbr maps Rank values to their corresponding abbreviations.
var rankToAbbr = func() map[Rank]string {
	res := make(map[Rank]string)
	for k, v := range abbrToRank {
		res[v] = k
	}
	return res
}()

// rankToString maps Rank values to their full string names.
var rankToString = map[Rank]string{
	UnknownRank:          "unknown",
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

// abbrToRank maps rank strings to their corresponding Rank values.
var stringToRank = func() map[string]Rank {
	res := make(map[string]Rank)
	for k, v := range rankToString {
		res[v] = k
	}
	return res
}()
