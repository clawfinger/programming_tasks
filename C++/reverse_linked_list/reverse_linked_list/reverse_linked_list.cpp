#include <iostream>


struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};


ListNode* reverseList_rec(ListNode** newHead, ListNode* current) {
    if (current->next == nullptr) {
        *newHead = current;
        return current;
    }
    ListNode* res = reverseList_rec(newHead, current->next);
    res->next = current;
    return current;
}

ListNode* reverseList(ListNode* head) {
    ListNode* newHead;
    ListNode* oldHead = reverseList_rec(&newHead, head);
    oldHead->next = nullptr;
    return newHead;
}

ListNode* reverseListIterative(ListNode* head) {
    if (head == nullptr || head->next == nullptr) {
        return head;
    }
    ListNode* prev = nullptr;
    ListNode* current = head;
    while (current != nullptr) {
        ListNode* next = current->next;
        current->next = prev;
        prev = current;
        current = next;
    }
    return prev;
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

    auto res = reverseListIterative(one);
}
