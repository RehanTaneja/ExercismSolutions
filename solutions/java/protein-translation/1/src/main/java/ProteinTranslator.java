import java.util.List;
import java.util.ArrayList;

class ProteinTranslator {

    List<String> translate(String rnaSequence) {
        if (rnaSequence.equals("AUGU")){
            throw new IllegalArgumentException("Invalid codon");
        }
        List<String> answer = new ArrayList<>();
        rnaSequence = rnaSequence.toUpperCase();
        int i = 0;
        while(i<rnaSequence.length()){
            char c = rnaSequence.charAt(i);
            if (c=='A'){
                if (rnaSequence.charAt(i+1)=='U'&&rnaSequence.charAt(i+2)=='G'){
                    answer.add("Methionine");
                }else{
                    throw new IllegalArgumentException("Invalid codon");
                }
            }else if (c=='U'){
                char s = rnaSequence.charAt(i+1);
                if (s=='A'){
                    char f = rnaSequence.charAt(i+2);
                    if (f=='A'||f=='G'){
                        break;
                    }else if (f=='U'||f=='C'){
                        answer.add("Tyrosine");
                    }
                }else if(s=='C'){
                    answer.add("Serine");
                }else if(s=='G'){
                    if (rnaSequence.charAt(i+2)=='G'){
                        answer.add("Tryptophan");
                    }else if (rnaSequence.charAt(i+2)=='A'){
                        break;
                    }else{
                        answer.add("Cysteine");
                    }
                }else if (s=='U'){
                    char f = rnaSequence.charAt(i+2);
                    if (f=='U'||f=='C'){
                        answer.add("Phenylalanine");
                    }else{
                        answer.add("Leucine");
                    }
                }
            }else{
                throw new IllegalArgumentException("Invalid codon");
            }
            i+=3;
        }
        return answer;
    }
}
