import java.util.Random;

class CaptainsLog {

    private static final char[] PLANET_CLASSES = new char[]{'D', 'H', 'J', 'K', 'L', 'M', 'N', 'R', 'T', 'Y'};

    private Random random;

    CaptainsLog(Random random) {
        this.random = random;
    }

    char randomPlanetClass() {
        int n = random.nextInt(PLANET_CLASSES.length);
        return PLANET_CLASSES[n];
    }

    String randomShipRegistryNumber() {
        int n = random.nextInt(9000)+1000;
        return "NCC-"+n;
    }

    double randomStardate() {
        double d = random.nextDouble(1000) + 41000;
        return d;
    }
}
