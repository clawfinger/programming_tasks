// playground.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>

int main()
{
    int testCount = 0;
    std::cin >> testCount;

    for (int i = 0; i < testCount; i++) {
        int a = 0;
        int b = 0;
        std::cin >> a >> b;
        std::cout << a + b << std::endl;
    }
}

