#include <iostream>
#include <unordered_map>

std::unordered_map<int, int> solved;

int fastFibonacci(int num) {
    if (num <= 1) {
        return num;
    }
    auto prevIt = solved.find(num - 1);
    int prev;
    int prevPrev;
    if (prevIt != solved.end()) {
        prev = prevIt->second;
    }
    else {
        prev = fastFibonacci(num - 1);
    }
    auto prevPrevIt = solved.find(num - 2);
    if (prevPrevIt != solved.end()) {
        prevPrev = prevPrevIt->second;
    }
    else {
        prevPrev = fastFibonacci(num - 2);
    }
    int res = prev + prevPrev;
    solved[num] = res;
    return res;
}

long long FibonacciLastDigit(long long num) {
    if (num <= 1) {
        return num;
    }
    long long prev = 0;
    long long current = 1;
    for (int i = 2; i <= num; i++) {
        int tmp = prev;
        prev = current;
        current = (current + tmp) % 10;
    }
    return current;
}

std::vector<long long> getPisanoPeriod(long long m) {
    std::vector<long long> res;

    res.push_back(0);
    res.push_back(1);

    int prev = 0;
    int current = 1;
    for (;;) {
        int tmp = prev;
        prev = current;
        current = (current + tmp) % m;
        if (prev == 0 && current == 1) {
            res.pop_back();
            return res;
        }
        res.push_back(current);
    }

    return res;
}

long long get_fibonacci_modulo(long long n, long long m) {
    if (n <= 1)
        return n;

    auto period = getPisanoPeriod(m);
    int periodIndex = n % period.size();
    return period[periodIndex];
}

int fibonacci_sum(long long n) {
    if (n <= 1)
        return n;

    auto period = getPisanoPeriod(10);
    int periodIndex = (n + 2) % period.size();
    int lastDigit = period[periodIndex];
    if (lastDigit == 0) {
        return 9;
    }
    else {
        return lastDigit - 1;
    }
}

long long get_fibonacci_partial_sum(long long from, long long to) {
    long long full = FibonacciLastDigit((to + 2) % 60);
    long long start = FibonacciLastDigit((from -1 + 2) % 60);
    int res = (full - start + 10) % 10;
    return res;
}

int fibonacci_sum_squares_naive(long long n) {
    int N = FibonacciLastDigit((n) % 60);
    int NPlusOne = FibonacciLastDigit((n + 1) % 60);
    return (N * NPlusOne) % 10;
}

int main()
{
    std::cout << get_fibonacci_modulo(10, 2);
}

