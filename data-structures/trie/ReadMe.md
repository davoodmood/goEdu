# Trie data structure

a Trie is a tree structure which is mostly used for storing words. Each `node` will represent a word or part of a word, and by connecting the path from the root to that node, we can store words in this particular structure.

## Use cases
as it allows, for fast seaching of words, the Trie data structure can be used for `auto-complete` features. similar to how google guesses your search query. 

## Structure
tries will have a root node at the top, which will have its child nodes as a list of array, where each index will represent all the possible `chars` in the alphabet. In fact, each node can have a child per each character in the alphabet. The ones that represent the children will have a `pointer` inside it, that points to that child node.

for example:
A node `D` has two children `o` and `g`. this means that this `D` node, has an array with an index form `0` to `25`, that respectively represent `a` to `z`.  this array at its `o` & `g` index equevelants stores a `pointer` to the addresses of the next nodes, while the other indexes are `null`/`nil`.

each node also has a `boolean` property value that indicates whether they are the ending letter of a word.

## Methods

# INSERT

