package protein
import "errors"
var ErrInvalidBase, ErrStop = errors.New("invalid base"), errors.New("stop")
func FromRNA(rna string) (s []string, e error) {
	for i := 0; i <= len(rna)- 3; i+=3{
        codon := string(rna[i: i+3])
        protein, err := FromCodon(codon)
        if err == ErrInvalidBase{
            return nil, ErrInvalidBase
        }else if err == ErrStop{
        	return s, e
        }
    	s = append(s, protein)
    }
	return s, e
}
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}
	return "", ErrInvalidBase
}