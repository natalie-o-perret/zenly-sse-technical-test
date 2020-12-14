# Senior Software Engineer Technical Test

## Summary

Today zenly's social graph is made of 30 million users with approximately 5 million daily active users. 

Our goal is to reach in the next two years a 100 million daily active users; that exercise aims at testing your capabilities to design a system which would handle that load.

The social graph must be able to store relationships between users:
- Your account (let's say your phone number)
- Your own address book
- Your contacts address books

Common queries asked on that social graph:
- Who are your contacts?
- Who has you in their contacts?
- Which people that are not in your contacts that you might know?

Let's rephrase this a bit more mathematically.

### Contact Graph

- We need to implement a contact graph composed of phone numbers and their relations.
- We need this data structure to answer the following queries for a dataset of size 100M nodes with 50 edges on average per node.

Normalized international phone numbers are formatted like `+1 123 456 789 00`

### Example Graph

```
x -> y
x -> z
y -> x
y -> z
p -> x
```

### Requirements

#### Lookup
For a given phone number `x`, get all the contacts of `x`:

```txt
lookup(x) = [y, z]
```

#### Reverse Lookup

For a given phone number `x`, get all the phone numbers who have `x` in their contacts:

```txt
rlookup(z) = [x, y]
```

#### Suggestions

For a given contact `x`, return the ordered list of top `10` suggested contacts using a consistent set of rules of your choosing:

```txt
suggest(p) = [y, z]
```

## What we're expecting

We need an in-memory implementation of a data structure that answers the 3 queries specified above on a single machine of reasonable size (commodity hardware).

Any optimization possible given the time dedicated to this test is appreciated. Ideally we want an algorithm that answers these queries as fast as possible.

You can generate random data to test the queries but the data should satisfy these constraints:
- 100M nodes in graph
- An average of 50 contacts per phone number
- A random distribution of contacts for each number

The solution can be coded in the language / stack of your choosing but given the memory requirements of the data structure, a system's programming language is recommended. 

We propose that you take a few hours to think about the problem and spend something like 6 hours. 

We appreciate that building a perfect solution in such a short amount of time is not possible but the result should be enough to let us discuss the ideal solution.

Feel free to ask additional questions, you can reach us at: [olivier@zen.ly](olivier@zen.ly) / [mehran@zen.ly](olivier@zen.ly)
