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
Considering the same example of the _adjacency list_, if we have five `vertices` in the graph, we will need a `5 x 5` matrix.

<table>
  <thead>
    <tr>
      <th rowspan="2" colspan="2"></th>
      <th rowspan="2">vertex</th>
      <th colspan="5">Columns</th>
    </tr>
    <tr>
      <th colspan="5">to</th>
    </tr>
    <tr>
      <th rowspan="6" colspan="2">vertex</th>
      <th>Keys</th>
      <th>1</th>
      <th>2</th>
      <th>3</th>
      <th>4</th>
      <th>5</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td rowspan="0"><strong>Rows</strong></td>
      <td rowspan="0"><strong>from</strong></td>
      <td><strong>1</strong></td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
    </tr>
    <tr>
      <td><strong>2</strong></td>
      <td>1</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
    </tr>
    <tr>
      <td><strong>3</strong></td>
      <td>1</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
      <td>1</td>
    </tr>
    <tr>
      <td><strong>4</strong></td>
      <td>0</td>
      <td>1</td>
      <td>0</td>
      <td>0</td>
      <td>1</td>
    </tr>
    <tr>
      <td><strong>5</strong></td>
      <td>1</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
      <td>0</td>
    </tr>
  </tbody>
</table>

If the edges in the graph had a weight, the matrix values instead of being binary (`0` or `1`), would be they weight _integer_ value. e.g. `2`, `5`, `8` or `10`.
In case of the undirected graphs, the matrix will be uncemetrical.

#### Adding edges
when we _add_ a new edge to the graph, we need to add a new integer to the `from` & `to` vertecies. 
#### Adding verteces
when we _add_ a new vertex, we must add a new _row_ & _column_ to the matrix. this means you have to make a new matrix and replace it with the current one.

## List vs Matrix Implementation Comparison
We can compare the two implementation approaces in the follwing:

### by Storage Space
if we have a sparse Graph, having an `Adjacency List` helps with saving more storage space. But an adcencancy matrix, we always need _v ** 2_ space, where _v_ is the number of the `verteces`.

### by Edge Lookup
If you want to check whether two vertices are connected or which individuals are following each other. In such case, the `Adjacency Matrix` is way faster for the lookup. In face it would take constant time of `O(1)`. In contrary a lookup action in  `Adjacency List` a traversing of the adjacent vertices is required, that would in worse case take `n` steps if the number of vertices is `n`.

## time complexity
|Topic|Adjacency Matrix|Adjacency List|
|-------|--------|--------|
|Edge Lookup|`O(n ** 2)`| `O(n + E)` |
|Edge Lookup|`O(1)`| `O(n)` |
|Adding Vertex|`O(n ** 2)`| `O(1)` |
|Removing Vertex|`O(n ** 2)`| `O(E)`* |
|Adding Edges|`O(1)`| `O(1)` |
|Removing Edges|`O(1)`| `O(n)`* |


`*`: `E = number of edges`

or in terms of performance:

|Topic|Adjacency Matrix|Adjacency List|
|-------|--------|--------|
|Edge Lookup|`Horrible`| `OK` |
|Edge Lookup|`Great`| `OK` |
|Adding Vertex|`Horrible`| `Great` |
|Removing Vertex|`Horrible`| `OK` |
|Adding Edges|`Great`| `Great` |
|Removing Edges|`Great`| `OK`* |
