# Linked Lists

In `linked lists` data is stored in a sequence of `nodes`. compared to an array where values are stored in each index of an array, 


linked lists can be recognozed in 2 types:

1. **Singly Linked List:** In this type, each `node` will link only to its next node.
2. **Doubly Linked List:** In this type, each `node` addtional to linking to its next node, will store the memory position of its previous node.


When it comes to link lists, we first go the HEAD then we follow the links i.e. if we want the 1000th node we should sequentially move from head to that node. The question is why would we be the use case for this as such application sees to be flawd. 

## Used Case

In fact, some applications require this data structure because it is extremely faster for __ADDING__ or __DELETING__ elemnts to/from the head/end node. It is much faster than arrays/slices in its time complexity of `O(1)`. this is because when adding or removing and item to the `0` index of an array, other items must first move an index forward or backward.
