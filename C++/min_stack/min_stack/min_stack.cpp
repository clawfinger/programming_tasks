#include <iostream>
#include <stack>

class MinStack {
public:
    MinStack() {

    }

    void push(int val) {
        main.push(val);
        if (min.empty() || val <= min.top()) {
            min.push(val);
        }
    }

    void pop() {
        if (min.top() == main.top()) {
            min.pop();
        } 
        main.pop();
    }

    int top() {
        return main.top();
    }

    int getMin() {
        return min.top();
    }
private:
    std::stack<int> main;
    std::stack<int> min;
};

int main()
{
    MinStack* st = new MinStack();
    st->push(0);
    st->push(1);
    st->push(0);
    std::cout << st->getMin() << std::endl;
    st->pop();
    std::cout << st->getMin() << std::endl;
}

