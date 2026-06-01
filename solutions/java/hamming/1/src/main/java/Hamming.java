public class Hamming {
    private String left;
    private String right;
    
    public Hamming(String leftStrand, String rightStrand) {
        if (leftStrand.length()!=rightStrand.length()){
            throw new IllegalArgumentException("strands must be of equal length");
        }
        this.left = leftStrand;
        this.right = rightStrand;
    }

    public int getHammingDistance() {
        int len;
        if (this.left.length()!=this.right.length()){
            throw new IllegalArgumentException("strands must be of equal length");
        }else{
            len = this.left.length();
        }
        int hd = 0;
        for (int i =0;i<len;i++){
            if (this.left.charAt(i)!=this.right.charAt(i)){
                hd++;
            }
        }
        return hd;
    }
}
