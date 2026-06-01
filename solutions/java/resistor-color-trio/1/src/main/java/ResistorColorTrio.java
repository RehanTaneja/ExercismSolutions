class ResistorColorTrio {
    String label(String[] colors) {
        String s = "";
        int a = 1;
        int numZeros = 0;
        switch (colors[0]){
            case "black":
                a *= 0;
                break;
            case "brown":
                a *= 10;
                break;
            case "red":
                a *= 20;
                break;
            case "orange":
                a *= 30;
                break;
            case "yellow":
                a *= 40;
                break;
            case "green":
                a *= 50;
                break;
            case "blue":
                a *= 60;
                break;
            case "violet":
                a *= 70;
                break;
            case "grey":
                a *= 80;
                break;
            case "white":
                a *= 90;
                break;
        }
        switch (colors[1]){
            case "black":
                numZeros += 1;
                a = a/10;
                break;
            case "brown":
                a += 1;
                break;
            case "red":
                a += 2;
                break;
            case "orange":
                a += 3;
                break;
            case "yellow":
                a += 4;
                break;
            case "green":
                a += 5;
                break;
            case "blue":
                a += 6;
                break;
            case "violet":
                a += 7;
                break;
            case "grey":
                a += 8;
                break;
            case "white":
                a += 9;
                break;
        }
        s += a;
        String u = "";
        switch (colors[2]){
            case "black":
                break;
            case "brown":
                numZeros += 1;
                break;
            case "red":
                numZeros += 2;
                break;
            case "orange":
                numZeros += 3;
                break;
            case "yellow":
                numZeros += 4;
                break;
            case "green":
                numZeros += 5;
                break;
            case "blue":
                numZeros += 6;
                break;
            case "violet":
                numZeros += 7;
                break;
            case "grey":
                numZeros += 8;
                break;
            case "white":
                numZeros += 9;
                break;
        }
        if (numZeros>=9){
            u += "gigaohms";
            numZeros -= 9;
        }else if (numZeros>=6){
            u += "megaohms";
            numZeros -= 6;
        }else if (numZeros>=3){
            u += "kiloohms";
            numZeros -= 3;
        }else{
            u += "ohms";
        }
        for (int i=0;i<numZeros;i++){
            s += "0";
        }
        s += " " + u;
        if (a==0){
            return "0 ohms";
        }
        return s;
    }
}
