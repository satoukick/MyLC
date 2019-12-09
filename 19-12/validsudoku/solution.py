from typing import List


class Solution:
    def isValidSudoku(self, board: 'List[List[str]]') -> 'bool':
        return self.is_col_valid(board) and self.is_row_valid(board) and \
            self.is_square_valid(board)

    def is_unit_valid(self, unit) -> 'bool':
        u = [i for i in unit if i != '.']
        return len(set(u)) == len(u)

    def is_row_valid(self, board):
        for row in board:
            if not self.is_unit_valid(row):
                return False
        return True

    def is_col_valid(self, board):
        for col in list(zip(*board)):
            if not self.is_unit_valid(col):
                return False
        return True

    def is_square_valid(self, board):
        for i in (0, 3, 6):
            for j in (0, 3, 6):
                square = [board[x][y]
                          for x in range(i, i+3) for y in range(j, j+3)]
                if not self.is_unit_valid(square):
                    return False
        return True
