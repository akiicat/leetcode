---
tags:
- Array
- Backtracking
- Bit Manipulation
- Matrix
---


# 980. Unique Paths III

=== "C++"

    **Time:** $O(N^2)$

    **Space:** $O(N^2)$

    ``` c++
    typedef pair<int,int> pii;

    class Solution {
    public:
        pii start;
        int res;

        int uniquePathsIII(vector<vector<int>>& grid) {

            // default start point
            int count = 1;

            for (int i = 0; i < grid.size(); i++) {
                for (int j = 0; j < grid[0].size(); j++) {
                    if (grid[i][j] == 1) {
                        start.first = i;
                        start.second = j;
                    }

                    if (grid[i][j] == 0) {
                        count++;
                    }
                }
            }

            res = 0;
            dfs(grid, start.first, start.second, count);
            return res;
        }
        
        void dfs(vector<vector<int>>& grid, int x, int y, int square) {

            if (!(x >= 0 && x < grid.size() && y >= 0 && y < grid[0].size() && grid[x][y] >= 0)) {
                return;
            }

            if (grid[x][y] == 2) {
                if (square == 0)
                    res++;
                return;
            }

            grid[x][y] ^= (1<<31);

            dfs(grid, x, y-1, square-1);
            dfs(grid, x-1, y, square-1);
            dfs(grid, x, y+1, square-1);
            dfs(grid, x+1, y, square-1);

            grid[x][y] ^= (1<<31);

            return;
        }
    };
    ```