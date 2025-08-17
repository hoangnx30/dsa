class Solution:
    def maximum69Number(self, num: int) -> int:
        str_num = str(num)
        res = ""
        flag = False

        for i in range(len(str_num)):
            if str_num[i] == '6' and not flag:
                res += '9'
                flag = True
                continue

            res += str_num[i]

        return int(res)


if __name__ == '__main__':
    test = Solution()
    assert test.maximum69Number(9669) == 9969
