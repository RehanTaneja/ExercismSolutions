class DifferenceOfSquaresCalculator {

    int computeSquareOfSumTo(int input) {
        int i = 1;
        int s = 0;
        while(i<=input){
            s += i;
            i++;
        }
        return s*s;
    }

    int computeSumOfSquaresTo(int input) {
        int s = 0;
        for(int i = 1;i<=input;i++){
            int d = i*i;
            s += d;
        }
        return s;
    }

    int computeDifferenceOfSquares(int input) {
        return computeSquareOfSumTo(input)-computeSumOfSquaresTo(input);
    }

}
