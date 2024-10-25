package coldp

import "strings"

type Gazetteer int

const (
	UnknownGz Gazetteer = iota
	TDWG
	ISO
	FAO
	LONGHURST
	TEOW
	IHO
	MRGID
	TextGz
)

var gazetteerToString = map[Gazetteer]string{
	TDWG:      "TDWG",
	ISO:       "ISO",
	FAO:       "FAO",
	LONGHURST: "LONGHURST",
	TEOW:      "TEOW",
	IHO:       "IHO",
	MRGID:     "MRGID",
	TextGz:    "TEXT",
}

var stringToGazetteer = func() map[string]Gazetteer {
	res := make(map[string]Gazetteer)
	for k, v := range gazetteerToString {
		res[v] = k
	}
	return res
}()

func (g Gazetteer) String() string {
	if res, ok := gazetteerToString[g]; ok {
		return res
	}
	return ""
}

func NewGazetteer(s string) Gazetteer {
	s = strings.ToUpper(s)
	if res, ok := stringToGazetteer[s]; ok {
		return res
	}
	return UnknownGz
}
