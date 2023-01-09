// https://leetcode.com/problems/max-points-on-a-line/

#include <iostream>
#include <vector>
#include <cmath>
#include <unordered_map>
#include <set>
#include <sstream>

int gcd(int a, int b)
{
    int result = std::min(a, b); // Find Minimum of a nd b
    while (result > 0) {
        if (a % result == 0 && b % result == 0) {
            break;
        }
        result--;
    }
    return result; // return gcd of a nd b
}

int maxPoints(std::vector<std::vector<int>>& points) {
    if (points.size() == 1) {
        return 1;
    }

    int max = std::numeric_limits<int>::min();

    std::unordered_map<std::string, std::set<std::pair<int, int>>> pointsAngles;
    for (int i = 0; i < points.size(); i++) {
        int maxForPoint = std::numeric_limits<int>::min();
        std::unordered_map < std::string, int> slopes;
        std::pair<int, int> first(points[i][0], points[i][1]);

        for (int j = 0; j < points.size(); j++) {
            if (i == j) {
                continue;
            }
            std::pair<int, int> second(points[j][0], points[j][1]);

            int y = second.second - first.second;
            int x =  second.first - first.first;

            int gdc = gcd(std::abs(x), std::abs(y));
            if (gdc != 0) {
                x /= gdc;
                y /= gdc;
            }


            if (x < 0 && y < 0) {
                x *= -1;
                y *= -1;
            }
            else if (x < 0 && y > 0) {
                x *= -1;
                y *= -1;
            }
            std::stringstream ss;
            ss << x << "-" << y;

            int current = slopes[ss.str()];
            current++;
            slopes[ss.str()] = current;
            if (current > maxForPoint) {
                maxForPoint = current;
            }
        }
        if (maxForPoint > max) {
            max = maxForPoint;
        }
    }
    return max + 1;
}

int main()
{
    std::vector<std::vector<int>> points{ {1,1 }, { 3,2 }, { 5,3 }, { 4,1 }, { 2,3 },{1, 4}};
    //std::vector<std::vector<int>> points{ {0,0 }, { 4,5 }, { 7,8 }, { 8,9 }, { 5,6 }, { 3,4 }, { 1,1 }};
    std::cout << maxPoints(points);
}

