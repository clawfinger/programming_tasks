// check_brackets.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <stack>
#include <string>

using namespace std;

bool isOpening(char c) {
    return c == '{' || c == '(' || c == '[';
};

bool isClosing(char c) {
    return c == '}' || c == ')' || c == ']';
};

bool isMatching(char current, char check) {
    if (current == ')') {
        return check == '(';
    }
    if (current == '}') {
        return check == '{';
    }
    if (current == ']') {
        return check == '[';
    }
    return false;
}

string isBalanced(string s) {
    stack<char> stack;

    for (int i = 0; i < s.size(); i++) {
        if (isOpening(s[i])) {
            stack.push(s[i]);
        }
        else if (isClosing(s[i])) {
            if (stack.empty()) {
                return "NO";
            }
            char top = stack.top();
            if (!isMatching(s[i], top)) {
                return "NO";
            }
            stack.pop();
        }
    }
    if (!stack.empty()) {
        return "NO";
    }
    return "YES";
}


int main()
{
    string s{ "{()}" };
    std::cout << isBalanced(s);
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
