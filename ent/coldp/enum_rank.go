package coldp

import "strings"

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
	case "superdom":
		return Superdomain
	case "dom":
		return Domain
	case "superreg":
		return Superkingdom
	case "reg":
		return Kingdom
	case "subreg":
		return Subkingdom
	case "infrareg":
		return Infrakingdom
	case "superphyl":
		return Superphylum
	case "phyl":
		return Phylum
	case "subphyl":
		return Subphylum
	case "infraphyl":
		return Infraphylum
	case "parvphyl":
		return Parvphylum
	case "microphyl":
		return Microphylum
	case "gigacl":
		return Gigaclass
	case "megacl":
		return Megaclass
	case "supercl":
		return Superclass
	case "cl":
		return Class
	case "subcl":
		return Subclass
	case "subtercl":
		return Subterclass
	case "parvcl":
		return Parvclass
	case "superdiv":
		return Superdivision
	case "div":
		return Division
	case "subdiv":
		return Subdivision
	case "infradiv":
		return Infradivision
	case "superleg":
		return Superlegion
	case "legion":
		return Legion
	case "subleg":
		return Sublegion
	case "infraleg":
		return Infralegion
	case "gigaord":
		return Gigaorder
	case "grandord":
		return Grandorder
	case "mirord":
		return Mirorder
	case "superord":
		return Superorder
	case "ord":
		return Order
	case "nanord":
		return Nanorder
	case "hypord":
		return Hypoorder
	case "minord":
		return Minorder
	case "subord":
		return Suborder
	case "infraord":
		return Infraorder
	case "parvord":
		return Parvorder
	case "supersect":
		return Supersection
	case "sect":
		return Section
	case "subsect":
		return Subsection
	case "megafam":
		return Megafamily
	case "grandfam":
		return Grandfamily
	case "superfam":
		return Superfamily
	case "fam":
		return Family
	case "subfam":
		return Subfamily
	case "infrafam":
		return Infrafamily
	case "supertrib":
		return Supertribe
	case "trib":
		return Tribe
	case "subtrib":
		return Subtribe
	case "infratrib":
		return Infratribe
	case "supergen":
		return Supergenus
	case "gen":
		return Genus
	case "subgen":
		return Subgenus
	case "infragen":
		return Infragenus
	case "superser":
		return Superseries
	case "ser":
		return Series
	case "subser":
		return Subseries
	case "sp":
		return Species
	case "infrasp":
		return InfraspecificName
	case "gx":
		return Grex
	case "subsp":
		return Subspecies
	case "convar":
		return Convariety
	case "infrasubsp":
		return InfrasubspecificName
	case "ab":
		return Aberration
	case "supervar":
		return Supervariety
	case "var":
		return Variety
	case "subvar":
		return Subvariety
	case "superf":
		return Superform
	case "f":
		return Form
	case "subf":
		return Subform
	case "pv":
		return Pathovar
	case "fsp":
		return FormaSpecialis
	case "cv":
		return Cultivar
	case "mut":
		return Mutatio
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
