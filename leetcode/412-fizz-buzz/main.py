from typing import List


class Solution:
    def fizzBuzz(self, nums: int) -> List[str]:
        result = []

        for num in range(1, nums + 1):
            if num % 3 == 0 and num % 5 == 0:
                result.append("FizzBuzz")
            elif num % 3 == 0:
                result.append("Fizz")
            elif num % 5 == 0:
                result.append("Buzz")
            else:
                result.append(str(num))

        return result


if __name__ == '__main__':
    assert Solution().fizzBuzz(15) == ["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz",
                                       "13", "14", "FizzBuzz"]
