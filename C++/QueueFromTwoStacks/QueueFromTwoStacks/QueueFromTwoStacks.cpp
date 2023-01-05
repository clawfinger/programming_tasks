#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
#include <stack>
using namespace std;

class StackQueue {
public:
    void enqueue(int entry);
    void dequeue();
    int top();
private:
    stack<int> m_push;
    stack<int> m_pop;
};


void StackQueue::enqueue(int entry)
{
    m_push.push(entry);
}

void StackQueue::dequeue()
{
    if (!m_pop.empty()) {
        m_pop.pop();
        return;
    }

    while (!m_push.empty()) {
        int value = m_push.top();
        m_push.pop();
        m_pop.push(value);
    }
    m_pop.pop();
}

int StackQueue::top()
{
    if (m_pop.empty()) {
        while (!m_push.empty()) {
            int value = m_push.top();
            m_push.pop();
            m_pop.push(value);
        }
    }
    return m_pop.top();
}

int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */     int queriesNum = 0;
    StackQueue queue;
    cin >> queriesNum;
    int type;
    for (int i = 0; i < queriesNum; i++) {
        cin >> type;
        switch (type) {
        case 1:
            int entry;
            cin >> entry;
            queue.enqueue(entry);
            break;
        case 2:
            queue.dequeue();
            break;
        case 3:
            cout << queue.top() << endl;
            break;
        default:
            return 9;
        }
    }
    return 0;
}
