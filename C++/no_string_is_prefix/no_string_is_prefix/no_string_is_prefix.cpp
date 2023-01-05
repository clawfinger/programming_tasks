// no_string_is_prefix.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>
#include <set>
#include <algorithm>

using namespace std;

void noPrefix(vector<string> words) {
    unordered_map<string, set<string>> cache;
    set<string> original;
    for (const string& str : words) {
        if (cache.find(str) != cache.end()) {
            cout << "BAD SET" << endl << str << endl;
            return;
        }
        for (int i = 0; i < str.size() - 1; i++) {
            cache[str].insert(str.substr(0, i + 1));
        }
        original.insert(str);
    }
    //for (const string& str : words) {
    //    auto it = cache.find(str);
    //    if (it != cache.end()) {
    //        cout << "BAD SET" << endl << it->second[0] << endl;
    //        return;
    //    }
    //}
    for (auto& word : words) {
        const auto& wordSet = cache[word];
        std::set<string> intersection;
        std::set_intersection(original.begin(), original.end(), wordSet.begin(), wordSet.end(),
            std::inserter(intersection, intersection.begin()));
        if (intersection.size() != 0) {
            cout << "BAD SET" << endl << word << endl;
            return;
        }
    }
    cout << "GOOD SET";
}

int main()
{
    std::vector<string> set{ "abcd", "bcd", "abcde", "bcde" };
    noPrefix(set);
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
