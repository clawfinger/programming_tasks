
#include <iostream>
#include <map>
#include <vector>
#include <algorithm>
#include <set>
void testRun() {
    int runnersCount = 0;
    std::cin >> runnersCount;
    std::map<int, std::vector<int>> runnersData; //key - time, value - runner
    std::map<int, int> places; // key - runner, value - place
    std::vector<int> times;
    for (int i = 1; i <= runnersCount; i++) {
        int time;
        std::cin >> time;
        if (runnersCount == 1) {
            std::cout << 1 << std::endl;
            return;
        }
        runnersData[time].push_back(i);
        times.push_back(time);
    }
    std::sort(times.begin(), times.end());
    int place = 1;
    int samePlace = 1;
    int prevTime = 0;
    bool first = true;
    for (auto time : times) {
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
                places[run] = place;
            }
            prevTime = time;
        }

    }
    for (int i = 1; i <= runnersCount; i++) {
        std::cout << places[i] << " ";
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
