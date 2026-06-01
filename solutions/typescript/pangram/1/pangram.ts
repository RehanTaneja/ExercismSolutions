export function isPangram(str : string) : boolean {
  if (str.length<26){
    return false;
  }
  str = str.toLowerCase();
  const ctr : number[] = [];
  for(let i : number = 0;i<26;++i){
    ctr[i] = 0;
  }
  for(let i : number = 0;i<str.length;++i){
    const a : number = str.charCodeAt(i) - 97;
    if (a<26){
      ctr[a] = ctr[a] + 1;
    }
  }
  for(let i : number = 0;i<26;++i){
    if (ctr[i]==0){
      return false;
    }
  }
  return true;
}
