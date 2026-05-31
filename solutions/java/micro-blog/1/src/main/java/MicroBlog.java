class MicroBlog {
    public String truncate(String input) {
        if (input==null){
            return null;
        }
        int codePoints = input.codePointCount(0,input.length());
        if (codePoints<=5){
            return input;
        }
        int endIndex = input.offsetByCodePoints(0,5);
        return input.substring(0,endIndex);
    }
}
