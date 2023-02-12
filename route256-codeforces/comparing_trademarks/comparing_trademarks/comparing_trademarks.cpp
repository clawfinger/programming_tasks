// comparing_trademarks.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>
#include <set>
void doWork() {
    int stringsCount = 0;
    std::cin >> stringsCount;
    std::vector<std::string> strings(stringsCount);
    std::set<std::string> cache;
    for (int i = 0; i < stringsCount; i++) {
        std::string str;
        std::cin >> str;
        char prev = str[0];
        int sameCount = 1;
        std::string processed;
        processed.push_back(prev);
        for (int j = 1; j < str.size(); j++) {
            if (prev == str[j]) {
                sameCount++;
            }
            else {
                sameCount = 1;
            }
            prev = str[j];

            if (sameCount <= 2) {
                processed.push_back(str[j]);
            }
        }
        cache.insert(processed);
    }
    std::cout << cache.size() << std::endl;
}

int main()
{
    int testNum = 0;
    std::cin >> testNum;
    for (int i = 0; i < testNum; i++) {
        doWork();
    }
}

