#include <iostream>


struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    ListNode* head = nullptr;
    ListNode* res = nullptr;

    int overflow = 0;
    while (true) {
        if (l1 == nullptr && l2 == nullptr) {
            break;
        }
        int current = 0;
        if (l1 != nullptr) {
            current += l1->val;
            l1 = l1->next;
        }
        if (l2 != nullptr) {
            current += l2->val;
            l2 = l2->next;
        }
        current += overflow;
        int meaning = current % 10;
        overflow = current / 10;
        ListNode* node = new ListNode(meaning);
        if (head == nullptr) {
            head = node;
            res = head;
        }
        else {
            head->next = node;
            head = node;
        }
    }
    if (overflow != 0) {
        head->next = new ListNode(overflow);
    }
    return res;
}

int main()
{
    ListNode* one = new ListNode(9);
    ListNode* two = new ListNode(9);
    ListNode* three = new ListNode(9);

    one->next = two;
    two->next = three;

    ListNode* four = new ListNode(1);
    //ListNode* five = new ListNode(6);
    //ListNode* six = new ListNode(4);

    //four->next = five;
    //five->next = six;
    ListNode* res = addTwoNumbers(one, four);
}
