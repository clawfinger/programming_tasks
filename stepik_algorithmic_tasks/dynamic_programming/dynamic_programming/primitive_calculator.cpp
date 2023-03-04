#include "primitive_calculator.h"
#include <limits>
#include <map>

std::vector<int> optimal_sequence(int n) {
	std::map<int, int> done;
	done[1] = 0;
	for (int i = 2; i <= n; ++i) {
		int current = done[i - 1] + 1;
		if (i % 2 == 0) {
			current = std::min(current, done[i / 2] + 1);
		}
		if (i % 3 == 0) {
			current = std::min(current, done[i / 3] + 1);
		}
		done[i] = current;
	}
	std::vector<int> result;

	while (n > 1) {
		result.push_back(n);
		if (done[n] == done[n - 1] + 1) {
			n = n - 1;
		}
		else if (n % 2 == 0 && (done[n] == done[n / 2] + 1)) {
				n = n / 2;
		}
		else if (n % 3 == 0 &&  (done[n] == done[n / 3] + 1)) {
				n = n / 3;
			}
 	}
	result.push_back(1);

	return result;
}