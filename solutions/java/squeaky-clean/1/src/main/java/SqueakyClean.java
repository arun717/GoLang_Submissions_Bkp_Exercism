class SqueakyClean {
    static String clean(String identifier) {
        StringBuilder sb = new StringBuilder();
        boolean capNext = false;
        for(int i = 0; i < identifier.length(); i++) {
            char c = identifier.charAt(i);
            if (c == ' ') {
                sb.append('_');
            } else if (c == '-') {
                capNext = true;
            } else if (!Character.isLetterOrDigit(identifier.charAt(i)) 
                       && identifier.charAt(i) != '-') {
                // do nothing
            } else if (c == '4') {
                sb.append('a');
            } else if (c == '3') {
                sb.append('e');
            } else if (c == '0') {
                sb.append('o');
            } else if (c == '1') {
                sb.append('l');
            } else if (c == '7') {
                sb.append('t');
            } else {
                if (capNext) {
                    sb.append(Character.toUpperCase(c));
                    capNext = false;
                } else {
                    sb.append(c);
                }
            }
        }
        return sb.toString();
    }
}