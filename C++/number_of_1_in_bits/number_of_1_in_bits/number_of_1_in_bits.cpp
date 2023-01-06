//https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/565/
#include <iostream>

int hammingWeight(uint32_t n) {
    int res = 0;
    for (int i = 0; i < 32; i++) {
        if (n & 1 << i) {
            res++;
        }
    }
    return res;
}

int main()
{
    std::cout << "Hello World!\n";
}
