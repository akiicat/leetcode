---
tags:
- Array
- Stack
- Queue
- Simulation
---


# 1700. Number of Students Unable to Eat Lunch

=== "C++"

    $Time: O(N)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        int countStudents(vector<int>& students, vector<int>& sandwiches) {
            int count[2] = {};
            int n = students.size();

            for (auto &i : students)
                count[i]++;

            int i;
            for (i = 0; i < n && count[sandwiches[i]]; i++)
                count[sandwiches[i]]--;

            return n - i;
        }
    };
    ```