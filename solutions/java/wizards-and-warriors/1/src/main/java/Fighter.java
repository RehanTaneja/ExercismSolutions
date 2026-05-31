class Fighter {

    boolean isVulnerable() {
        return true;
    }

    int getDamagePoints(Fighter fighter) {
        return 1;
    }
}

// TODO: define the Warrior class
class Warrior extends Fighter{
    
    @Override
    boolean isVulnerable(){
        return false;
    }

    @Override
    int getDamagePoints(Fighter fighter){
        if (fighter.isVulnerable()){
            return 10;
        }
        return 6;
    }

    @Override
    public String toString(){
        return "Fighter is a Warrior";
    }
}

// TODO: define the Wizard class
class Wizard extends Fighter{

    private boolean preparedSpell = false;

    @Override
    public String toString(){
        return "Fighter is a Wizard";
    }

    void prepareSpell(){
        this.preparedSpell = true;
    }

    @Override
    boolean isVulnerable(){
        if (this.preparedSpell==true){
            return false;
        }
        return true;
    }

    @Override
    int getDamagePoints(Fighter fighter){
        if (this.preparedSpell==true){
            return 12;
        }
        return 3;
    }
}