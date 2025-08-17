class Solution:
    def divisorSubstrings(self, num: int, k: int) -> int:
        result = 0
        str_num = str(num)
        for i in range(len(str_num) - k):
            sub_num = int(str_num[i: i + k])

            if num % sub_num == 0 and sub_num > 0:
                result += 1

        return result


if __name__ == '__main__':
    assert Solution().divisorSubstrings(240, 2) == 2
    assert Solution().divisorSubstrings(430043, 2) == 2

