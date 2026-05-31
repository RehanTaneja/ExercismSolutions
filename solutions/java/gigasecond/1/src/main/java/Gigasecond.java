import java.time.LocalDate;
import java.time.LocalDateTime;

public class Gigasecond {

    private LocalDate ld;
    private LocalDateTime ldt;
    
    public Gigasecond(LocalDate moment) {
        this.ld = moment;
    }

    public Gigasecond(LocalDateTime moment) {
        this.ldt = moment;
    }

    public LocalDateTime getDateTime() {
        if (ld==null){
            return ldt.plusSeconds(1000000000);
        }else{
            return ld.atStartOfDay().plusSeconds(1000000000);
        }
    }
}
