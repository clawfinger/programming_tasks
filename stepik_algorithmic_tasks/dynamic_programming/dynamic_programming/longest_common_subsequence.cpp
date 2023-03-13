#include "longest_common_subsequence.h"

int lcs2_recursive(std::vector<int>& a, std::vector<int>& b, int i, int j) {
	if (i == 0 || j == 0) {
		return 0;
	}
	int insertion = lcs2_recursive(a, b, i, j - 1);
	int deletion = lcs2_recursive(a, b, i - 1, j);
	int equal = lcs2_recursive(a, b, i - 1, j - 1) + 1;
	int res = std::max(insertion, deletion);
	if (a[i - 1] == b[j - 1]) {
		int equal = lcs2_recursive(a, b, i - 1, j - 1) + 1;
		res = std::max(res, equal);
	}
	return res;
}

int lcs2(std::vector<int>& a, std::vector<int>& b) {
	return lcs2_recursive(a, b, a.size(), b.size());
}

int lcs2_iterative(std::vector<int>& a, std::vector<int>& b) {
	std::vector<std::vector<int>> results(a.size() + 1, std::vector<int>(b.size() + 1));
	for (int i = 0; i <= a.size(); ++i) {
		for (int j = 0; j <= b.size(); ++j) {
			if (i == 0 || j == 0) {
				results[i][j] = 0;
				continue;
			}
			int deletion = results[i - 1][j];
			int insertion = results[i][j - 1];
			results[i][j] = std::max(deletion, insertion);
			if (a[i - 1] == b[j - 1]) {
				results[i][j] = std::max(results[i][j], results[i - 1][j - 1] + 1);
			}
		}
	}
	return results[a.size()][b.size()];
}