package protein

import (
	"errors"
)

var ErrStop error = errors.New("stop")
var ErrInvalidBase error = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	sequence := make([]string, 0)

	for i := 0; i < len(rna); i = i + 3 {
		codon := rna[i : i+3]
		ptn, err := FromCodon(codon)

		if err == ErrStop {
			break
		}

		if err == ErrInvalidBase {
			return nil, ErrInvalidBase
		}

		sequence = append(sequence, ptn)
	}

	return sequence, nil
}

func FromCodon(codon string) (string, error) {
	table := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGC": "Cysteine",
		"UGU": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	ptn, ok := table[codon]
	if !ok {
		return "", ErrInvalidBase
	}

	if ptn == "STOP" {
		return "", ErrStop
	}

	return ptn, nil
}
