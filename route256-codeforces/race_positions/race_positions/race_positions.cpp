
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
    std::vector<int> places(runnersCount +1); // key - runner, value - place
    std::priority_queue<int, std::vector<int>, std::greater<int>> times;
    for (int i = 1; i <= runnersCount; i++) {
        int time;
        std::cin >> time;
        if (runnersCount == 1) {
            std::cout << 1 << std::endl;
            return;
        }
        runnersData[time].push_back(i);
        times.push(time);
    }
    int place = 1;
    int samePlace = 1;
    int prevTime = 0;
    bool first = true;
    while (!times.empty()) {
        int time = times.top();
        times.pop();
        if (first) {
            const auto& runners = runnersData[time];
            for (auto run : runners) {
                places[run] = place;
            }
            prevTime = time;
            first = false;
        }
        else {
            if (time - prevTime <= 1) {
                samePlace++;
            }
            else {
                place += samePlace;
                samePlace = 1;
            }
            const auto& runners = runnersData[time];
            for (auto run : runners) {
                std::cout << place << " ";
            }
            prevTime = time;
        }

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
