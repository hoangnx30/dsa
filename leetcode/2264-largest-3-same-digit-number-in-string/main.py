class Solution:
    def largestGoodInteger(self, num: str) -> str:
        result = ""

        for i in range(len(num) - 2):
            if num[i] == num[i + 1] == num[i + 2]:
                if result == "" or int(num[i]) > int(result[0]):
                    result = num[i] * 3

        return result


# Example usage:
solution = Solution()
result = solution.largestGoodInteger("6777133339")  # Output: "777"
print(result)  # Output: "777"
