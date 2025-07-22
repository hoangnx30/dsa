import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    public static int[] twoSum(int[] nums, int target) {

        Map<Integer, Integer> hashMap = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            if (hashMap.containsKey(target - nums[i])) {
                return new int[]{i, hashMap.get(target - nums[i])};
            } else {
                hashMap.put(nums[i], i);
            }
        }

        return new int[]{};
    }

    public static void main(String[] args) {
        int[] nums = new int[]{1, 2, 3, 4};
        int target = 6;
        int[] result = twoSum(nums, target);
        System.out.printf("with nums: %s, target: %d => result: %s", Arrays.toString(nums), target, Arrays.toString(result));
    }
}
