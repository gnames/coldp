package coldp

import "strings"

type Rank int

func (r Rank) String() string {
	var res string
	var ok bool

	if res, ok = rankToAbbr[r]; ok {
		return res + "."
	}
	if res, ok = rankToString[r]; ok {
		return res
	}
	// something has to be added if you see this return
	return "missing, fix me"
}

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

const (
	UnknownRank Rank = iota
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
	Unranked
	Variety
)

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

var rankToAbbr = func() map[Rank]string {
	res := make(map[Rank]string)
	for k, v := range abbrToRank {
		res[v] = k
	}
	return res
}()

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

var stringToRank = func() map[string]Rank {
	res := make(map[string]Rank)
	for k, v := range rankToString {
		res[v] = k
	}
	return res
}()
