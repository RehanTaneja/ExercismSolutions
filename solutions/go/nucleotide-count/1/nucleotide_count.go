package dna

import(
    "errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

func (d DNA) Counts() (Histogram, error) {
	var h Histogram = make(map[rune]int)
    h['A']=0
    h['C']=0
    h['G']=0
    h['T']=0
    hasInvalid := false
    for _,x := range d{
        switch x{
            case 'A','G','C','T':
            	h[x]++
            default:
            	hasInvalid = true;
        }
    }
    if hasInvalid{
        return h,errors.New("INVALID")
    }
	return h, nil
}
