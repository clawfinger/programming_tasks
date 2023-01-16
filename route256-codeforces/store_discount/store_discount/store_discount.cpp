// store_discount.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <unordered_map>

int main()
{
    int testCount = 0;
    std::cin >> testCount;
    for (int i = 0; i < testCount; i++) {
        int total = 0;
        std::cin >> total;
        std::unordered_map<int, int> counter;
        for (int j = 0; j < total; j++) {
            int product = 0;
            std::cin >> product;
            counter[product]++;
        }
        int sum = 0;
        for (const auto& entry : counter) {
            int product = entry.first;
            int amount = entry.second;
            if (amount < 3) {
                sum += amount * product;
            }
            else {
                int leftovers = amount % 3;
                sum += leftovers * product;

                int packagesCount = amount / 3;
                sum += packagesCount * 2 * product;
            }
        }
        std::cout << sum << std::endl;
    }
}

