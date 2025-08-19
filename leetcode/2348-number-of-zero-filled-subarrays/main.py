from typing import List


class Solution:
    def zeroFilledSubarray(self, nums: List[int]) -> int:
        count, streak = 0, 0
        for num in nums:
            if num == 0:
                streak += 1
                count += streak
            else:
                streak = 0

        return count

if __name__ == '__main__':
    assert Solution().zeroFilledSubarray([1, 2, 3, 0]) == 1
    assert Solution().zeroFilledSubarray([1, 0, 1, 0]) == 2
