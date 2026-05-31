import java.util.List;
import java.util.ArrayList;
import java.util.Collections;

class HighScores {

    private List<Integer> s;

    public HighScores(List<Integer> highScores) {
        s = highScores;
    }

    List<Integer> scores() {
        return this.s;
    }

    Integer latest() {
        return s.get(s.size()-1);
    }

    Integer personalBest() {
        int best = this.s.get(0);
        for(int i=1;i<this.s.size();i++){
            if (this.s.get(i)>best){
                best = this.s.get(i);
            }
        }
        return best;
    }

    List<Integer> personalTopThree() {
        if (this.s.size()<=3){
            Collections.sort(this.s,Collections.reverseOrder());
            return this.s;
        }
        int best1 = -100;
        int best2 = -100;
        int best3 = -100;
        for(int i=0;i<this.s.size();i++){
            if (this.s.get(i)>best1){
                best3 = best2;
                best2 = best1;
                best1 = this.s.get(i);
            }else if(this.s.get(i)>best2){
                best3 = best2;
                best2 = this.s.get(i);
            }else if(this.s.get(i)>best3){
                best3 = this.s.get(i);
            }
        }
        List<Integer> a = new ArrayList<>();
        a.add(best1);
        a.add(best2);
        a.add(best3);
        return a;
    }

}
