#include <iostream>
#include <vector>

#include "binary_search.h"

int main()
{
    std::vector<int> arr{ 1,2,3,4,5,6,7,8,9 };
    std::cout << binarySearchIterative(arr, 2);
}
