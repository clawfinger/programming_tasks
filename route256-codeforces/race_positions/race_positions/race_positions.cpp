
#include <iostream>
#include <map>
#include <vector>
#include <algorithm>
#include <set>
#include <queue>

void testRun() {
    int runnersCount = 0;
    std::cin >> runnersCount;
    std::map<int, std::vector<int>> runnersData; //key - time, value - runner
    std::vector<int> positions(runnersCount + 1);
    for (int i = 1; i <= runnersCount; i++) {
        int time;
        std::cin >> time;
        if (runnersCount == 1) {
            std::cout << 1 << std::endl;
            return;
        }
        auto it = runnersData.find(time);
        if (it != runnersData.end()) {
            it->second.push_back(i);
            continue;
        }
        it = runnersData.find(time - 1);
        if (it != runnersData.end()) {
            it->second.push_back(i);
            continue;
        }
        it = runnersData.find(time + 1);
        if (it != runnersData.end()) {
            for (auto runner : it->second) {
                runnersData[time].push_back(runner);
            }
            runnersData.erase(time + 1);
            continue;
        }
        runnersData[time].push_back(i);
    }

    int pos = 0;
    int prev = 1;
    for (const auto& entry : runnersData) {
        for (auto runner : entry.second) {
            positions[runner] = pos + prev;
        }

        prev = entry.second.size();
        pos++;
    }

    for (int i = 1; i <= runnersCount; i++) {
        std::cout << positions[i] << " ";
    }
    std::cout << std::endl;
}

int main()
{
    int testCounts = 0;
    std::cin >> testCounts;
    for (int i = 0; i < testCounts; i++) {
        testRun();
    }
}
