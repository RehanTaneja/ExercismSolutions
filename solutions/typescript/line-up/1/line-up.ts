export function format(name: string, number: number): unknown {
  let lastDigit = number%10;
  let secondlastDigit = ((number%100) - lastDigit)/10;
  if (lastDigit==1 && secondlastDigit!=1) {
    return name + ", you are the " + number + "st customer we serve today. Thank you!"
  } else if (lastDigit==2 && secondlastDigit!=1){
    return name + ", you are the " + number + "nd customer we serve today. Thank you!"
  } else if (lastDigit==3 && secondlastDigit!=1){
    return name + ", you are the " + number + "rd customer we serve today. Thank you!"
  } else{
    return name + ", you are the " + number + "th customer we serve today. Thank you!"
  }
}
