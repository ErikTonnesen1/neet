package betterSolutions;
import betterSolutions.TwoPointer.*;

public class Main {
    public static void main(String[] args){
        int[] numbers = {1,2,3,4};
        int target = 3;

        int[] numbers2 = {-1, 0};
        int target2 = -1;

        // int[] solution = sol.twoSum(numbers, target);
        int[] solution = twoPointer.twoSum(numbers2, target2);

        for(int x : solution){
            System.out.println(x);
        }
    }
}
