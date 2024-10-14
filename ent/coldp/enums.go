package coldp

import "strings"

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

func NewNomCode(s string) NomCode {
	s = strings.ToLower(s)
	switch s {
	case "1", "bacterial":
		return Bacterial
	case "2", "botanical":
		return Botanical
	case "3", "cultivars":
		return Cultivars
	case "4", "phytosociological":
		return PhytoSociological
	case "5", "virus":
		return Virus
	case "6", "zoological":
		return Zoological
	default:
		return UnknownCode
	}
}

type ArchiveType int

const (
	ArchUnknown ArchiveType = iota
	ArchName
	ArchNameUsage
)

type Gender int

const (
	Masculine Gender = iota + 1
	Feminine
	Neutral
)

type FileType int

const (
	UnknownFileType FileType = iota
	JSON
	YAML
	CSV
	TSV
	// pipe-separated file
	PSV
)
