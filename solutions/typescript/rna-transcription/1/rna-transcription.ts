export function toRna(dna : string) : string{
  let rna : string = '';
  for (let i : number = 0;i < dna.length;++i){
    switch (dna[i]){
      case 'G':
        rna = rna + 'C';
        break;
      case 'C':
        rna = rna + 'G';
        break;
      case 'T':
        rna = rna + 'A';
        break;
      case 'A':
        rna = rna + 'U';
        break;
      default:
        throw new Error("Invalid input DNA.");
    }
  }
  return rna;
}
