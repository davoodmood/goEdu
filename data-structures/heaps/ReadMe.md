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

### Finding child nodes

Once implmented, child nodes can be determined by simply running the parent node index through the folowing formula:

| Parent index X  | Operators        | Child Index Position |
|-----------------|------------------|----------------------|
| index           | x 2 + 1 =        | left Position Index  |
| index           | x 2 + 2 =        | right Position Index |

### Inserting Keys 

Whenever we insert a key, we add it to the bottom right node of the tree, i.e. the last index of the Slice (Array). Then, we need to rearrenge the nodes, so we can maintain the heap proority. That is to keep the parent key larger than its children. To do this, we compare the last node to its parrent node and swap it if the value is larger than its parent. We follow up the tree, and keep repeating this process until it gets to its place. This process of rearrenging the indexes is called `Heapify`. 

The heapifying process also occures when we want to take an item out of the tree.

### Extracting the Highest Key

To extract the proirity key, we simply take out the root element of the Heap, which is the largest key. After that we need to heapify (rearrenge the nodes), so the empty root can be filled, while the heap property is maintained. To do this, we will initially fill the root with the last node (last index). Then we start heapifying from the top. In perespective, We will compare the children of the root element for the bigger value, then compare the biggest child with the parent node. and repeat this process until the Heap is in its propper state.

### Time complexity

The time complexity of extracting and inserting in the heap will depend on the heapifying process. While the actions of Insertion and extraction are not the dominant part of the complexity, the hight of the tree is the factor to determine the Time Complexity of the operation. Therefore, if we recognize the hight of the tree as `h` the time complexity would be `O(h)`. Now, if we consider `n` as the number of nodes in the tree, its realtion to the hight can be consider `log n` as it has logarithmic relation to the nodes. 

> As a result, we can consider the time complexity if Heap Insert and Extract operations as `O(log n)`


## Golang Implementation

To build the structure, we are going to follow this pattern bellow:

1. We start by creating a struct for the Max Heap
2. Creating an Insert Method
3. Writing the Extract Method

We will also build the following methods to help us handle the logic:

* `maxHeapifyUp`
* `maxHeapifyDown`
* `swap`

And the following functions:

* `parent`
* `left`
* `right`
