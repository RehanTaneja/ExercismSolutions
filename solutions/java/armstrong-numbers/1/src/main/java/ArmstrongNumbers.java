import java.lang.Math;

class ArmstrongNumbers {

    boolean isArmstrongNumber(int numberToCheck) {
        int n = numberToCheck;
        if (numberToCheck==0){
            return true;
        }
        int num = numberToCheck;
        int numDigits = 0;
        while (num>0){
            num = num/10;
            numDigits++;
        }
        int s = 0;
        while (numberToCheck>0){
            int d = numberToCheck%10;
            s += Math.pow(d,numDigits);
            numberToCheck = numberToCheck / 10;
        }
        return s==n;
    }
}
