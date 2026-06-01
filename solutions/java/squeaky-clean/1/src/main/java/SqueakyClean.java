class SqueakyClean {
    static String clean(String identifier) {
        identifier =  identifier.replaceAll(" ","_");
        String last = "";
        int i = 0;
        boolean sawDash = false;
        while(i<identifier.length()){
            char c = identifier.charAt(i);
            if (c=='-'){
                sawDash = true;
                i++;
            }else{
                if (sawDash){
                    last = last + Character.toUpperCase(c);
                    sawDash = false;
                    i++;
                }else{
                    last = last + c;
                    i++;
                }
            }
        }
        identifier = last;
        identifier = identifier.replaceAll("4","a");
        identifier = identifier.replaceAll("3","e");
        identifier = identifier.replaceAll("0","o");
        identifier = identifier.replaceAll("1","l");
        identifier = identifier.replaceAll("7","t");
        String aakhri = "";
        for (int j=0;j<identifier.length();j++){
            char c = identifier.charAt(j);
            if (Character.isLetter(c)||c=='_'||Character.isDigit(c)){
                aakhri = aakhri + c;
            }
        }
        return aakhri;
    }
}
