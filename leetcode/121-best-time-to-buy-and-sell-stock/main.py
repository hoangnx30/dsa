from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        min_sellable = prices[0]
        max_profit = 0
        for sell in prices:
            max_profit = max(max_profit, sell - min_sellable)
            min_sellable = min(min_sellable, sell)

        return max_profit

    def maxProfit1(self, prices: List[int]) -> int:
        max_profit = 0
        left, right = 0, 1
        while right < len(prices):
            if prices[left] < prices[right]:
                max_profit = max(max_profit, prices[right] - prices[left])
            else:
                left = right

            right += 1

        return max_profit

    def maxProfit2(self, prices: List[int]) -> int:
        max_profit = 0
        for i in range(len(prices) - 1):
            for j in range(i + 1, len(prices)):
                max_profit = max(max_profit, prices[j] - prices[i])

        return max_profit


if __name__ == '__main__':
    assert Solution().maxProfit([1, 2, 3, 4]) == 3
    assert Solution().maxProfit1([1, 2, 3, 4]) == 3
    assert Solution().maxProfit2([1, 2, 3, 4]) == 3
