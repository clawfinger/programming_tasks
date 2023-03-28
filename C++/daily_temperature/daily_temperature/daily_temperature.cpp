// daily_temperature.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>
#include <stack>

std::vector<int> dailyTemperatures(std::vector<int>&& temperatures) {
    std::vector<int> res(temperatures.size(), 0);
    std::stack<std::pair<int, int>> st;
    for (int i = 0; i < temperatures.size(); ++i) {
        while (!st.empty() && st.top().first < temperatures[i]) {
            const auto& current = st.top();
            st.pop();
            res[current.second] = i - current.second;
        }
        st.push({ temperatures[i], i });
    }

    return res;
}

int main()
{
    auto res = dailyTemperatures({ 73,74,75,71,69,72,76,73 });
}

