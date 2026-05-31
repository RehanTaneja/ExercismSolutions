import java.util.List;
import java.util.ArrayList;
import java.util.Random;
import java.lang.Math;

class DnDCharacter {
    int strength;
    int dexterity;
    int constitution;
    int intelligence;
    int wisdom;
    int charisma;
    int modifier;
    int hitpoints;

    int ability(List<Integer> scores) {
        int s = 0;
        int smallest = scores.get(0);
        for(int i =0;i<scores.size();i++){
            if (scores.get(i)<smallest){
                smallest = scores.get(i);
            }
            s+=scores.get(i);
        }
        return s - smallest;
    }

    List<Integer> rollDice() {
        Random random = new Random();
        List<Integer> a = new ArrayList<>();
        a.add(random.nextInt(6)+1);
        a.add(random.nextInt(6)+1);
        a.add(random.nextInt(6)+1);
        a.add(random.nextInt(6)+1);
        return a;
    }

    int modifier(int input) {
        return (int)Math.floor(((float)input-10)/2);
    }

    int getStrength() {
        if (this.strength==0){
            this.strength = ability(rollDice());
        }
        return this.strength;
    }

    int getDexterity() {
        if (this.dexterity==0){
            this.dexterity = ability(rollDice());
        }
        return this.dexterity;
    }

    int getConstitution() {
        if (this.constitution==0){
            this.constitution = ability(rollDice());
        }
        return this.constitution;
    }

    int getIntelligence() {
        if (this.intelligence==0){
            this.intelligence = ability(rollDice());
        }
        return this.intelligence;
    }

    int getWisdom() {
        if (this.wisdom==0){
            this.wisdom = ability(rollDice());
        }
        return this.wisdom;
    }

    int getCharisma() {
        if (this.charisma==0){
            this.charisma = ability(rollDice());
        }
        return this.charisma;
    }

    int getHitpoints() {
        return 10 + modifier(this.constitution);
    }
}
