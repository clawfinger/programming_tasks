// https://leetcode.com/problems/is-subsequence/

#include <iostream>
#include <string>

bool isSubsequence(std::string s, std::string t) {
    int sIdx = 0;
    for (int i = 0; i < t.size(); i++) {
        if (t[i] == s[sIdx]) {
            sIdx++;
        }
    }
    return sIdx == s.size();
}

int main()
{
    std::cout << isSubsequence("abc", "ahbgdc");
}
