from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1

        return - 1

    def binarySearch(self, nums: List[int], target: int) -> int:
        return self.recursive(0, len(nums) - 1, nums, target)

    def recursive(self, left: int, right: int, nums: List[int], target: int) -> int:
        mid = (left + right) // 2

        if left > right:
            return - 1

        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            return self.recursive(mid + 1, right, nums, target)
        return self.recursive(left, mid - 1, nums, target)


if __name__ == '__main__':
    solution = Solution()
    assert solution.search(nums=[1, 2, 3, 4, 5, 6, 7], target=3) == 2
    assert solution.binarySearch(nums=[1, 2, 3, 4, 5, 6, 7], target=1) == 0
    assert solution.binarySearch(nums=[-1, 0, 3, 5, 9, 12], target=2) == -1
