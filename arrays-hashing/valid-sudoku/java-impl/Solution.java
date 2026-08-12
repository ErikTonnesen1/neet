import java.util.HashSet;
import java.util.Set;
import java.util.HashMap;
import java.util.ArrayList;
import java.util.Objects;

public class Solution {
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

        Solution sol = new Solution();

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

        System.out.println("Is Board Valid? : " + sol.isValidSudoku(board2));

    }

    /**
     * Traverses the board place-by-place, adding each value to a row, column, 3x3
     * square HashSet
     * For each index, the value is checked for duplicates in each HashSet
     * 
     * If after traversing the entire board, no duplicate is found
     * 
     * @param board The sudoku board to be checked
     * @return boolean Returns true if no duplicates found in the entire board,
     *         false otherwise.
     */
    public boolean isValidSudoku(char[][] board) {
        Set<Character>[] rows = new Set[9];
        Set<Character>[] cols = new Set[9];
        HashMap<Square, HashSet<Character>> sqrs = new HashMap<Square, HashSet<Character>>();

        for (int x = 0; x < board.length; x++) {
            for (int y = 0; y < board[0].length; y++) {
                if (board[x][y] == '.') {
                    continue;
                }

                Square currSqr = new Square(x/3, y/3);
                if (rows[x] != null && rows[x].contains(board[x][y]) ||
                        cols[y] != null && cols[y].contains(board[x][y]) ||
                        sqrs.get(currSqr) != null && sqrs.get(currSqr).contains(board[x][y])) {
                    return false;
                }
                if (rows[x] == null) {
                    rows[x] = new HashSet<Character>();
                }
                if (cols[y] == null) {
                    cols[y] = new HashSet<Character>();
                }
                if (sqrs.get(currSqr) == null) {
                    sqrs.put(currSqr, new HashSet<Character>());
                }

                rows[x].add(board[x][y]);
                cols[y].add(board[x][y]);
                sqrs.get(currSqr).add(board[x][y]);

                
            }
        }

        // sqrs.forEach((square, set) -> {
        //     System.out.println("Square: [" + square.x + ", " + square.y + "]. Set: [" + set + "]");
        // });
        //
        return true;
    }

    private class Square {
        private int x, y;

        public Square(int x, int y) {
            this.x = x;
            this.y = y;
        }

        @Override
        public boolean equals(Object o) {
            if (this == o)
                return true;
            if (!(o instanceof Square))
                return false;
            Square square = (Square) o;
            return x == square.x && y == square.y;
        }

        @Override
        public int hashCode() {
            // return 31 * x * y;
            return Objects.hash(x, y);
        }
    }
}
