class ResistorColorDuo {
    int value(String[] colors) {
        int a = 1;
        switch(colors[0]){
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
        switch(colors[1]){
            case "black":
                a += 0;
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
        return a;
    }
}
