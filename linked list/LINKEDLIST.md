# Linked list Data Structure

A linked list is a linear data structure that includes a series of connected nodes. Here, each node stores the **data** and the **address** of the next node. For example,

![linked list data structure](https://cdn.programiz.com/sites/tutorial2program/files/linked-list-concept.png "Linked list Data Structure")

Linked list Data Structure

You have to start somewhere, so we give the address of the first node a special name called HEAD. Also, the last node in the linked list can be identified because its next portion points to NULL.

Linked lists can be of multiple types: **singly**, **doubly**, and **circular linked list**. In this article, we will focus on the **singly linked list**. To learn about other types, visit [Types of Linked List](https://www.programiz.com/dsa/linked-list-types "Types of Linked List").

**Note:** You might have played the game Treasure Hunt, where each clue includes the information about the next clue. That is how the linked list operates.

---

## Representation of Linked List

Let's see how each node of the linked list is represented. Each node consists:

- A data item
- An address of another node

We wrap both the data item and the next node reference in a struct as:

```C
struct node
{
  int data;
  struct node *next;
};
```

Understanding the structure of a linked list node is the key to having a grasp on it.

Each struct node has a data item and a pointer to another struct node. Let us create a simple Linked List with three items to understand how this works.
```C
/* Initialize nodes */
struct node *head;
struct node *one = NULL;
struct node *two = NULL;
struct node *three = NULL;

/* Allocate memory */
one = malloc(sizeof(struct node));
two = malloc(sizeof(struct node));
three = malloc(sizeof(struct node));

/* Assign data values */
one->data = 1;
two->data = 2;
three->data=3;

/* Connect nodes */
one->next = two;
two->next = three;
three->next = NULL;

/* Save address of first node in head */
head = one;
```

In just a few steps, we have created a simple linked list with three nodes.

![representing linked list by connecting each node with next node using address of next node](https://cdn.programiz.com/sites/tutorial2program/files/linked-list-with-data.png "Linked list Representation")

Linked list Representation

The power of a linked list comes from the ability to break the chain and rejoin it. E.g. if you wanted to put an element 4 between 1 and 2, the steps would be:

- Create a new struct node and allocate memory to it.
- Add its data value as 4
- Point its next pointer to the struct node containing 2 as the data value
- Change the next pointer of "1" to the node we just created.

Doing something similar in an array would have required shifting the positions of all the subsequent elements.

In python and Java, the linked list can be implemented using classes as shown in [the codes below](https://www.programiz.com/dsa/linked-list#code).

---

## Linked List Utility

Lists are one of the most popular and efficient data structures, with implementation in every programming language like C, C++, Python, Java, and C#.

Apart from that, linked lists are a great way to learn how pointers work. By practicing how to manipulate linked lists, you can prepare yourself to learn more advanced data structures like graphs and trees.

---

## Linked List Implementations in C

```C
#include <stdio.h>
#include <stdlib.h>

// Creating a node
struct node {
  int value;
  struct node *next;
};

// print the linked list value
void printLinkedlist(struct node *p) {
  while (p != NULL) {
    printf("%d ", p->value);
    p = p->next;
  }
}

int main() {
  // Initialize nodes
  struct node *head;
  struct node *one = NULL;
  struct node *two = NULL;
  struct node *three = NULL;

  // Allocate memory
  one = malloc(sizeof(struct node));
  two = malloc(sizeof(struct node));
  three = malloc(sizeof(struct node));

  // Assign value values
  one->value = 1;
  two->value = 2;
  three->value = 3;

  // Connect nodes
  one->next = two;
  two->next = three;
  three->next = NULL;

  // printing node-value
  head = one;
  printLinkedlist(head);
}
```

---

## Linked List Complexity

Time Complexity

||Worst case|Average Case|
|---|---|---|
|**Search**|O(n)|O(n)|
|**Insert**|O(1)|O(1)|
|**Deletion**|O(1)|O(1)|

Space Complexity: `O(n)`
---

## Linked List Applications

- Dynamic memory allocation
- Implemented in stack and queue
- In **undo** functionality of softwares
- Hash tables, Graphs

---

### Recommended Readings

### 1. Tutorials

- [Linked List Operations (Traverse, Insert, Delete)](https://www.programiz.com/dsa/linked-list-operations "Linked List Operations")
- [Types of Linked List](https://www.programiz.com/dsa/linked-list-types "Types of Linked List")
- [Java LinkedList](https://www.programiz.com/java-programming/linkedlist "Java Linked List")

### 2. Examples

- [Get the middle element of Linked List in a single iteration](https://www.programiz.com/java-programming/examples/get-middle-element-of-linkedlist "Get middle element of Linked List")
- [Convert the Linked List into an Array and vice versa](https://www.programiz.com/java-programming/examples/linkedlist-array-conversion "Convert Linked List into Array")
- [Detect loop in a Linked List](https://www.programiz.com/java-programming/examples/detect-loop-in-linkedlist "Detect Loop in a Linked List")