from typing import List
from unittest import result


class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        left, right, zeros = 0, 0, 0

        for right in range(len(nums)):
            if nums[right] == 0:
                zeros += 1

    def longest_subarray_brute_force(self, nums: List[int]) -> int:

        result = 0

        for i in range(len(nums)):
            zeros = 0
            total_one = 0
            for j in range(i, len(nums)):

                if nums[j] == 0:
                    zeros += 1

                if nums[j] == 1:
                    total_one += 1

                if zeros == 2:
                    break

                if j == len(nums) - 1 and zeros == 0:
                    result = max(result, total_one - 1)
                    break

                result = max(result, total_one)

        return result


if __name__ == '__main__':
    assert Solution().longest_subarray_brute_force([1, 1, 1]) == 2
    assert Solution().longest_subarray_brute_force([1, 1, 0, 1]) == 3
    assert Solution().longest_subarray_brute_force([0, 1, 1, 1, 0, 1, 1, 0, 1]) == 5
