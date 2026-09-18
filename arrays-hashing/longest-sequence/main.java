import java.util.Arrays;

public class main {
    public static void main(String[] args){
    int[] nums = {0,3,2,5,4,6,1,1};
    System.out.println(longestConsecutive(nums));
}

    public static int longestConsecutive(int[] nums) {
        if (nums.length == 0){
            return 0;
        }

        Arrays.sort(nums);
        int longestSeq = 0;
        int currentSeq = 0;
        for(int i = 0; i < nums.length - 1; i++){
            if(nums[i+1] - nums[i] == 1){
                currentSeq++;
            } else {
                currentSeq = 0;

            } 

            if (currentSeq > longestSeq) {
                longestSeq = currentSeq;
            }
        }

        return longestSeq;
    }

}
