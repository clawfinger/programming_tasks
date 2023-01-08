// https://leetcode.com/explore/featured/card/top-interview-questions-easy/102/math/745/

#include <iostream>

bool isPowerOfThree(int n) {
    if (n == 0) {
        return false;
    }

    while (n != 1) {
        if (n % 3 != 0) {
            return false;
        }
        n = n / 3;
    }
    return true;
}

int main()
{
    std::cout << isPowerOfThree(1);
}

