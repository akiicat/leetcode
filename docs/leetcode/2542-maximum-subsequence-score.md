---
tags:
- Array
- Greedy
- Sorting
- Heap (Priority Queue)
---


# 2542. Maximum Subsequence Score

=== "C++"

    $Time: O(Nlog(N))$

    $Space: O(N)$

    ```c++
    class Solution {
    public:
        long long maxScore(vector<int>& speed, vector<int>& efficiency, int k) {
            vector<pair<int, int>> es;
            for (int i = 0; i < speed.size(); i++) {
                es.push_back({efficiency[i], speed[i]});
            }

            sort(rbegin(es), rend(es));

            priority_queue<int, vector<int>, greater<int>> pq; // min queue
            long long sum = 0;
            long long res = 0;

            for (auto& [e, s] : es) {
                pq.push(s);
                sum += s;
                if (pq.size() > k) {
                    sum -= pq.top();
                    pq.pop();
                }
                if (pq.size() == k) {
                    res = max(res, sum*e);
                }
            }
            
            return res;
        }
    };
    ```