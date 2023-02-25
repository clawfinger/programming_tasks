
#include <iostream>
#include <unordered_map>

int main()
{
    int count = 0;
    std::cin >> count;
    for (int i = 0; i < count; i++) {
        std::unordered_map<int, int> cache;

        for (int i = 0; i < 10; i++) {
            int entry = 0;
            std::cin >> entry;
            cache[entry]++;
        }
        if (cache[1] == 4 && cache[2] == 3 && cache[3] == 2 && cache[4] == 1) {
            std::cout << "YES" << std::endl;
        }
        else {
            std::cout << "NO" << std::endl;
        }
    }

}
