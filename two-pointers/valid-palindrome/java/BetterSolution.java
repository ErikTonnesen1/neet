import java.lang.StringBuilder;
import java.lang.Character;

public class BetterSolution {
    public static void main(String[] args) {

        String s = "Was it a car or a cat I saw?";
        String s2 = "tab a cat";
        BetterSolution sol = new BetterSolution();
        System.out.println(sol.isPalindrome(s));
        System.out.println(sol.isPalindrome(s2));
    }

    public boolean isPalindrome(String s) {
        StringBuilder newStr = new StringBuilder();
        for(char c : s.toCharArray()){
            if(Character.isLetterOrDigit(c)) newStr.append(Character.toLowerCase(c));
        }
        return newStr.toString().equals(newStr.reverse().toString());
    }
}
