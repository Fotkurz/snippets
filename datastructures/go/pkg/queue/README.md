# Queue

Queue data structures are characterized as FIFO or First In First Out.

This means that the first elements added to the queue are the first to
be removed from the queue.

Queues are usually used today in event drive architectures to asynchronous process messages.

Some of te core methods of queues are:

- IsEmpty() -> Checks if queue is empty
- Front() -> Returns the first element of the queue (the next to be processed).
- Back() -> Returns the last element of the queue.
- Dequeue() -> Remove the las telement of the queue.
- Enqueue() -> Add a new element to the queue.

Some implementations of the Queue with DoublyLinked Lists, can achieve
Dequeue and Enqueue time complexity of O(1). (We should try this here).
