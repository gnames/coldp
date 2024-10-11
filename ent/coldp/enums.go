package coldp

type NomCode int

const (
	Bacterial NomCode = iota + 1
	Botanical
	Cultivars
	PhytoSociological
	Virus
	Zoological
)

type Gender int

const (
	Masculine Gender = iota + 1
	Feminine
	Neutral
)
