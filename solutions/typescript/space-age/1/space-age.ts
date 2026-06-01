export function age(planet: string, seconds: number): number {
  if (seconds < 0){
    throw new Error("Invalid Input: seconds cannot be negative.");
  }
  let years : number = seconds / (60*60*24*365.25);
  switch (planet){
    case 'mercury':
      return Math.round((years / 0.2408467)*100)/100;
    case 'venus':
      return Math.round((years / 0.61519726)*100)/100;
    case 'earth':
      return Math.round(years*100)/100;
    case 'mars':
      return Math.round((years / 1.8808158)*100)/100;
    case 'jupiter':
      return Math.round((years / 11.862615)*100)/100;
    case 'saturn':
      return Math.round((years / 29.447498)*100)/100;
    case 'uranus':
      return Math.round((years / 84.016846)*100)/100;
    case 'neptune':
      return Math.round((years / 164.79132)*100)/100;
    default:
      throw new Error("Invalid input: invalid planet.");
  }
}
