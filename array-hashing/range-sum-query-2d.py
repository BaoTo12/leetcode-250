class NumMatrix:
    def __init__(self, matrix):
        if not matrix or not matrix[0]:
            self.prefixSum = [[]]
            return

        m, n = len(matrix), len(matrix[0])
        print(m, n)
        self.prefixSum = [[0] * (n + 1) for _ in range(m + 1)]
        print(self.prefixSum)  # 6 x 6
        for i in range(m):
            for j in range(n):
                self.prefixSum[i+1][j+1] = (self.prefixSum[i][j+1]
                                            + self.prefixSum[i+1][j]
                                            - self.prefixSum[i][j]
                                            + matrix[i][j])
                print("self.prefixSum[i+1][j+1]", self.prefixSum[i+1][j+1])
                print()
        print(self.prefixSum)

    def sumRegion(self, row1, col1, row2, col2):
        return (self.prefixSum[row2+1][col2+1]
                - self.prefixSum[row2+1][col1]
                - self.prefixSum[row1][col2+1]
                + self.prefixSum[row1][col1])


matrix = [
    [3, 0, 1, 4, 2],
    [5, 6, 3, 2, 1],
    [1, 2, 0, 1, 5],
    [4, 1, 0, 1, 7],
    [1, 0, 3, 0, 5]
]
num_matrix = NumMatrix(matrix)

sum_1 = num_matrix.sumRegion(2, 1, 4, 3)
# print(f"Tổng vùng từ (2, 1) đến (4, 3) là: {sum_1}")

# sum_2 = num_matrix.sumRegion(1, 1, 2, 2)
# print(f"Tổng vùng từ (1, 1) đến (2, 2) là: {sum_2}")

# sum_3 = num_matrix.sumRegion(0, 0, 1, 2)
# print(f"Tổng vùng từ (0, 0) đến (1, 2) là: {sum_3}")
