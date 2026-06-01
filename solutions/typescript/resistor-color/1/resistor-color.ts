export const colorCode = (color : string) : number => {
  const lowercasedString : string = color.toLowerCase();
  return COLORS.indexOf(lowercasedString);
}

export const COLORS = ['black','brown','red','orange','yellow','green','blue','violet','grey','white']
