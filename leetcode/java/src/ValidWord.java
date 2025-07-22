import java.util.ArrayList;


public class ValidWord {
    public static boolean isValid(String word) {
        if (word.length() < 3) {
            return false;
        }

        System.out.println(word);

        int vovels = 0;
        int consonants = 0;

        for (char c : word.toCharArray()) {

            if (!Character.isLetterOrDigit(c)) {
//                System.out.println(word);
                return false;
            }

            if (Character.isLetter(c)) {
                char lowerCase = Character.toLowerCase(c);
                if (lowerCase == 'a' || lowerCase == 'e' || lowerCase == 'i' || lowerCase == 'u' || lowerCase == 'o') {
                    vovels++;
                } else {
                    consonants++;
                }
            }
        }

        return vovels > 0 && consonants > 0;
    }

    public static void main(String[] args) {
        String word = "234Adas";
        ArrayList<Integer> list = new ArrayList<>();


        boolean result = isValid(word);

        System.out.printf("word: %s. Result: %b \n", word, result);

        word = "3$a";
        result = isValid(word);

        System.out.printf("word: %s. Result: %b \n", word, result);

        word = "aya";
        result = isValid(word);

        System.out.printf("word: %s. Result: %b \n", word, result);

        word = "XV2";
        result = isValid(word);

        System.out.printf("word: %s. Result: %b \n", word, result);


    }
}
