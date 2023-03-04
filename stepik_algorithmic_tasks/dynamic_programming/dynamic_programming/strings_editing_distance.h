#pragma once
#include <string>
#include <sstream>

int distance_recursive(const std::string& str1, const std::string& str2, int i, int j, std::stringstream* ss);

int edit_distance(const std::string& str1, const std::string& str2);