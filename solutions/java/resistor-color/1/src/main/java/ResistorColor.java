class ResistorColor {
    int colorCode(String color) {
        color = color.toLowerCase();
        for(int i = 0;i<10;i++){
            if (colors()[i].equals(color)){
                return i;
            }
        }
        return -1;
    }

    String[] colors() {
        return new String[]{"black","brown","red","orange","yellow","green","blue","violet","grey","white"};
    }
}
