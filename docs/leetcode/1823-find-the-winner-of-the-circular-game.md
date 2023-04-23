---
tags:
- Array
- Math
- Recursion
- Queue
- Simulation
- Unsolved
---


# 1823. Find the Winner of the Circular Game

=== "C++"

    $Time: O(N*K)$

    $Space: O(N)$

    ```c++
    class Solution {
    public:
        int findTheWinner(int n, int k) {
            queue<int> q;
            for (int i = 1; i <= n; i++)
                q.push(i);

            while(q.size() > 1) {
                int x = k;
                while (--x > 0) {
                    q.push(q.front());
                    q.pop();
                }
                cout << q.front() << endl;;
                q.pop();
            }
            
            return q.front();
        }
    };
    ```

=== "C++"

    $Time: O(N)$

    $Space: O(N)$

    ```c++
    class Solution {
    public:
        int findTheWinner(int n, int k) {
            return r(n, k) + 1;
        }

        int r(int n, int k) {
            if (n == 1)
                return 0;
            return (r(n-1, k) + k) % n;
        }
    };
    ```

=== "C++"

    $Time: O(N)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        int findTheWinner(int n, int k) {
            int ans = 0;
            for (int i = 1; i <= n; i++) {
                ans = (ans + k) % i;
            }
            return ans + 1;
        }
    };
    ```