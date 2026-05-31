public class LogLevels {
    
    public static String message(String logLine) {
        String[] splitted = logLine.split(":");
        String message = splitted[1];
        return message.trim();
    }

    public static String logLevel(String logLine) {
        String[] splitted = logLine.split("]");
        String pOne = splitted[0];
        String level = pOne.substring(1);
        return level.toLowerCase();
    }

    public static String reformat(String logLine) {
        return message(logLine) + " (" + logLevel(logLine) + ")";
    }
}
