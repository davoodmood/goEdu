# Graphs

Graphs are like tree structures, meaning they have `nodes` and `edges`. But unlik `trees` which the nodes are pointing from parent to children, in **Graphs** nodes can point to other nodes including thier parent, to itself or nodes in differnet `branches` and even make `cycles`. That means this new data structure can not be a tree anymore, and it's now become a `Graph` data structure.

## What is a graph?
In the graph we use similar concepts in trees and call them differently.

|Data Structure|Slots|Connectors|
|`Trees`|Nodes|Pointers|
|`Graph`|Vertex|Edges|

When a **Graph** has many `edges` we say it is **`dense`** and if it has less edges, we recognize it as **`sparse`**.

Graphs can be `Directed` which means that they have a directional edges between nodes or `Undirected`.
Graphs can be `Cyclic`, or you can place weights on the edges which will make it a `Weighted Graph`.

by this definition, we can also see a normal tree is also a type of graph, but it has more limitations e.g. it cannot connect child nodes to thier parent, or each node can only have 2 direct children (in a binary tree)

There are a lot of diverse implementations of a graph as it can be calibrated into different needs.
- edges might need to be weighted
- edges might not need directions
- graph may allow cycles
- building a networ system computer
- adding edges very often/rarely
- graph might not need to care about performance

## Graph in code

The 2 common ways to represent a graphs are:
1. Adjecency List: a graph is expressed a list of verteces
2. Adjacency Matrix

### Adjecency List
In this implementation, a graph is expressed a list of verteces, while each vertex will have a list of its neigboring `vertecies`. Essentially, It's looking like a list of lists. these lists can be any other type e.g. a `linked-list`, `arrays`, or `dynamic slices`, i.e anything that can implement a list.

Imagine _vertex_ with _key_ of `1` points to two other vertices, `2` and `4`. and `2` is pointing to `1` (meaning they are `bi-directional`). note that if we work with `undirected graph` all the nodes are considered `bi-directional` by default, contrary that is not the case in `directed graphs`.

we call the `integer` values of the list `Keys` and if they are going to represent a set of information each key must be _unique_.
#### Adding edges
when we _add_ a new edge to the graph, we can add a new element to the adgecency list. 
#### Adding verteces
when we _add_ a new vertex to the graph, we will add a vertex to the graph list. 

#### A Table view example
|Vertex|Keys|
|------|----|
|`1`|`2`,`4`|
|`2`|`1`|
|`3`|`1`,`4`,`5`|
|`4`|`2`,`5`|
|`5`|`1`|
|`6`||

### Adjecency Matrix
In this implementation, a graph is expressed by a 2-dimentional (2D) array. The `vertecies` of the graph as `columns` & `rows` which will respectively represent the `from` and `to` vertex.

