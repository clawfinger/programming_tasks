// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/

#include <iostream>
#include <vector>
#include <map>
#include <set>
#include <queue>
#include <stack>

std::vector<int> countSubTreesSlow(int n, std::vector<std::vector<int>>& edges, std::string labels) {
    std::vector<std::vector<int>> connections(n);
    for (int i = 0; i < edges.size(); i++) {
        connections[edges[i][1]].push_back(edges[i][0]);
        connections[edges[i][0]].push_back(edges[i][1]);
    }
    std::vector<bool> visited(n, false);
    std::queue<int> queue;
    std::vector<int> parents(n);

    queue.push(0);
    visited[0] = true;

    parents[0] = -1;
    std::vector<int> res(n, 1);
    while (!queue.empty()) {
        int current = queue.front();
        queue.pop();
        const auto& adjascent = connections[current];

        for (int i = 0; i < adjascent.size(); i++) {
            int adj = adjascent[i];
            if (!visited[adj]) {
                parents[adj] = current;
                queue.push(adj);
                visited[adj] = true;
            }
        }
    }


    for (int i = 0; i < n; i++) {
        char c = labels[i];
        if (c == labels[i]) {
            res[i]++;
        }
        int parent = parents[i];
        while (parent != -1) {
            if (labels[parent] == c) {
                res[parent]++;
            }
            int newParent = parents[parent];
            parent = newParent;
        }
    }
    return res;
}

int main()
{
    //std::vector<std::vector<int>> edges{ {0,1 }, { 0,2 }, { 1,4 }, { 1,5 }, { 2,3 },{2, 6}};
    std::vector<std::vector<int>> edges{ {0,1 }, { 1,2 }, { 0,3 }};
    std::string labels = "bbbb";
    //std::string labels = "abaedcd";
    auto res = countSubTreesSlow(labels.size(), edges, labels);
}

