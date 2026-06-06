import java.util.Map;
import java.util.HashMap;

class Scrabble {
    
    private String word;
    private Map<Character,Integer> lookup;
    
    Scrabble(String word) {
        this.word = word.toLowerCase();
        this.lookup = new HashMap<Character,Integer>();
        prepareLookup();
    }

    void prepareLookup() {
        this.lookup.put('a',1);
        this.lookup.put('e',1);
        this.lookup.put('i',1);
        this.lookup.put('o',1);
        this.lookup.put('u',1);
        this.lookup.put('l',1);
        this.lookup.put('n',1);
        this.lookup.put('r',1);
        this.lookup.put('s',1);
        this.lookup.put('t',1);
        this.lookup.put('d',2);
        this.lookup.put('g',2);
        this.lookup.put('b',3);
        this.lookup.put('c',3);
        this.lookup.put('m',3);
        this.lookup.put('p',3);
        this.lookup.put('f',4);
        this.lookup.put('h',4);
        this.lookup.put('v',4);
        this.lookup.put('w',4);
        this.lookup.put('y',4);
        this.lookup.put('k',5);
        this.lookup.put('j',8);
        this.lookup.put('x',8);
        this.lookup.put('z',10);
        this.lookup.put('q',10);
    }

    int getScore() {
        int score = 0;
        for (char c : this.word.toCharArray()){
            score += this.lookup.get(c);
        }
        return score;
    }

}
