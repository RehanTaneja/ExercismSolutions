import java.math.BigInteger;

class Grains {

    BigInteger grainsOnSquare(final int square) {
        if (square<=0 || square>64){
            throw new IllegalArgumentException("square must be between 1 and 64");
        }
        BigInteger a = BigInteger.valueOf(1);
        for (int i=1;i<square;i++){
            a = a.multiply(BigInteger.valueOf(2));
        }
        return a;
    }

    BigInteger grainsOnBoard() {
        BigInteger a = BigInteger.valueOf(0);
        for(int i=1;i<=64;i++){
            a = a.add(grainsOnSquare(i));
        }
        return a;
    }

}
