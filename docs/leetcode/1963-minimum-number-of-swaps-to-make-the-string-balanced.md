---
tags:
- Two Pointers
- String
- Stack
- Greedy
---


# 1963. Minimum Number of Swaps to Make the String Balanced

=== "C++"

    $Time: O(N)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        int minSwaps(string s) {
            int stack = 0;

            // count how many '[' left
            for (auto& ch : s) {
                if (ch == '[') {
                    stack++;
                } else if (stack > 0) {
                    stack--;
                }
            }
            
            return (stack+1)/2;
        }
    };
    ```