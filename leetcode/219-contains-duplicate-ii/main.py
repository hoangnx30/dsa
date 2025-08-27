from typing import List


class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        seen = {}

        for idx, val in enumerate(nums):
            if val in seen and abs(seen[val] - idx) <= k:
                return True
            seen[val] = idx

        return False