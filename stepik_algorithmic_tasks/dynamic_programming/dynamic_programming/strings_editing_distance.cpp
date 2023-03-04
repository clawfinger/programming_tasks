#include "strings_editing_distance.h"
#include <unordered_map>
#include <sstream>
std::unordered_map<std::string, int> cache;

int distance_recursive(const std::string& str1, const std::string& str2, int i, int j, std::stringstream* ss) {
	*ss << i << " " << j;
	std::string key{ ss->str() };
	ss->str("");
	ss->clear();

	auto it = cache.find(key);
	if (it != cache.end()) {
		return it->second;
	}
	if (i == 0) {
		return j;
	}
	if (j == 0) {
		return i;
	}
	int insertion = distance_recursive(str1, str2, i, j - 1, ss) + 1;
	int deletion = distance_recursive(str1, str2, i - 1, j, ss) + 1;
	int match = distance_recursive(str1, str2, i - 1, j - 1, ss);
	int mismatch = distance_recursive(str1, str2, i - 1, j - 1, ss) + 1;

	int res;
	if (str1[i - 1] == str2[j - 1]) {
		res =  std::min(insertion, std::min(deletion, match));
	}
	else {
		res =  std::min(insertion, std::min(deletion, mismatch));
	}
	cache[key] = res;
	return res;
}

int edit_distance(const std::string& str1, const std::string& str2) {
	std::stringstream ss;
	return distance_recursive(str1, str2, str1.size(), str2.size(), &ss);
}
