import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

class AppointmentScheduler {
    public LocalDateTime schedule(String appointmentDateDescription) {
        DateTimeFormatter parser = DateTimeFormatter.ofPattern("MM/dd/yyyy HH:mm:ss");
        LocalDateTime answer = LocalDateTime.parse(appointmentDateDescription,parser);
        return answer;
    }

    public boolean hasPassed(LocalDateTime appointmentDate) {
        LocalDateTime current = LocalDateTime.now();
        return appointmentDate.isBefore(current);
    }

    public boolean isAfternoonAppointment(LocalDateTime appointmentDate) {
        LocalDate b = appointmentDate.toLocalDate();
        LocalDate e = appointmentDate.toLocalDate();
        LocalDateTime bd = b.atTime(12,0,0);
        LocalDateTime ed = e.atTime(18,0,0);
        return (appointmentDate.equals(bd) || appointmentDate.isAfter(bd)) && appointmentDate.isBefore(ed);
    }

    public String getDescription(LocalDateTime appointmentDate) {
        String answer = "You have an appointment on ";
        switch (""+appointmentDate.getDayOfWeek()){
            case "MONDAY":
                answer += "Monday, ";
                break;
            case "TUESDAY":
                answer += "Tuesday, ";
                break;
            case "WEDNESDAY":
                answer += "Wednesday, ";
                break;
            case "THURSDAY":
                answer += "Thursday, ";
                break;
            case "FRIDAY":
                answer += "Friday, ";
                break;
        }
        switch (""+appointmentDate.getMonth()){
            case "JANUARY":
                answer += "January ";
                break;
            case "FEBRUARY":
                answer += "February ";
                break;
            case "MARCH":
                answer += "March ";
                break;
            case "APRIL":
                answer += "April ";
                break;
            case "MAY":
                answer += "May ";
                break;
            case "JUNE":
                answer += "June ";
                break;
            case "JULY":
                answer += "July ";
                break;
            case "AUGUST":
                answer += "August ";
                break;
            case "SEPTEMBER":
                answer += "September ";
                break;
            case "OCTOBER":
                answer += "October ";
                break;
            case "NOVEMBER":
                answer += "November ";
                break;
            case "DECEMBER":
                answer += "December ";
                break;
        }
        answer += appointmentDate.getDayOfMonth()+", "+appointmentDate.getYear()+", at ";  
        DateTimeFormatter f = DateTimeFormatter.ofPattern("h:mm a");
        answer += appointmentDate.format(f)+".";
        return answer;
    }

    public LocalDate getAnniversaryDate() {
        LocalDate now = LocalDate.now();
        return LocalDate.of(now.getYear(),9,15);
    }
}
