public class Say {

    public String say(long number) {
        if(number<0){
            throw new IllegalArgumentException("Number should not be less than 0.");
        }else if(number>999999999999L){
            throw new IllegalArgumentException("Number exceeds upper limit.");
        }
        String[] basics = {"zero","one","two","three","four","five","six","seven","eight","nine","ten","eleven","twelve","thirteen","fourteen","fifteen","sixteen","seventeen","eighteen","nineteen"};
        String[] tens = {"zero","ten","twenty","thirty","forty","fifty","sixty","seventy","eighty","ninety"};
        if (number<20){
            return basics[(int) number];
        } else if (number<100){
            String ans = "";
            long d = number%10;
            long d2 = number/10;
            ans = ans + tens[(int)d2];
            if (d!=0){
                ans = ans + "-" + basics[(int)d];
            }
            return ans;
        } else if(number<1000){
            long left = number%100;
            long d = number/100;
            String ans = basics[(int)d] + " hundred";
            if (left!=0){
                ans = ans + " " + say(left);
            }
            return ans;
        } else if (number<1000000){
            long d = number/1000;
            long left = number%1000;
            String ans = say(d) + " thousand";
            if(left!=0){
                ans = ans + " " + say(left);
            }
            return ans;
        } else if (number<1000000000){
            long d = number/1000000;
            long left = number%1000000;
            String ans = say(d) + " million";
            if(left!=0){
                ans = ans + " " + say(left);
            }
            return ans;
        } else{
            long d = number/1000000000;
            long left = number%1000000000;
            String ans = say(d) + " billion";
            if(left!=0){
                ans = ans + " " + say(left);
            }
            return ans;
        }
    }
}
