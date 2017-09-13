# 36.Valid Sudoku
class Solution(object):
    def isValidSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: bool
        """
        return self.is_valid_row(board) and self.is_valid_square(board) and self.is_valid_column(board)

    def is_valid_unit(self, unit):
        before = [i for i in unit if i != '.']
        return len(before) == len(set(before))


    def is_valid_row(self,board):
        for row in board:
            if not self.is_valid_unit(row):
                return False
        return True


    def is_valid_column(self,board):
        for col in zip(*board):
            if not self.is_valid_unit(col):
                return False
        return True

    def is_valid_square(self,board):
        for i in (0,3,6):
            for j in (0,3,6):
                square = [board[x][y] for x in range(i,i+3) for y in range(j,j+3)]
                if not self.is_valid_unit(square):
                    return False
        return True
