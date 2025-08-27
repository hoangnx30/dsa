from typing import List

class Solution:
    def areaOfMaxDiagonal(self, dimensions: List[List[int]]) -> int:
        result = 0
        max_diagonal = 0
        for i in range(len(dimensions)):
            width = dimensions[i][0]
            height = dimensions[i][1]
            diagonal = height * height + width * width

            area = width * height

            if diagonal > max_diagonal:
                max_diagonal = diagonal
                result = area

            if diagonal == max_diagonal:
                result = max(area, result)

        return result


if __name__ == '__main__':
    so = Solution()
    assert so.areaOfMaxDiagonal([[1, 2], [3, 4]]) == 12
