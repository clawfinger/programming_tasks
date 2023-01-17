// task_sequenses.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>
#include <unordered_set>

void work() {
    int num = 0;
    std::cin >> num;
    std::vector<int> tasks(num);
    for (int i = 0; i < num; i++) {
        std::cin >> tasks[i];
    }
    int prev = -1;
    std::unordered_set<int> worked;

    for (int i = 0; i < tasks.size(); i++) {
        if (tasks[i] != prev) {
            if (worked.find(tasks[i]) != worked.end()) {
                std::cout << "NO" << std::endl;
                return;
            }
            prev = tasks[i];
            worked.insert(prev);
        }
    }
    std::cout << "YES" << std::endl;
}

int main()
{
    int testNum = 0;
    std::cin >> testNum;
    for (int i = 0; i < testNum; i++) {
        work();
    }
}
