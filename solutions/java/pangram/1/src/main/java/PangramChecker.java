import java.util.Map;
import java.util.HashMap;

public class PangramChecker {

    public boolean isPangram(String input) {
        if (input==null || input.length()<26){
            return false;
        }
        input = input.toLowerCase();
        Map<Character,Integer> a = new HashMap<>();
        for (int i = 0;i<input.length();i++){
            if (Character.isLetter(input.charAt(i))){
                a.put(input.charAt(i),a.getOrDefault(input.charAt(i),0)+1);
            }
        }
        return a.keySet().size() >= 26;
    }

}
