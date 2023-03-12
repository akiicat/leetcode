---
tags:
- Array
- Hash Table
- Math
- Binary Search
- Bit Manipulation
- Sorting
- Unsolved
---


# 268. Missing Number

=== "C++"

    $Time: O(N)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        int missingNumber(vector<int>& nums) {
            int res = nums.size();

            for (int i = 0; i < nums.size(); i++)
                res ^= nums[i] ^ i;

            return res;
        }
    };
    ```

=== "C++"

    $Time: O(N)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        int missingNumber(vector<int>& nums) {
            int s = accumulate(nums.begin(), nums.end(), 0);
            int n = nums.size();
            return ((n * (n+1)) / 2) - s;
        }
    };
    ```