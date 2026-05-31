public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        int successRate;
        if (speed==0){
            successRate = 0;
        }else if (speed<=4){
            successRate = 100;
        }else if (speed<=8){
            successRate = 90;
        }else if (speed==9){
            successRate = 80;
        }else{
            successRate = 77;
        }
        double success = (double) successRate/100;
        return 221*success*speed;
    }

    public int workingItemsPerMinute(int speed) {
        double perHour = productionRatePerHour(speed);
        return (int) perHour/60;
    }
}
