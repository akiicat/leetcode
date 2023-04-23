---
tags:
- Math
- Simulation
---


# 2177. Find Three Consecutive Integers That Sum to a Given Number

=== "C++"

    $Time: O(1)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        vector<long long> sumOfThree(long long num) {
            if (num % 3 != 0)
                return {};

            long long x = num / 3;
            return {x-1,x,x+1};
        }
    };
    ```