# Data Structures - Singly LinkedList

Linked-list are one of the fundamentals data structures for computer science.

Linked-lists are like arrays of 2 elements being the first element, the data and the second element, a pointer to the location of the next element in the list.

This way the linked lists can be stored anywhere in the memory while keeping the track
of the elements.

![singly-linked-list](/assets/linkedlist.drawio.png)

This represents the Singly Linked-list since there is only a Next pointer.

There is also the Doubly Linked-list where each **Node** has the reference of the previous and next element.

![doubly-linkedlist](/assets/doublylinkedlist.drawio.png)

### Advantages

- Linked-lists are dynamic, so they can grow and shrink at runtime.
- No wastes of memory since you are not using more than you require.
- Insertion and deletion operations are quite easier to run and implement, since there is no need to shift elements position like in arrays.

### Disadvantages

- Uses more memory than arrays, since each element, in a single linked list, is made of 2 nodes.
- Traversal is harder in linked-list than in arrays and direct access to an element is not possible.
- Search and iteration operations are more expensive than in arrays.