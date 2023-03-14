#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
 };
 
bool hasCycle(ListNode* head) {
    if (head == nullptr) {
        return false;
    }
    ListNode* slow = head;
    ListNode* fast = head;

    while (slow != nullptr && fast != nullptr && fast->next != nullptr) {
        slow = slow->next;
        fast = fast->next->next;
        if (slow == fast) {
            return true;
        }
    }

    return false;
}

int main()
{
    ListNode* one = new ListNode(1);
    ListNode* two = new ListNode(2);
    ListNode* three = new ListNode(3);
    ListNode* four = new ListNode(4);
    one->next = two;
    two->next = three;
    three->next = four;
    //four->next = two;
    std::cout << hasCycle(one);
}
