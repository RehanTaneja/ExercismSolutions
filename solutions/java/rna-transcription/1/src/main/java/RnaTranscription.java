class RnaTranscription {

    String transcribe(String dnaStrand) {
        String a = "";
        for(int i = 0;i<dnaStrand.length();i++){
            switch (dnaStrand.charAt(i)){
                case 'G':
                    a += "C";
                    break;
                case 'C':
                    a += "G";
                    break;
                case 'T':
                    a += "A";
                    break;
                case 'A':
                    a += "U";
                    break;
            }
        }
        return a;
    }

}
