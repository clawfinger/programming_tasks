#pragma once
#include <vector>
#include <algorithm>

inline double buy(const std::vector<int>& setNum, const std::vector<double>& setPrice, double singlePrice, int total){
	if (total == 0) {
		return 0;
	}
	double res = 0;
	double single = buy(setNum, setPrice, singlePrice, total - 1);
	res += single + singlePrice;

	for (int i = 0; i < setNum.size(); ++i) {
		if (total >= setNum[i]) {
			double current = buy(setNum, setPrice, singlePrice, total - setNum[i]);
			res = std::min(res, current + setPrice[i]);
		}
	}
	return res;
}

inline double buy_iterative(const std::vector<int>& setNum, const std::vector<double>& setPrice, double singlePrice, int total){
	std::vector<double> dp(total + 1, 0);
	dp[0] = 0;
	for (int i = 1; i <= total; i++) {
		double res = 0;
		double single = dp[i - 1];
		res += single + singlePrice;
		dp[i] = res;
		for (int j = 0; j < setNum.size(); ++j) {
			if (i >= setNum[j]) {
				double current = dp[i - setNum[j]];
				res = std::min(res, current + setPrice[j]);
				dp[i] = res;
			}
		}
	}
	return dp[total];
}