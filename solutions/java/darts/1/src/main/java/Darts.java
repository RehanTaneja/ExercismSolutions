import java.lang.Math;

class Darts {
    int score(double xOfDart, double yOfDart) {
        double dist = Math.pow(xOfDart,2) + Math.pow(yOfDart,2);
        if (dist>100){
            return 0;
        }else if (dist>25){
            return 1;
        }else if (dist>1){
            return 5;
        }
        return 10;
    }
}
