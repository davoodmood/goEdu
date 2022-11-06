# Hash Tables
`Hash Tables` is a data structure that is uniqley fast for inserting, deleting & searching data. Knowing about it can help in implementation of a its more complex forms namely `HashMap`/`HashSet`.

To understand hash tables better, let's first imagine an `array`/`slice` with a bunch of content stored in it. To search your intended value in that array, you must check array items slot by slot from index 0 to n. This can be a super simple, but in reality arrays can become much larger than using brute-force to find values in an array. As you may guess its pretty inefficient. 

Now, let's say there is a way to find out which index our intended value is saved in. We can understand the index by the content of the value we are searching. This is the idea of a hashtable.

The thing that can turn our intended value into the right array index is called a `Hash Function`. The value is the input of the `Hash Function` and the output would be the index of that value in a hash table. Hash function's output is also known as a `Hash Code`, which will be location of our input value. This would be possible, because when we are storing data we are putting the input value in a hash function which will give us the array index. there for we store it in that indexed location, and later when we are trying to search for the same value, we can recover it by looking at the same hashcode. a data structure with this mechanism is called a `Hash Table`.


## A Hash Function
Let's look at a simple Hashing algorithm, let's say our input value is the string , "Dave". Each letter of this word "Dave" will be converted to its `ascii` code i.e. truning each letter to a number. Then we sum them up and devide them by the size of our array, _e.g. 100_, and we get the remainder. This remainder will be a number between `0` and `Array Size - 1` which in our exampe is `99` and this output is our `hash code`

This means for every input value, we calculate the `hash code` by running it through the `Hash Function` and look into its output index to find/store the value. 

## Collision Handling
But what if two different values will have the same indexin a `Hash Table`? Honestly, it can be a common matter to have such a collision. 
there are two suggested ways to handle this.

- Open Addressing
- Seperate Chaining

### Open Addressing
If the spot for the second value is already filled in the hash table, then we store it in the next index. i.e. we will first look at the outputed `hash code` and if the slot is not empty, but the value isn't there, then we move to the next index. and we follow this process until we either find the value or an empty slot. even though we will seach slot by slot in this case, it would still be faster than searching all the way from index 0.

#### Use Case
When there is a limited number of inputs, which its total is predictable. 
#### Con
As the hashtable gets more populated, the values will go farther away from their intended location, and we start to lose the benefits of the hashtable.


### seperate chaining
A simple way to explain Seperate Chains, is to sotre mulitple values in one index. This will be possible with a linked list data-structure. instead of saving the input in the index, we save it in a `linked list` and save a pointer to the `linked list`'s `head` in that index.

`linked list` will be able to grow or shrink dynamically. Therfore it would have no cap over the size of data we can store. In face, this combination of an `hashtable` and `Linked list` allows **seperate chaining** a very fast, effective and efficient.

## Hash Table Methods
We will vall the value to be stored as `keys` and the linked-lists in the array index can be reffered to as `buckets`. we call the name keys because hash tables are used to store key-value pairs.

### Insert
We first get the `hash code` by putting the `key` in the `hash function` e.g. `hash("key")`. Then we go to the output index. In order to place our `key` in the linked-list we first need to search for that key in the `bucket` to know if it exists or not. If it doesn't already exist, we will add the node represented by the key into the linked-list.
### Delete
In order to delete a key, we will find it's allocated bucket by the retrieved `hash code`, look into the linked-list and we would disconnect it from the bucket if found. We disconnect it by letting the previous node, skip the `key`'s node and make it point to the next node.
### Search
When searching for the `key` we first run it through the `hash function` to find its allocated bucket. Then we go to the bucket and search it's nodes one by one to find the key.

## Time Complexity

The time complexity of Hash tables is expected to be constant in an average case.

|HashTables          |
|--------------------|
|Average Case        |
|--------------------|
|INSERT|SEARCH|DELETE|
|------|------|------|
|O(1)  |O(1)  |O(1)  |

<table>
  <thead>
    <tr>
      <th colspan="3">HashTables</th>
    </tr>
    <tr>
      <th colspan="3">Average Case</th>
    </tr>
    <tr>
      <td>INSERT</td>
      <td>SEARCH</td>
      <td>DELETE</td>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>O(1)</td>
      <td>O(1)</td>
      <td>O(1)</td>
    </tr>
  </tbody>
</table>

### The time complexity of constant O(1) by comparison of data structures:

|Data Structure|`INSERT`|`SEARCH`|`DELETE`|
|--------------|--------|--------|--------|
|Array|❌|❌|❌|
|LinkedList|✔️|❌|✔️|
|Hash Table|✔️|✔️|✔️|