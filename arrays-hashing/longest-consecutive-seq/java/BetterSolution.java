import java.util.Arrays;

public class BetterSolution {

    public static void main(String[] args){

    }

    public int longestConescutive(int[] nums){
        Arrays.sort(nums);

        int seq = 1;
        int maxSeq = 1;
        for(int i = 0; i < nums.length-1; i++){
            if(nums[i+1] == nums[i]) continue;
            if(nums[i+1] == nums[i] + 1){
                seq++;
            } else{
                maxSeq = Math.max(maxSeq, seq);
                seq = 1;
            }
        }

        return Math.max(maxSeq, seq);
    }


}
