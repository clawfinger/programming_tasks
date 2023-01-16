
#include <iostream>
#include <unordered_map>
#include <vector>
#include <algorithm>
#include <queue>

void pairsSlow() {
    int count = 0;
    std::cin >> count;
    std::unordered_map<int, std::priority_queue<int, std::vector<int>, std::greater<int> >> programmersStrength; // key - strength, value programmer;
    std::vector<int> strength(count, 0);
    for (int j = 0; j < count; j++) {
        int current = 0;
        std::cin >> current;
        strength[j] = current;
        programmersStrength[current].push(j);
    }
    std::sort(strength.begin(), strength.end());
    std::vector<std::pair<int, int >> pairs;
    for (int k = 1; k < strength.size(); k += 2) {
        int leftStr = strength[k - 1];
        int rightStr = strength[k];
        int firstProgrammer = programmersStrength[leftStr].top();
        programmersStrength[leftStr].pop();
        int secondProgrammer = programmersStrength[rightStr].top();
        programmersStrength[rightStr].pop();
        pairs.push_back({ std::min(firstProgrammer, secondProgrammer), std::max(firstProgrammer, secondProgrammer) });
    }
    std::sort(pairs.begin(), pairs.end(), [&](std::pair<int, int> left, std::pair<int, int> right) -> bool {
        return left.first < right.first;
        });
    for (const auto& pair : pairs) {
        std::cout << pair.first + 1 << " " << pair.second + 1 << std::endl;
    }
}

void pairs() {
    int count = 0;
    std::cin >> count;
    std::vector<int> strength(count, 0);
    for (int j = 0; j < count; j++) {
        int current = 0;
        std::cin >> current;
        strength[j] = current;
    }
    std::unordered_map<int, int> used;
    for (int i = 0; i < strength.size(); i++) {
        if (used.find(i) != used.end()) {
            continue;
        }
        int minDiff = std::numeric_limits<int>::max();
        int pair = -1;
        for (int j = i + 1; j < strength.size(); j++) {
            if (used.find(j) != used.end()) {
                continue;
            }
            int diff = std::abs(strength[i] - strength[j]);
            if (diff < minDiff) {
                minDiff = diff;
                pair = j;
            }
        }
        used[pair]++;
        std::cout << i + 1 << " " << pair + 1 << std::endl;
    }
}

int main()
{
    int testCount = 0;
    std::cin >> testCount;
    for (int i = 0; i < testCount; i++) {
        pairs();
    }
}
