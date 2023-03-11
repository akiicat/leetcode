---
tags:
- String
- Greedy
---


# 1974. Minimum Time to Type Word Using Special Typewriter

=== "C++"

    $Time: O(N)$

    $Space: O(1)$

    ``` c++
    class Solution {
    public:
        int minTimeToType(string word) {
            int distance = word[0] - 'a';
            int count = min(distance, 26 - distance) + 1;

            for (int i = 0; i < word.size()-1; i++) {
                distance = abs(word[i+1] - word[i]);
                count += min(distance, 26 - distance) + 1;
            }

            return count;
        }
    };
    ```