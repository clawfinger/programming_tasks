// string_editor.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <stack>

void deleteLast(std::string& s, int k) {
    s = s.substr(0, s.size() - k);
}

void printCharAtIndex(std::string s, int k) {
    std::cout << s[k - 1] << std::endl;
}

int main()
{
    int opCount = 0;
    std::cin >> opCount;
    std::string str;
    std::stack<std::string> stateStack;
    for (int i = 0; i < opCount; i++) {
        int op;
        std::cin >> op;
        switch (op) {
        case 1:
        {
            std::string append;
            std::cin >> append;
            stateStack.push(str);
            str.append(append);
            break;
        }
        case 2:
        {
            int k;
            std::cin >> k;
            stateStack.push(str);
            deleteLast(str, k);
        }
            break;
        case 3:
            int k;
            std::cin >> k;
            std::cout << str[k - 1] << std::endl;
            break;
        case 4:
            str = stateStack.top();
            stateStack.pop();
            break;
        default:
            return 1;
        }

    }
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
