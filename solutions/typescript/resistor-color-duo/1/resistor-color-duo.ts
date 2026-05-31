export function decodedValue(colors : string[]) : number {
  const COLOR : string[] = ['black','brown','red','orange','yellow','green','blue','violet','grey','white']
  const firstLowercase : string = colors[0].toLowerCase();
  const secondLowerCase : string = colors[1].toLowerCase();
  return (COLOR.indexOf(firstLowercase) * 10 ) + COLOR.indexOf(secondLowerCase);
}
