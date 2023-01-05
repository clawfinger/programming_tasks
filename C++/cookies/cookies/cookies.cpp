// cookies.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>
#include <queue>
using namespace std;

int cookies(int k, vector<int> A) {
    std::priority_queue<int, std::vector<int>, std::greater<int>> queue;
    for (int i : A) {
        queue.push(i);
    }
    if (queue.empty() || queue.size() < 2) {
        return -1;
    }
    int res = 0;
    while (queue.top() < k) {
        if (queue.size() < 2) {
            break;
        }
        int first = queue.top();
        queue.pop();
        int second = queue.top();
        queue.pop();
        int merged = first + 2 * second;
        res++;
        queue.push(merged);
    }
    if (queue.top() < k) {
        return -1;
    }
    return res;

}

int main()
{
    std::cout << cookies(7, std::vector<int>{1, 2, 3, 9, 10, 12});
}

// Run program: Ctrl + F5 or Debug > Start Without Debugging menu
// Debug program: F5 or Debug > Start Debugging menu

// Tips for Getting Started: 
//   1. Use the Solution Explorer window to add/manage files
//   2. Use the Team Explorer window to connect to source control
//   3. Use the Output window to see build output and other messages
//   4. Use the Error List window to view errors
//   5. Go to Project > Add New Item to create new code files, or Project > Add Existing Item to add existing code files to the project
//   6. In the future, to open this project again, go to File > Open > Project and select the .sln file
