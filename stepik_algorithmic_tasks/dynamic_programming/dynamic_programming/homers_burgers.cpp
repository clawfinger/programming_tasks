#include "homers_burgers.h"
#include <vector>

int burgers_recursive(int n, int m, int time) {
	if (time == 0) {
		return 0;
	}

	int countN = -1;
	int countM = -1;
	if (n <= time) {
		countN = burgers_recursive(n, m, time - n);
	}
	if (m <= time) {
		countM = burgers_recursive(n, m, time - m);
	}
	if (countN == -1 && countM == -1) {
		return -1;
	}
	return std::max(countM, countN) + 1;
}

int burgers_iterative(int n, int m, int time) {
	std::vector<int> dp(time + 1, -1);
	dp[0] = 0;
	for (int i = 1; i <= time; i++) {
		int countN = -1;
		int countM = -1;
		if (n <= i) {
			countN = dp[i - n];
		}
		else {
			countN = -1;
		}
		if (m <= i) {
			countM = dp[i - m];
		}
		else {
			countM = -1;
		}
		if (countN == -1 && countM == -1) {
			dp[i] = - 1;
		}
		else {
			dp[i] = std::max(countM, countN) + 1;
		}
	}
	return dp[time];
}

int burgers(int n, int m, int time) {
	return burgers_iterative(n, m, time);
}
