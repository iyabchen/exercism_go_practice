package protein

import ()

const testVersion = 1

// Codon                 | Protein
// :---                  | :---
// AUG                   | Methionine
// UUU, UUC              | Phenylalanine
// UUA, UUG              | Leucine
// UCU, UCC, UCA, UCG    | Serine
// UAU, UAC              | Tyrosine
// UGU, UGC              | Cysteine
// UGG                   | Tryptophan
// UAA, UAG, UGA         | STOP

var codonProteinMap = make(map[string]string)

// construct codonProteinMap
func init() {
	codonProteinMap["AUG"] = "Methionine"
	codonProteinMap["UUU"] = "Phenylalanine"
	codonProteinMap["UUC"] = "Phenylalanine"
	codonProteinMap["UUA"] = "Leucine"
	codonProteinMap["UUG"] = "Leucine"
	codonProteinMap["UCU"] = "Serine"
	codonProteinMap["UCC"] = "Serine"
	codonProteinMap["UCA"] = "Serine"
	codonProteinMap["UCG"] = "Serine"
	codonProteinMap["UAU"] = "Tyrosine"
	codonProteinMap["UAC"] = "Tyrosine"
	codonProteinMap["UGU"] = "Cysteine"
	codonProteinMap["UGC"] = "Cysteine"
	codonProteinMap["UGG"] = "Tryptophan"
	codonProteinMap["UAA"] = "STOP"
	codonProteinMap["UAG"] = "STOP"
	codonProteinMap["UGA"] = "STOP"
}

// alternatively, the map can be defined like
// var codons = map[string](string){
// 	"AUG": "Methionine",
// 	"UUU": "Phenylalanine",
// 	"UUC": "Phenylalanine",
// 	"UUA": "Leucine",
// 	"UUG": "Leucine",
// 	"UCU": "Serine",
// 	"UCC": "Serine",
// 	"UCA": "Serine",
// 	"UCG": "Serine",
// 	"UAU": "Tyrosine",
// 	"UAC": "Tyrosine",
// 	"UGU": "Cysteine",
// 	"UGC": "Cysteine",
// 	"UGG": "Tryptophan",
// 	"UAA": "STOP",
// 	"UAG": "STOP",
// 	"UGA": "STOP",
// }

// For a given codon, return the protein name
func FromCodon(src string) string {
	return codonProteinMap[src]

}

func FromRNA(src string) []string {
	// the length of src must be multipliers of 3
	if len(src)%3 != 0 {
		return nil
	}
	// must only contains AUGC
	for _, v := range src {
		switch v {
		case 'A':
		case 'C':
		case 'G':
		case 'U':
		default:
			return nil
		}
	}

	ret := []string{}
	for i := 0; i < len(src); i = i + 3 {
		s := string(src[i : i+3])
		protein := FromCodon(s)
		if protein != "STOP" {
			ret = append(ret, protein)
		} else {
			break
		}
	}
	return ret

}
