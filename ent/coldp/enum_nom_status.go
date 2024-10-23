package coldp

import (
	"log/slog"
	"strings"
)

// NomStatus represents the nomenclatural status of a scientific name.
type NomStatus int

// Constants for different nomenclatural statuses.
const (
	UnknownNomStatus NomStatus = iota
	Established                // The name is validly published and available.
	NotEstablished             // The name is not validly published or unavailable.
	Acceptable                 // The name is potentially valid or acceptable for use.
	Unacceptable               // The name is illegitimate or unacceptable for use.
	Conserved                  // The name is conserved against a competing name.
	Rejected                   // The name is rejected in favor of a competing name.
	Doubtful                   // The application of the name is uncertain.
	Manuscript                 // The name is only published in a manuscript.
	Chresonym                  // The name is a chresonym.
)

// NewNomStatus creates a new NomStatus from a string representation.
// It handles various synonyms and normalizations to ensure consistent matching.
func NewNomStatus(s string) NomStatus {
	sOrig := s
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ".", "")
	switch s {
	case "":
		return UnknownNomStatus
	case "nomenvalidum", "available", "established":
		return Established
	case "nominval", "nomeninvalidum", "unavailable", "notestablished":
		return NotEstablished
	case "nomenlegitimum", "potentiallyvalid", "acceptable":
		return Acceptable
	case "nomilleg", "nomenillegitimum", "objectivelyinvalid", "unacceptable":
		return Unacceptable
	case "nomcons", "nomenconservandum", "conservedname", "conserved":
		return Conserved
	case "nomrej", "nomenrejiciendum", "rejected":
		return Rejected
	case "nomdub", "nomeddubium", "doubtful":
		return Doubtful
	case "manuscript name", "manuscript":
		return Manuscript
	case "chresonym":
		return Chresonym
	default:
		slog.Warn("Cannot find nom. status", "input", sOrig)
		return UnknownNomStatus
	}
}

// StringByCode returns the string representation of the nomenclatural status
// according to the provided nomenclatural code.
func (n NomStatus) StringByCode(c NomCode) string {
	switch n {
	case Established:
		switch c {
		case Botanical:
			return "nomen validum"
		case Zoological:
			return "available"
		default:
			return "established"
		}

	case NotEstablished:
		switch c {
		case Botanical:
			return "nomen invalidum"
		case Zoological:
			return "unavailable"
		default:
			return "not established"
		}

	case Acceptable:
		switch c {
		case Botanical:
			return "nomen legitimum"
		case Zoological:
			return "potentially valid"
		default:
			return "acceptable"
		}

	case Unacceptable:
		switch c {
		case Botanical:
			return "nomen illegitimum"
		case Zoological:
			return "objectively invalid"
		default:
			return "unacceptable"
		}

	case Conserved:
		switch c {
		case Botanical:
			return "nomen conservandum"
		case Zoological:
			return "conserved name"
		default:
			return "conserved"
		}

	case Rejected:
		switch c {
		case Botanical:
			return "nomen rejiciendum"
		case Zoological:
			return "rejected"
		default:
			return "rejected"
		}

	case Doubtful:
		switch c {
		case Botanical:
			return "nomen dubium"
		case Zoological:
			return "doubtful"
		default:
			return "doubtful"
		}

	case Manuscript:
		switch c {
		case Botanical:
			return "manuscript name"
		case Zoological:
			return "manuscript name"
		default:
			return "manuscript"
		}

	case Chresonym:
		return "chresonym"
	}
	return "missing status, fixme"
}
