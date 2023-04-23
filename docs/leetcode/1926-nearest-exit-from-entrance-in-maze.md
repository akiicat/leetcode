---
tags:
- Array
- Breadth-First Search
- Matrix
---


# 1926. Nearest Exit from Entrance in Maze

=== "C++"

    $Time: O(N*M)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        int nearestExit(vector<vector<char>>& maze, vector<int>& entrance) {
            queue<pair<int,int>> q;
            q.push({entrance[0], entrance[1]});

            int m = maze.size();
            int n = maze[0].size();
            int moves = 1;

            int dir[5] = {0, 1, 0, -1, 0};

            maze[entrance[0]][entrance[1]] = '+';

            while(!q.empty()) {

                for (int l = q.size(); l > 0; l--) {
                    pair<int,int> &t = q.front();

                    for (int i = 0; i < 4; i++) {
                        int x = t.first + dir[i];
                        int y = t.second + dir[i+1];

                        if (x >= 0 && y >= 0 && x <= m-1 && y <= n-1 && maze[x][y] == '.') {
                            if (x == 0 || y == 0 || x == m-1 || y == n-1) {
                                return moves;
                            }
                            
                            maze[x][y] = '+';
                            q.push({x, y});
                        }
                    }
                    
                    q.pop();
                }

                moves++;
            }

            return -1;
        }
    };
    ```