#include "binary_search.h"

int bs(const std::vector<int>& arr, int n, int from, int to) {
	if (from > to) {
		return -1;
	}
	int mid = from + (to - from) / 2;
	if (arr[mid] == n) {
		return mid;
	}
	if (arr[mid] > n) {
		return bs(arr, n, from, mid - 1);
	}
	else {
		return bs(arr, n, mid + 1, to);
	}
}

int binarySearch(const std::vector<int>& arr, int n) {
	return bs(arr, n, 0, arr.size());
}

int binarySearchIterative(const std::vector<int>& arr, int n) {
	int left = 0;
	int right = arr.size();

	while (left <= right) {
		int mid = left + (right - left) / 2;

		if (arr[mid] == n) {
			return mid;
		}
		if (arr[mid] < n) {
			left = mid + 1;
		}
		else {
			right = mid;
		}
	}
	return -1;
}