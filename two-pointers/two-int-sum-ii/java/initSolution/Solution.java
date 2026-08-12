import java.util.List;
import java.util.ArrayList;

public class Solution {
    public static void main(String[] args){
        Solution sol = new Solution();
        int[] numbers = {1,2,3,4};
        int target = 3;

        int[] numbers2 = {-1, 0};
        int target2 = -1;

        // int[] solution = sol.twoSum(numbers, target);
        int[] solution = sol.twoSum(numbers2, target2);

        for(int x : solution){
            System.out.println(x);
        }



    } 


    public int[] twoSum(int[] numbers, int target){

        // int end = 0;
        // for(int i=numbers.length - 1; i >= 0; i--){
        //     if(numbers[i] <= target){
        //         end = i; 
        //         break;
        //     }
        // }

        int end = numbers.length - 1;
        for(int x=end; x >= 0; x--){
            for(int z =0; z < end; z++){
                if(numbers[z] + numbers[x] > target) continue;
                if(numbers[z] + numbers[x] == target) return new int[]{(z+1), (x+1)};
            }
        }

        return new int[0];
    }
}
