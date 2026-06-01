import java.util.List;
import java.util.Set;
import java.util.HashSet;

class GottaSnatchEmAll {

    static Set<String> newCollection(List<String> cards) {
        Set<String> s = new HashSet<>();
        for(int i=0;i<cards.size();i++){
            s.add(cards.get(i));
        }
        return s;
    }

    static boolean addCard(String card, Set<String> collection) {
        boolean answer = collection.contains(card);
        collection.add(card);
        return !answer;
    }

    static boolean canTrade(Set<String> myCollection, Set<String> theirCollection) {
        boolean answer1 = false;
        for(String s : myCollection){
            if (!theirCollection.contains(s)){
                answer1 = true;
                break;
            }
        }
        boolean answer2 = false;
        for (String s : theirCollection){
            if (!myCollection.contains(s)){
                answer2 = true;
            }
        }
        return answer1 && answer2;
    }

    static Set<String> commonCards(List<Set<String>> collections) {
        Set<String> answer = new HashSet<>();
        for(String s: collections.get(0)){
            boolean a = true;
            for(Set<String> i : collections){
                if (!i.contains(s)){
                    a = false;
                }
            }
            if (a){
                answer.add(s);
            }
        }
        return answer;
    }

    static Set<String> allCards(List<Set<String>> collections) {
        Set<String> answer = new HashSet<>();
        for(Set<String> s: collections){
            for(String i: s){
                answer.add(i);
            }
        }
        return answer;
    }
}
