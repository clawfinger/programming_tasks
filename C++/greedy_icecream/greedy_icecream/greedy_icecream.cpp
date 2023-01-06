#include <iostream>
#include <vector>
#include <algorithm>

//https://leetcode.com/problems/maximum-ice-cream-bars/
int maxIceCream(std::vector<int>& costs, int coins) {
    std::sort(costs.begin(), costs.end());
    int res = 0;
    for (int i = 0; i < costs.size(); i++) {
        coins -= costs[i];
        if (coins < 0) {
            break;
        }
        res++;
    }
    return res;
}

#include <iostream>

int main()
{
    std::vector<int> costs{ 10,6,8,7,7,8 };
    std::cout << maxIceCream(costs, 5);
}

