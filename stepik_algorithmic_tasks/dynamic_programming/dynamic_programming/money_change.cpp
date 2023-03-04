#include "money_change.h"
#include <limits>
#include <algorithm>
#include <iostream>
#include <unordered_map>

std::unordered_map<int, int> cache;

// coins of 1 3 4
int get_change(int m) {
	auto it = cache.find(m);
	if (it != cache.end()) {
		return it->second;
	}
	if (m == 0) {
		return 0;
	}
	int arr[3]{ 1,3,4 };

	int res = std::numeric_limits<int>::max();
	for (int coin : arr) {
		if (coin <= m) {
			res = std::min(res, get_change(m - coin) + 1);
			cache[m] = res;
		}
	}
	return res;
}