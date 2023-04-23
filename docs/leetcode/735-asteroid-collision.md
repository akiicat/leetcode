---
tags:
- Array
- Stack
---


# 735. Asteroid Collision

=== "C++"

    $Time: O(N)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        vector<int> asteroidCollision(vector<int>& asteroids) {

            vector<int> res;

            for (int i = 0; i < asteroids.size(); i++) {
                int &a = asteroids[i];
                if (res.empty() || res.back() < 0 || a > 0) {
                    res.push_back(a);
                } else if (res.back() <= -a) {
                    if (res.back() < -a) i--;
                    res.pop_back();
                }
            }

            return res;
        }
    };
    ```