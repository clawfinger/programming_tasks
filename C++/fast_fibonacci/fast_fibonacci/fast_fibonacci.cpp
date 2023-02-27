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

int FibonacciLastDigit(int num) {
    if (num <= 1) {
        return num;
    }
    int prev = 0;
    int current = 1;
    for (int i = 2; i <= num; i++) {
        int tmp = prev;
        prev = current;
        current = (current + tmp) % 10;
    }
    return current;
}

int main()
{
    std::cout << FibonacciLastDigit(10);
}

