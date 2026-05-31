class NaturalNumber {
    private int num; 
    
    NaturalNumber(int number) {
        if (number <=0){
            throw new IllegalArgumentException("You must supply a natural number (positive integer)");
        }
        this.num = number;
    }

    Classification getClassification() {
        int s = 0;
        for (int i = 1;i<this.num;i++){
            if (this.num%i==0){
                s+=i;
            }
        }
        if (s==this.num){
            return Classification.PERFECT;
        }else if (s>this.num){
            return Classification.ABUNDANT;
        }else{
            return Classification.DEFICIENT;
        }
    }
}
