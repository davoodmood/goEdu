
# Binary Search Tree

The top node is called the root. The nodes that are directly under a nodes are called children, while the node above them is known as _Parent_. 
The nodes at the bottom of the tree without any children are known as _leaves_. A tree node with its parents having no more than two children each, is called a _Binary Tree_.
Therefor, each parent has a left and a right child. The smaller child is placed on the left and the larger on the right, we call it a Binary Search Tree.

## Advantages

### Speed
The advantage of using binary tree is literally its speed. In a somewhat well balanced binary search tree, the `Big O` is being  `O(h)`. i.e. way better than `O(n)`.


## Insert and Search

in both cases we want to start from the root and compare a given value with root node. if the value is smaller than the node, we will move it down to the left otherwise to the right. we follow this process untill we can land the value on leaf.