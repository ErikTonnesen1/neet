import java.util.TreeSet;

public class Solution {
    public static void main(String[] args) {
        Solution sol = new Solution();

        // int[] nums = { 2, 20, 4, 10, 3, 4, 5 };
        // System.out.println("Output: " + sol.longestConsecutive(nums));

        int[] nums2 = { 0, 3, 2, 5, 4, 6, 1, 1 };
        
        int[] nums3 = {9,1,4,7,3,-1,0,5,8,-1,6};
        System.out.println("Output: " + sol.longestConsecutive(nums3));


    }

    /**
     *
     * Must first find smallest value
     * Must then traverse array (Keeping track of last value), if we find a value
     * that's +1 greater,
     * then we add to the consecutive count and set as last value
     *
     * Need to do in O(n)
     *
     * Could traverse, add all values into hashmap @ the index of their value; then
     * traverse hashmap and find longest consecutive count
     */
    public int longestConsecutive(int[] nums) {
        TreeSet<Integer> sortedNums = new TreeSet<Integer>();

        if(nums.length == 0){
            return 0;
        }

        for (int num : nums) {
            sortedNums.add(num);
        }
        System.out.println("Sorted Nums: " + sortedNums.toString());

        int largestSequence = 0;
        int currSequenceCount = 1;
        int lastLargest = sortedNums.first();
        for (int num : sortedNums) {
            if (num - lastLargest == 1) {
                currSequenceCount++;
            } else if (num == lastLargest) {
                continue;
            } else {
                if (currSequenceCount > largestSequence) {
                    largestSequence = currSequenceCount;
                }
                currSequenceCount = 1;
            }

            System.out.printf("Checkpoint: LargestSeq: %d, currentSeqCount: %d, lastLargest: %d\n", largestSequence,
                    currSequenceCount, lastLargest);

            lastLargest = num;
        }

        return currSequenceCount > largestSequence ? currSequenceCount : largestSequence;

    }
}
