#include <stdlib.h>
// 引入 stdbool 以使用 bool 类型
#include <stdbool.h>

// --- Linked List ---
typedef List Node;
typedef struct List {
    int val;
    Node next;
} List;

void listPushFront(Node* head, int x) {
    Node tmp = malloc(sizeof(*tmp));
    tmp->val = x;
    tmp->next = head->next;
    head->next = tmp;
}

void listDelete(Node* head, int x) {
    for (Node t = head->next; t; t = t->next) {
        if (t->val == x) {
            Node tmp = t->next;
            t->next = tmp->next;
            free(tmp);
            // 因为 HashSet 中不能有相同的元素
            // 所以这里直接退出
            break;
        }
    }
}

bool listContains(Node* head, int x) {
    for (Node t = head->next; t; t = t->next) {
        if (t->val == x) {
            return true;
        }
    }
    return false;
}

void listFree(Node* head) {
    for (; head->next;) {
        Node tmp = head->next;
        head->next = tmp->next;
        free(tmp);
    }
}

// --- HashSet ---
const int base = 769;

typedef struct {
    Node* data;
} MyHashSet;

int hash(int key) {
    return key % base;
}

MyHashSet* myHashSetCreate() {
    MyHashSet* set = malloc(sizeof(*set));
    // Array of Node
    set->data = malloc(sizeof(Node) * base);
    for (int i = 0; i < base; i++) {
        set->data[i].val = 0;
        set->data[i].next = NULL;
    }
    return set;
}

bool myHashSetContains(MyHashSet* s, int key) {
    int h = hash(key);
    return listContains(&(s->data[h]), key);
}

void myHashAdd(MyHashSet* s, int key) {
    if (!myHashSetContains(s, key)) {
        int h = hash(key);
        listPushFront(&(s->data[h]), key);
    }
}

void myHashSetRemove(MyHashSet* s, int key) {
    int h = hash(key);
    listDelete(&(s->data[h]), key);
}

void myHashSetFree(MyHashSet* s) {
    for (int i = 0; i < base; i++) {
        listFree(&(s->data[i]));
    }
    free(s->data);
}