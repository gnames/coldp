package coldp

import "strings"

type GazetteerEnt int

const (
	UnknownGz GazetteerEnt = iota
	TDWG
	ISO
	FAO
	LONGHURST
	TEOW
	IHO
	MRGID
	TextGz
)

var gazetteerToString = map[GazetteerEnt]string{
	TDWG:      "TDWG",
	ISO:       "ISO",
	FAO:       "FAO",
	LONGHURST: "LONGHURST",
	TEOW:      "TEOW",
	IHO:       "IHO",
	MRGID:     "MRGID",
	TextGz:    "TEXT",
}

var stringToGazetteer = func() map[string]GazetteerEnt {
	res := make(map[string]GazetteerEnt)
	for k, v := range gazetteerToString {
		res[v] = k
	}
	return res
}()

func (g GazetteerEnt) String() string {
	if res, ok := gazetteerToString[g]; ok {
		return res
	}
	return ""
}

func NewGazetteerEnt(s string) GazetteerEnt {
	s = strings.ToUpper(s)
	if res, ok := stringToGazetteer[s]; ok {
		return res
	}
	return UnknownGz
}
