// sliding_window_maximum.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>
#include <deque>
#include <algorithm>

std::vector<int> maxSlidingWindow(std::vector<int>&& nums, int k) {
    std::deque<int> window;
    std::vector<int> res;
    int left = 0;
    for (int right = 0; right < nums.size(); right++) {
        while (!window.empty() && nums[window.back()] < nums[right]) {
            window.pop_back();
        }
        window.push_back(right);

        if (left > window.front()) {
            window.pop_front();
        }
        if (right + 1 >= k) {
            res.push_back(nums[window.front()]);
            left++;
        }
    }

    return res;
}

int main()
{
    auto res = maxSlidingWindow({ 1,3,-1,-3,5,3,6,7 }, 3);
}
