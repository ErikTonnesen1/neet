public class Solution {

    public static void main(String[] args) {

        String s = "Was it a car or a cat I saw?";
        String s2 = "tab a cat";
        Solution sol = new Solution();
        System.out.println(sol.isPalindrome(s2));
    }

    public boolean isPalindrome(String s) {

        if (s.isBlank()) {
            return true;
        }

        s = s.toLowerCase();
        s = s.replaceAll(" ", "");
        s = s.replaceAll("[^a-z0-9]", "");
        s = s.trim();

        System.out.println("Sanitized String: " + s);

        char[] charArr = s.toCharArray();
        int lhs = 0;
        int rhs = s.length() - 1;
        ;

        while (rhs > lhs) {
            if (charArr[lhs] != charArr[rhs]) {
                return false;
            }
            lhs++;
            rhs--;
        }
        return true;
    }

}
