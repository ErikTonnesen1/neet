import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class BetterSolution {

    public static void main(String[] args) {

        char[][] board = {
                { '1', '2', '.', '.', '3', '.', '.', '.', '.' },
                { '4', '.', '.', '5', '.', '.', '.', '.', '.' },
                { '.', '9', '8', '.', '.', '.', '.', '.', '3' },
                { '5', '.', '.', '.', '6', '.', '.', '.', '4' },
                { '.', '.', '.', '8', '.', '3', '.', '.', '5' },
                { '7', '.', '.', '.', '2', '.', '.', '.', '6' },
                { '.', '.', '.', '.', '.', '.', '2', '.', '.' },
                { '.', '.', '.', '4', '1', '9', '.', '.', '8' },
                { '.', '.', '.', '.', '8', '.', '.', '7', '9' } };

        BetterSolution sol = new BetterSolution();

        char[][] board2 = {
                { '1', '2', '.', '.', '3', '.', '.', '.', '.' },
                { '4', '.', '.', '5', '.', '.', '.', '.', '.' },
                { '.', '9', '1', '.', '.', '.', '.', '.', '3' },
                { '5', '.', '.', '.', '6', '.', '.', '.', '4' },
                { '.', '.', '.', '8', '.', '3', '.', '.', '5' },
                { '7', '.', '.', '.', '2', '.', '.', '.', '6' },
                { '.', '.', '.', '.', '.', '.', '2', '.', '.' },
                { '.', '.', '.', '4', '1', '9', '.', '.', '8' },
                { '.', '.', '.', '.', '8', '.', '.', '7', '9' }
        };

        System.out.println("Is Board Valid? : " + sol.isValidSudoku(board));
        System.out.println("Is Board Valid? : " + sol.isValidSudoku(board2));

    }

    public boolean isValidSudoku(char[][] board) {
        Map<Integer, HashSet<Character>> rows = new HashMap();
        Map<Integer, HashSet<Character>> cols = new HashMap();
        Map<String, HashSet<Character>> sqrs = new HashMap();

        for (int x = 0; x < 9; x++) {
            for (int y = 0; y < 9; y++) {
                if (board[x][y] == '.')
                    continue;
                String sqr = x / 3 + "," + y / 3;

                if (rows.computeIfAbsent(x, k -> new HashSet<Character>()).contains(board[x][y]) ||
                        cols.computeIfAbsent(y, k -> new HashSet<Character>()).contains(board[x][y]) ||
                        sqrs.computeIfAbsent(sqr, k -> new HashSet<Character>()).contains(board[x][y])) {
                    return false;
                }
                rows.get(x).add(board[x][y]);
                cols.get(y).add(board[x][y]);
                sqrs.get(sqr).add(board[x][y]);
            }
        }

        return true;
    }

}
