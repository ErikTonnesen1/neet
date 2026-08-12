package betterSolutions.TwoPointer;

public class twoPointer{

    /**
     * since numbers is already sorted, can use that to optimize time complexity to O(n) & space complexity O(1)
     * we can find the sum of the values at 0 and length-1 index
     *      - if sum > target --> move r pointer to the left 
     *      - if sum < target --> move l pointer to the right 
     */
    public static int[] twoSum(int[] numbers, int target){
        int l = 0; 
        int r = numbers.length - 1; 
        while(l < r){
            int sum = numbers[l] + numbers[r];

            if(sum > target) r--;
            if(sum < target) l++;
            if(sum == target) return new int[]{(l+1), (r+1)};
        }

        return new int[0];
    }
}
