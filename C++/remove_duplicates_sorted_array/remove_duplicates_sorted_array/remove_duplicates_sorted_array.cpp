// remove_duplicates_sorted_array.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>

int removeDuplicates(std::vector<int>& nums) {
    if (nums.size() == 1) {
        return 1;
    }
    std::vector<int> res;
    int current = nums[0];
    res.push_back(current);
    for (int i = 1; i < nums.size(); i++) {
        if (current != nums[i]) {
            res.push_back(nums[i]);
            current = nums[i];
        }
    }
    nums = res;
    return res.size();
    //int dupIndex = 0;
    //int mainIndex = 0;
    //while (mainIndex < nums.size())
    //{
    //    if (nums[dupIndex] == nums[mainIndex]) {
    //        mainIndex++;
    //    } else {
    //        dupIndex++;
    //        nums[dupIndex] = nums[mainIndex];
    //    }
    //}
    //return ++dupIndex;
}
}

int main()
{
    std::vector<int> nums{ 0, 0, 1, 1, 2, 3, 4, 5 };
    std::cout << removeDuplicates(nums);
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
