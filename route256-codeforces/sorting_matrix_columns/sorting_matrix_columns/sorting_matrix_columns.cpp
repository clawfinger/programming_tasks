// sorting_matrix_columns.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>
#include <algorithm>

void sortRows(std::vector<std::vector<int>>& matrix, int op) {
    auto predicate = [&](const std::vector<int>& left, const std::vector<int>& right) -> bool {
        return left[op - 1] < right[op - 1];
    };
    std::stable_sort(matrix.begin(), matrix.end(), predicate);
}

void print(std::vector<std::vector<int>>& matrix) {
    for (int i = 0; i < matrix.size(); i++) {
        for (int j = 0; j < matrix[i].size(); j++) {
            std::cout << matrix[i][j] << " ";
        }
        std::cout << std::endl;
    }
}

void work() {
    int rows = 0;
    int columns = 0;
    std::cin >> rows >> columns;
    std::vector<std::vector<int>> matrix(rows, std::vector<int>(columns));
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < columns; j++) {
            std::cin >> matrix[i][j];
        }
    }
    int opsNum = 0;
    std::cin >> opsNum;
    for (int i = 0; i < opsNum; i++) {
        int op = -1;
        std::cin >> op;
        sortRows(matrix, op);
    }
    print(matrix);
}

int main()
{
    int testCount = 0;
    std::cin >> testCount;
    for (int i = 0; i < testCount; i++) {
        work();
    }
}
