// https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/

#include <iostream>
#include <vector>
#include <map>
#include <set>
#include <queue>

int minTime(int n, std::vector<std::vector<int>>& edges, std::vector<bool>& hasApple) {
    std::vector<std::vector<int>> connections(n);
    std::set<int> res;
    for (int i = 0; i < edges.size(); i++) {
        connections[edges[i][1]].push_back(edges[i][0]);
        connections[edges[i][0]].push_back(edges[i][1]);
    }
    std::set<int> visited;
    std::queue<int> queue;
    std::map<int, int> parents;

    queue.push(0);
    visited.insert(0);
    while (!queue.empty()) {
        int current = queue.front();
        queue.pop();
        const auto& adjascent = connections[current];

        for (int i = 0; i < adjascent.size(); i++) {
            int adj = adjascent[i];
            if (visited.find(adj) == visited.end()) {
                parents[adj] = current;
                queue.push(adj);
                visited.insert(adj);
            }
        }
    }

    for (int i = 0; i < hasApple.size(); i++) {
        if (hasApple[i]) {
            res.insert(i);
            int parent = parents[i];
            while (parent != 0) {
                res.insert(parent);
                int newParent = parents[parent];
                parent = newParent;
            }
        }
    }

    return res.size() * 2;
}

int main()
{
    //std::vector<std::vector<int>> edges{ {0,1 }, { 0,2 }, { 1,4 }, { 1,5 }, { 2,3 },{2, 6} };
    std::vector<std::vector<int>> edges{ {0,1 }, { 1,2 }, { 0,3 }};
    //std::vector<bool> hasApple{ false,false,true,false,true,true,false };
    std::vector<bool> hasApple{ true,true,true,true };
    std::cout << minTime(hasApple.size(), edges, hasApple);
}

