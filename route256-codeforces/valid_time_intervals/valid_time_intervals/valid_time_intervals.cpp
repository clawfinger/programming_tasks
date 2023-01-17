
#include <iostream>
#include <sstream>
#include <tuple>
#include <vector>
#include <algorithm>

using tuple = std::tuple<int, int, int>;
using pair = std::pair<tuple, tuple>;

pair parse(const std::string& str) {
    std::stringstream ss(str);

    int firstHours;
    int firstMins;
    int firstSecs;

    std::string tmp;
    std::getline(ss, tmp, ':');
    std::istringstream(tmp) >> firstHours;
    std::getline(ss, tmp, ':');
    std::istringstream(tmp) >> firstMins;
    std::getline(ss, tmp, '-');
    std::istringstream(tmp) >> firstSecs;

    int secondHours;
    int secondMins;
    int secondSecs;

    std::getline(ss, tmp, ':');
    std::istringstream(tmp) >> secondHours;
    std::getline(ss, tmp, ':');
    std::istringstream(tmp) >> secondMins;
    std::getline(ss, tmp);
    std::istringstream(tmp) >> secondSecs;
    return std::make_pair(tuple{ firstHours, firstMins, firstSecs }, tuple{ secondHours, secondMins, secondSecs });
}

bool checkValid(const tuple& entry) {
    const auto& hours = std::get<0>(entry);
    if (hours < 0 || hours > 23) {
        return false;
    }
    const auto& minutes = std::get<1>(entry);
    if (minutes < 0 || minutes > 59) {
        return false;
    }
    const auto& seconds = std::get<2>(entry);
    if (seconds < 0 || seconds > 59) {
        return false;
    }
    return true;
}

bool checkOrder(const pair& entry) {
    if (entry.first > entry.second) {
        return false;
    }
    return true;
}

void work() {
    int setsCount = 0;
    std::cin >> setsCount;
    std::vector<pair> pairs;
    for (int i = 0; i < setsCount; i++) {
        std::string entry;
        std::cin >> entry;
        
        pair current = parse(entry);
        pairs.push_back(current);

    }
    std::sort(pairs.begin(), pairs.end(), [&](const pair& left, const pair& right) ->bool {
        return left.first < right.first;
        });
    
    for (int i = 0; i < pairs.size(); i++) {
        if (!checkValid(pairs[i].first)) {
            std::cout << "NO" << std::endl;
            return;
        }
        if (!checkValid(pairs[i].second)) {
            std::cout << "NO" << std::endl;
            return;
        }
        if (!checkOrder(pairs[i])) {
            std::cout << "NO" << std::endl;
            return;
        }
        if (i + 1 < pairs.size()) {
            if (pairs[i].second >= pairs[i + 1].first) {
                std::cout << "NO" << std::endl;
                return;
            }
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

