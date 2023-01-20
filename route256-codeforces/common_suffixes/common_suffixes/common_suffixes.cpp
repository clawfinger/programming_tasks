// common_suffixes.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>
#include <unordered_map>
#include <map>

using dict = std::unordered_map<std::string, std::vector<std::string>>;

void subsctrings(const std::string& source, dict& dict) {
    //auto it = source.begin();
    //std::advance(it, 1);
    for (auto it = source.begin(); it != source.end(); it++) {
        dict[std::string(it, source.end())].push_back(source);
    }
}

int main()
{
    int dictSize = 0;
    std::cin >> dictSize;
    dict dictionary;
    for (int i = 0; i < dictSize; i++) {
        std::string dictEntry;
        std::cin >> dictEntry;
        subsctrings(dictEntry, dictionary);
    }
    int wordsNum = 0;
    std::cin >> wordsNum;
    std::vector<std::string> words;
    for (int i = 0; i < wordsNum; i++) {
        std::string word;
        std::cin >> word;
        words.push_back(word);
    }
    for (const auto word : words) {
        auto it = word.begin();
        std::advance(it, 1);
        //std::map<int, std::vector<std::string>> res;
        bool found = false;
        for (; it != word.end(); it++) {
            std::string suffix(it, word.end());
            auto wordFromDictIt = dictionary.find(suffix);
            if (wordFromDictIt != dictionary.end()) {
                //res[suffix.size()].push_back(*wordFromDictIt->second.begin());
                for (const auto& dictWord : wordFromDictIt->second) {
                    if (dictWord != word) {
                        std::cout << dictWord << std::endl;
                        found = true;
                        break;
                    }
                }
            }
            if (found) {
                break;
            }
        }

        if (!found) {
            std::cout << *dictionary.begin()->second.begin() <<std::endl;
        }
    }
}
