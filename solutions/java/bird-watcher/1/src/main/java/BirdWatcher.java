
class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        return new int[]{0,2,5,3,7,8,4};
    }

    public int getToday() {
        return this.birdsPerDay[this.birdsPerDay.length-1];
    }

    public void incrementTodaysCount() {
        this.birdsPerDay[this.birdsPerDay.length-1]++; 
    }

    public boolean hasDayWithoutBirds() {
        for(int x: this.birdsPerDay){
            if (x==0){
                return true;
            }
        }
        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        if (numberOfDays>this.birdsPerDay.length){
            numberOfDays = this.birdsPerDay.length;
        }
        int cnt = 0;
        for (int i=0;i<numberOfDays;i++){
            cnt+=this.birdsPerDay[i];
        }
        return cnt;
    }

    public int getBusyDays() {
        int busy = 0;
        for(int i: this.birdsPerDay){
            if (i>=5){
                busy++;
            }
        }
        return busy;
    }
}
