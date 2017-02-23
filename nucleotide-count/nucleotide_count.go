package dna

import (
	"errors"
)

const testVersion = 2

// Given a DNA string, compute how many times each nucleotide occurs in the string.
// A T C G, Each symbol represents a nucleotide,

type DNA string

type Histogram map[byte]int

// Return how many times a symbol shows
func (dna DNA) Count(b byte) (int, error) {
	if b != 'A' && b != 'C' && b != 'T' && b != 'G' {
		return -1, errors.New("invalid nucleotide")
	}
	h := make(map[byte]int)
	h['A'] = 0
	h['C'] = 0
	h['T'] = 0
	h['G'] = 0
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'A':
			h['A']++
		case 'G':
			h['G']++
		case 'T':
			h['T']++
		case 'C':
			h['C']++
		default:
			return -1, errors.New("invalid DNA")
		}
	}
	return h[b], nil
}

// Return all symbols and the counts
func (dna DNA) Counts() (Histogram, error) {
	h := make(map[byte]int)
	h['A'] = 0
	h['C'] = 0
	h['T'] = 0
	h['G'] = 0
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'A':
			h['A']++
		case 'G':
			h['G']++
		case 'T':
			h['T']++
		case 'C':
			h['C']++
		default:
			return nil, errors.New("invalid DNA")
		}
	}
	return h, nil

}
