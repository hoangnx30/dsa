from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]
        return []

    def twoSumOptimized(self, nums: List[int], target: int) -> List[int]:    
        dictionary = {}
        for idx, num in enumerate(nums):
            rest = target - num
            if rest in dictionary: 
                return [dictionary[rest], idx]
            dictionary[num] = idx
            
        return []
        
        
        
solution = Solution()
# Example usage:
nums = [2, 7, 11, 15]
target = 9
result = solution.twoSumOptimized(nums, target)
print(result)  # Output: [0, 1]
