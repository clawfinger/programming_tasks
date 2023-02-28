// GCD.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>

int gcd(int a, int b) {
    while (a > 0 && b > 0) {
        if (a >= b) {
            a = a % b;
        }
        else {
            b = b % a;
        }
    }
    return std::max(a, b);
}

int main()
{
    std::cout << "Hello World!\n";
}

