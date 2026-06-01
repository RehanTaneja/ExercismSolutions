class Acronym {

    private String phrase;

    Acronym(String phrase) {
        this.phrase = phrase;
    }

    String get() {
        String a ="";
        if (this.phrase==null){
            return null;
        }
        int i = 0;
        if (!Character.isLetter(this.phrase.charAt(i))){
            i++;
        }
        while(i<this.phrase.length()){
            if (Character.isLetter(this.phrase.charAt(i))){
                if (i==0){
                    a += Character.toUpperCase(this.phrase.charAt(i));
                }else if(this.phrase.charAt(i-1)==' '||this.phrase.charAt(i-1)=='-'||this.phrase.charAt(i-1)=='_'){
                    a += Character.toUpperCase(this.phrase.charAt(i));
                }
            }
            i++;
        }
        return a;
    }

}
