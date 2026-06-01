public class LogLine {

    private String line;
    public LogLine(String logLine) {
        this.line = logLine;
    }

    public LogLevel getLogLevel() {
        switch (this.line.substring(1,4)){
            case "TRC":
                return LogLevel.TRACE;
            case "DBG":
                return LogLevel.DEBUG;
            case "INF":
                return LogLevel.INFO;
            case "WRN":
                return LogLevel.WARNING;
            case "ERR":
                return LogLevel.ERROR;
            case "FTL":
                return LogLevel.FATAL;
        }
        return LogLevel.UNKNOWN;
    }

    public String getOutputForShortLog() {
        return ""+getLogLevel().get()+":"+this.line.substring(7);
    }
}
