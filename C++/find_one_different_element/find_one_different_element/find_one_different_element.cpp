#include <iostream>
#include <vector>

int singleNumber(std::vector<int>& nums) {
    int res = 0;
    for (int& entry : nums) {
        res ^= entry;
    }
    return res;
}

int main()
{
    std::vector<int> test{ 2,2,4,4,1, 1};
    auto res = singleNumber(test);
}
