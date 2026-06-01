import java.util.Map;
import java.util.HashMap;

public class DialingCodes {

    Map<Integer,String> codes = new HashMap<>();

    public Map<Integer, String> getCodes() {
        return this.codes;
    }

    public void setDialingCode(Integer code, String country) {
        this.codes.put(code,country);
    }

    public String getCountry(Integer code) {
        return this.codes.get(code);
    }

    public void addNewDialingCode(Integer code, String country) {
        boolean answer = true;
        for (String s : this.codes.values()){
            if (s.equals(country)){
                answer = false;
            }
        }
        if (answer && !this.codes.containsKey(code)){
            this.codes.put(code,country);
        }
    }

    public Integer findDialingCode(String country) {
        for (int i : this.codes.keySet()){
            if (this.codes.get(i).equals(country)){
                return i;
            }
        }
        return null;
    }

    public void updateCountryDialingCode(Integer code, String country) {
        boolean answer = true;
        for (Integer i : this.codes.keySet()){
            if (this.codes.get(i).equals(country)){
                this.codes.remove(i);
                this.codes.put(code,country);
            }
        }
    }
}
