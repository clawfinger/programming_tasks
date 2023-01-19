

#include <iostream>
#include <vector>
#include <map>
#include <unordered_set>
#include <limits>

int main()
{
    int usersNum = 0;
    int entryNum = 0;
    std::cin >> usersNum >> entryNum;

    std::vector<std::unordered_set<int>> usersFriends(usersNum + 1);

    for (int i = 0; i < entryNum; i++) {
        int user1;
        int user2;
        std::cin >> user1 >> user2;

        usersFriends[user1].insert(user2);
        usersFriends[user2].insert(user1);
    }
    std::vector<std::map<int, int>> res(usersNum);
    for (int i = 1; i <= usersNum; i++) {
        const auto& friends = usersFriends[i];

        for (const auto& fren : friends) {
            const auto& friendFriends = usersFriends[fren];
            for (const auto& frenFren : friendFriends) {
                if (frenFren == i) {
                    continue;
                }
                if (friends.find(frenFren) == friends.end()) {
                    res[i - 1][frenFren]++;
                }
            }
        }
    }
    for (const auto& entry : res) {
        if (entry.empty()) {
            std::cout << "0";
        }
        else {
            int max = std::numeric_limits<int>::min();
            for (const auto& recommendation : entry) {
                if (recommendation.second > max) {
                    max = recommendation.second;
                }
            }
            for (const auto& recommendation : entry) {
                if (recommendation.second == max) {
                    std::cout << recommendation.first << " ";
                }
            }
        }
        std::cout << std::endl;
    }
}
