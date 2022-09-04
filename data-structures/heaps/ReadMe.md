# Heaps in Golang

## What is a heap?

A heap is a data-structure that is considered as a tree and is built by using a `Slice` (an Array). It was originally introduced as part of an algorithm called `Heap-sort`. The data-structure that was used in `heap-sort` was also useful for other applications, such as:

*   Priority Queues `(abstract data type)`
*   Selection Algorithms
*   Graph Algorithms

In this section we will go deeper into priority Queues.

### Priority queues Vs. Normal queues

A normal queue is based on a "first in first out" approach, but in the Priority Queues "the highest priority is being served first". This is where the `Heaps` come in, it is a data-structure that can be expressed as a complete tree. A complete tree means that all the levels in the tree are full, except for the lowest level. Where some nodes can be empty but the existings are all on the left. 


## Heap models

`Heaps` can are commonly used in the form of `max heap` and `min heap`. 

In the case of `max heap`, largest key will be stored in the root, and for every node in the tree, every parent will contain a higher key than its children. This is extremely important, because in the case of Priority Queues we are looking to operate or take the key with the largest priorty value, which will always be in the root. therefore, getting the largest key will be so fast, regardless of the heap size.

On the opposit side, is the `min heap` where the root is the smallest key. therefore in this model, each parent will have the smallest key than its children.

In this code, we are focusing on `max heaps` as the `min heaps` are simply the reverse of this implementation.

## Max Heap Implementation

As mentioned above, while heaps are visualized as a tree structure. Behind the scene, they are simply Slices (Arrays) whose `index 0` act as the `root` and the following indexes depict the children of each parent node from left to right.

This is possible through the folowing formula:

### left child node

| Parent index X  | Operators        | Child Index Position |
|-----------------|------------------|----------------------|
| index           | x 2 + 1 =        | left Position Index  |
| index           | x 2 + 2 =        | right Position Index |