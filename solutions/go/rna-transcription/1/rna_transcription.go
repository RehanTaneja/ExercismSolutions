package strand

func ToRNA(dna string) string {
	s := ""
    for _,x := range dna{
        switch x{
            case 'G':
            	s += string('C')
            case 'C':
            	s += string('G')
            case 'T':
            	s += string('A')
            case 'A':
            	s += string('U')
            default:
            	return ""
        }
    }
    return s
}
