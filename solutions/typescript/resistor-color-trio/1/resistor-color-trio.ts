export function decodedResistorValue(colors : string[]) : string {
  const COLOR : string[] = ['black','brown','red','orange','yellow','green','blue','violet','grey','white'];
  let n : number = (COLOR.indexOf(colors[0].toLowerCase()) * 10) + COLOR.indexOf(colors[1].toLowerCase());
  if (n === 0){
    return '0 ohms';
  }
  let nZeros : number = COLOR.indexOf(colors[2].toLowerCase());
  if (n % 10 === 0){
    n = n/10;
    nZeros = nZeros + 1;
  }
  let unit : string = '';
  if (nZeros < 3){
    unit = 'ohms';
  }else if (nZeros < 6){
    unit = 'kiloohms';
    nZeros = nZeros - 3;
  }else if (nZeros < 9){
    unit = 'megaohms';
    nZeros = nZeros - 6;
  }else {
    unit = 'gigaohms';
    nZeros = nZeros - 9;
  }

  for (let i : number = 0; i < nZeros; ++i){
    n = n * 10;
  }
  return n.toString() + ' ' + unit;
}
