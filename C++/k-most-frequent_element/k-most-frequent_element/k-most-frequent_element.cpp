// https://leetcode.com/problems/top-k-frequent-elements/

#include <iostream>
#include <vector>
#include <algorithm>
#include <map>

std::vector<int> topKFrequent(std::vector<int>& nums, int k) {
    std::map<int, int> counter;
    std::map<int, std::vector<int>, std::greater<int>> revCounter;
    for (auto entry : nums) {
        counter[entry]++;
    }
    std::vector<int> res;
    for (const auto& entry : counter) {
        revCounter[entry.second].push_back(entry.first);
    }
    for (const auto& entry : revCounter) {
        for (int i : entry.second) {
            res.push_back(i);
            if (res.size() == k) {
                return res;
            }
        }

    }
    return res;
}

int main()
{
    std::vector<int> nums{ 4,1,-1,2,-1,2,3 };
    auto res = topKFrequent(nums, 2);
    std::cout << "Hello World!\n";
}
