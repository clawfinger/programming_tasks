#include <vector>
#include <iostream>

int getDenominations(const std::vector<int>& coins, int sum) {
    std::vector<std::vector<int>> combinations;
    for (int a = 0; a <= sum / coins[0]; a++) {
        for (int b = 0; b <= sum / coins[1]; b++) {
            for (int c = 0; c <= sum / coins[2]; c++) {
                for (int d = 0; d <= sum / coins[3]; d++) {
                    combinations.push_back({ a, b, c, d });
                }
            }
        }
    }
    int smallestNumber = std::numeric_limits<int>::max();
    for (const auto& combination : combinations) {
        int currentSum = 0;
        int numCoins = 0;
        for (int i = 0; i < combination.size(); i++) {      
            std::cout << combination[i] << " ";
            currentSum += combination[i] * coins[i];
            numCoins += combination[i];
        }
        std::cout << std::endl;
        if (currentSum == sum) {
            if (smallestNumber > numCoins) {
                smallestNumber = numCoins;
            }
        }
    }
    return smallestNumber;
}

int main()
{
    std::vector<int> coins{ 25, 20, 10, 5 };
    std::cout << getDenominations(coins, 40);
}

