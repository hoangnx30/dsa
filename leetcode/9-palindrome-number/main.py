class Solution:
    def isPalindrome(self, x: int) -> bool:
        return str(x) == str(x)[::-1]

    def isPalindrome2(self, x: int) -> bool:
        temp = x
        new = 0
        while temp > 0:
            new = new * 10 + temp % 10
            temp = temp // 10

        return new == x

if __name__ == "__main__":
    test = Solution()
    assert test.isPalindrome(121) == True
    assert test.isPalindrome(122) == False

    assert test.isPalindrome2(121) == True
    assert test.isPalindrome2(122) == False
