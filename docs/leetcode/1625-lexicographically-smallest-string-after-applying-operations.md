---
tags:
- String
- Breadth-First Search
---


# 1625. Lexicographically Smallest String After Applying Operations

=== "C++"

    $Time: O(N^2)$

    $Space: O(N)$

    ```c++
    class Solution {
    public:
        string findLexSmallestString(string s, int a, int b) {
            int n = s.size();
            vector<int> v(n); // s - '0'
            vector<int> m; // min 

            for (int i = 0; i < n; i++)
                v[i] = s[i]-'0';
            m = v;

            int r = 0; // rotate: O(lcm(b, n) / b) < O(n)
            do {
                if (b % 2)
                    add(v, 0, a); // n / 2
                add(v, 1, a);     // n / 2
                m = min(m,v);     // n
                rotate(v, b);     // 2n
                r = (r + b) % n;
            } while (r);


            string ans;
            for (int i = 0; i < n; i++)
                ans += m[i] + '0';

            return ans;
        }

        // add a to interval index many times to let s[i] minimize
        void add(vector<int> &s, int i, int a) {
            int diff = 10 - s[i]; // 1,3,7,9

            if (a % 2 == 0)       // 2,4,6,8
                diff += (s[i] % 2);

            if (a == 5)           // 5
                diff = (s[i] >= 5) ? 5 : 0;

            for (; i < s.size(); i+=2)
                s[i] = (s[i] + diff) % 10;
        }

        void rotate(vector<int> &s, int b) {
            reverse(s.begin(), s.end());
            reverse(s.begin(), s.begin()+b);
            reverse(s.begin()+b, s.end());
        }
    };
    ```