import java.util.List;
import java.util.ArrayList;
import java.util.Collections;

class HandshakeCalculator {

    int convertToBinary(int number) {
        int multiplier = 1;
        int ans = 0;
        while (number>0) {
            int digit = number%2;
            number = number/2;
            ans += multiplier*digit;
            multiplier *= 10;
        }
        return ans;
    }

    List<Signal> calculateHandshake(int number) {
        int binary = convertToBinary(number);
        System.out.println(binary);
        int cnt = 0;
        List<Signal> ans = new ArrayList<>();
        while (binary>0 && cnt<5) {
            int d = binary%10;
            binary = binary/10;
            if (d==1) {
                if (cnt==0) {
                    ans.add(Signal.WINK);
                } else if (cnt==1) {
                    ans.add(Signal.DOUBLE_BLINK);
                } else if (cnt==2) {
                    ans.add(Signal.CLOSE_YOUR_EYES);
                } else if (cnt==3) {
                    ans.add(Signal.JUMP);
                } else if (cnt==4) {
                    Collections.reverse(ans);
                } 
            }
            cnt++;
        }
        return ans;
    }

}
