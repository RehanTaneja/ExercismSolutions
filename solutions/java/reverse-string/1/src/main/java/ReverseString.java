class ReverseString {

    String reverse(String inputString) {
        String n = "";
        int i = inputString.length()-1;
        while (i>=0){
            n += inputString.charAt(i);
            i -= 1;
        }
        return n;
    }
  
}
