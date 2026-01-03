def rotate(matrix: list[list[int]]) -> None:
    # 48
    n = len(matrix)
    for i in range(n // 2):
        for j in range(i, n - i - 1):
            # Pattern match to swap 4 positions
            match (matrix[i][j], 
                   matrix[j][n-1-i], 
                   matrix[n-1-i][n-1-j], 
                   matrix[n-1-j][i]):
                case (top, right, bottom, left):
                    matrix[i][j] = left
                    matrix[j][n-1-i] = top
                    matrix[n-1-i][n-1-j] = right
                    matrix[n-1-j][i] = bottom

def test_rotate():
    matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
    rotate(matrix)
    assert matrix == [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

if __name__ == "__main__":
    test_rotate()
