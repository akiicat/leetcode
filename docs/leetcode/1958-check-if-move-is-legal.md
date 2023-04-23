---
tags:
- Array
- Matrix
- Enumeration
---


# 1958. Check if Move is Legal

=== "C++"

    $Time: O(N)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        bool checkMove(vector<vector<char>>& board, int rMove, int cMove, char color) {
            int dir[9] = { 1, 0, -1, 0, 1, 1, -1, -1, 1 };
            for (int d = 0; d < 8; d++) {
                for (int x = 1; ; x++) {
                    int r = rMove + x * dir[d];
                    int c = cMove + x * dir[d+1];
                    if (r < 0 || c < 0 || r >= 8 || c >= 8 || board[r][c] == '.') {
                        break;
                    }
                    if (board[r][c] == color) {
                        if (x > 1) {
                            return true;
                        }
                        break;
                    }
                }
            }
            return false;
        }
    };
    ```