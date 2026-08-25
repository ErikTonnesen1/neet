import java.util.Deque;
import java.util.HashMap;
import java.util.Map;
import java.util.ArrayDeque;

public class ValidParenthesis {
    public static void main(String[] args){
        ParenthesisChecker checker = new ParenthesisChecker();
        String[] inputs = {"()", "[]", "{}", "[{()}]", "(((())))", "]", "}}", "({}})"};

        for (String input : inputs){
            System.out.printf("Is %s a valid Parenthesis? --> %b\n", input, checker.isValid(input));
        }

    }

    public static class ParenthesisChecker {
        private boolean isValid(String s) {
            String[] characters = s.split("");
            Deque<String> stack = new ArrayDeque<>();
            Map<String, String> parenthesisPairs = new HashMap<>(Map.of("(", ")", "[", "]", "{", "}"));

            for(int i = 0; i < characters.length; i++){
                String character = characters[i];

                if (parenthesisPairs.containsKey(character) == true){
                    stack.push(character);
                } else {
                    if (stack.size() == 0){
                        return false;
                    }
                    String lastOpenCharacter = stack.pop();
                    String requiredClosingCharacter = parenthesisPairs.get(lastOpenCharacter); 

                    boolean validParenthesisMatch = character.equals(requiredClosingCharacter);
                    if (!validParenthesisMatch) {
                        return false;
                    }
                }
            }

            if (stack.size() == 0){
                return true;
            } else {
                return false;
            }
        }
    }
}
