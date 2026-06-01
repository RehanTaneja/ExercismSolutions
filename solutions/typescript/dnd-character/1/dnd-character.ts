export class DnDCharacter {
  public strength : number;
  public dexterity : number;
  public constitution : number;
  public intelligence : number;
  public wisdom : number;
  public charisma : number;
  public hitpoints : number;
  public static generateAbilityScore(): number {
    const start : number = 1;
    const max : number = 6;
    let min : number = 7;
    let result : number = 0;
    for (let i : number = 0;i<4;++i){
      const ans : number = Math.floor(Math.random()*(max-start+1)) + start;
      result = result + ans;
      if (ans<min){
        min = ans;
      }
    }
    return result - min;
  }

  public static getModifierFor(abilityValue: number): number {
    const ans : number = Math.floor((abilityValue-10)/2);
    return ans;
  }
  constructor(){
    this.strength = DnDCharacter.generateAbilityScore();
    this.dexterity = DnDCharacter.generateAbilityScore();
    this.constitution = DnDCharacter.generateAbilityScore();
    this.intelligence = DnDCharacter.generateAbilityScore();
    this.wisdom = DnDCharacter.generateAbilityScore();
    this.charisma = DnDCharacter.generateAbilityScore();
    this.hitpoints = 10 + DnDCharacter.getModifierFor(this.constitution);
  }
}
