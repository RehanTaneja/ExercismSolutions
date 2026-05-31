class Badge {
    public String print(Integer id, String name, String department) {
        String answer = "";
        if (id!=null){
            answer += "["+id+"] - ";
        }
        answer += name + " - ";
        if (department==null){
            department = "owner";
        }
        answer += department.toUpperCase();
        return answer;
    }
}
