// https://leetcode.com/problems/gas-station/

#include <iostream>
#include <vector>

int getNextIndex(int initialIndex, int move, int maxIndex) {
    int nextIndex = initialIndex + move;
    return nextIndex % maxIndex;
}

//std::vector<int> gas{ 4,5,2,6,5,3 };
//std::vector<int> cost{ 3,2,7,3,2,9 };

int canCompleteCircuit(std::vector<int>& gas, std::vector<int>& cost) {
    for (int i = 0; i < gas.size(); i++) {
        int currentPos = i;
        int currentGas = gas[currentPos];

        for (;;) {
            // moving
            int currentCost = cost[currentPos];
            currentGas -= currentCost;

            // did we make it?
            if (currentGas >= 0) {

                // updating position
                currentPos++;
                currentPos = currentPos % gas.size();
                // full circle?
                if (currentPos == i) {
                    return i;
                }
                // filling tank
                currentGas += gas[currentPos];
            }
            else {
                break;
            }
        }
    }
    return -1;
}

int main()
{
    //std::vector<int> gas{ 1,2,3,4,5 };
    //std::vector<int> cost{ 3,4,5,1,2 };
    //std::vector<int> gas{ 2,3,4 };
    //std::vector<int> cost{ 3,4,3 };
    std::vector<int> gas{ 4,5,2,6,5,3 };
    std::vector<int> cost{ 3,2,7,3,2,9 };
    std::cout << canCompleteCircuit(gas, cost);
    std::cout << std::endl;
}