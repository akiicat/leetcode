---
tags:
- Math
- Dynamic Programming
- Combinatorics
---


# 2400. Number of Ways to Reach a Position After Exactly k Steps

=== "C++"

    $Time: O(K)$

    $Space: O(1)$

    ```c++
    class Solution {
    public:
        const int mod = 1e9+7;

        int numberOfWays(int startPos, int endPos, int k) {
            int diff = abs(endPos - startPos);

            if (diff > k || (k - diff) % 2)
                return 0;

            return ncr(k,(k-diff)/2);
        }

        unsigned long long int expo(unsigned long long int x, int y, int p) {
            unsigned long long int res = 1;
            x %= p;

            while (y > 0) {
                if (y & 1)
                    res = (res * x) % p;
                y >>= 1;
                x = (x * x) % p;
            }

            return res;
        }

        int inverse(int n, int p) {
            return expo(n, p-2, p);
        }

        int ncr(int n, int k) {
            if (n <= 1) return 1;
            if (n - k < k) k = n - k;

            unsigned long long int res = 1;
            for (int i = 1; i <= k; i++) {
                res = (res * (n - i + 1)) % mod;
                res = (res * inverse(i,mod)) % mod;  // (res / i) % mod;
            }
            
            return res;
        }
    };
    ```