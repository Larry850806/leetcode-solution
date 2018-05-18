/*
Detect a cycle in a linked list. Note that the head pointer may be 'NULL' if the
list is empty.

A Node is defined as:
    struct Node {
        int data;
        struct Node* next;
    }
*/

// struct Node {
//   int data;
//   struct Node *next;
// };

bool has_cycle(Node *head) {
  Node *fast = head, *slow = head;
  if (head == 0) {
    return false;
  }
  while (true) {
    fast = fast->next;
    if (fast == 0) {
      return false;
    }
    fast = fast->next;
    if (fast == 0) {
      return false;
    }
    slow = slow->next;
    if (fast == slow) {
      return true;
    }
  }
  // Complete this function
  // Do not write the main method
  return true;
}

// int main() {}