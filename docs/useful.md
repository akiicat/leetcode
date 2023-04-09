# Cheat Sheet

## Module Operation

- 2400

### Exponential

```c++
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
```

### Inverse

```c++
int inverse(int n, int p) {
    return expo(n, p-2, p);
}
```

### Number of Combination

```c++
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
```