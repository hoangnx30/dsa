from typing import List


class Solution:
    def getRow(self, rowIndex: int) -> List[int]:

        if rowIndex == 0:
            return [1]

        if rowIndex == 1:
            return [1, 1]

        dp = [[1], [1, 1]]

        for i in range(2, rowIndex + 1):
            sub = [1]
            old_sub = dp[i - 1]
            for j in range(len(old_sub) - 1):
                sub.append(old_sub[j] + old_sub[j + 1])

            sub.append(1)
            dp.append(sub)

        return dp[rowIndex]


if __name__ == '__main__':
    assert Solution().getRow(0) == [1]
    assert Solution().getRow(1) == [1, 1]
    assert Solution().getRow(2) == [1, 2, 1]

