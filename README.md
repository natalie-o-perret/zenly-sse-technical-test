# Zenly: Senior Software Engineer - Technical Test

Note: Make sure your current working directory is `.\src`.

## How to run the mem benchmark

```bash
go run main.go
```

## How to run the (cheap) unit tests and (even cheaper) speed benchmarks

```bash
go test .\lib -bench=. -v
```


## Why GO btw?

Since Zenly is atm primarily relying on the Go programming language,
I thought it will be good for me if I could give it another try.

It helped to bump into a few typical go-ish hiccups (eg. [append](https://yourbasic.org/golang/gotcha-append/)).

Also investigated about [how the GC works](https://blog.golang.org/ismmkeynote)
As well as performances when it comes to:
- [Heap and stack in Go](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services)
- [Go: Memory Management and Allocation](https://medium.com/a-journey-with-go/go-memory-management-and-allocation-a7396d430f44)
- [Avoiding high GC overhead with large heaps](https://blog.gopheracademy.com/advent-2018/avoid-gc-overhead-large-heaps)
- [Bitset implementations in Go](https://medium.com/@val_deleplace/7-ways-to-implement-a-bit-set-in-go-91650229b386)

Notes:

- I could probably have better scored in terms of speed and memory with a language not relying on a GC like using Rust, C++ or even C.
- Actually, looking back if I had more time I think I would love to do this code assignment with [nim](https://nim-lang.org), cause it can actually transpile among other langs, to C ~~.

##  Solution Rationale

### Storage Design

- The `Node` `N` structure represents a contact `C` and holds
    - An array `to`, ie. nodes / contacts `CT` of `C`.
      These nodes represent the 1st-degree connections.
    - Conversely, another array `from`, ie. nodes / contacts `CF` who have `C` in their phone book.
      Also 1st-degree, but the other way around.

- The `Graph` `G`, holding the key-value pairs is just a hashmap mapping a phone number `P` of a contact `C` to its node `N`.

If we take the example given in the instructions:
```
x -> y
x -> z
y -> x
y -> z
p -> x
```

The storage then looks like:

| Phone Number | `Node.from` | `Node.to` |
|--------------|-------------|-----------|
| `x`          | `[y, p]`    | `[y, z]`  |
| `y`          | `[x]`       | `[x, z]`  |
| `z`          | `[x, y]`    | `[]`      |
| `p`          | `[]`        | `[x]`     |

### Algorithm

#### `lookup`

##### Objective

Returns all the contacts for the given phone number.

##### Example

For the given example above:
```go
lookup(x) = [y, z]
```

##### How to

1. Find the relevant contact node `N` using the phone number as a key in the hashmap
2. We return the content of the `to` property in `N`
3. Preferably, we can return a copy (ie. immutability)
   if the array is a direct reference to avoid the client of the library
   to mess with the internal storage (also depends on the language and collection).

##### Complexity

`~O(1)`: constant

To be really accurate, depends on whether we want to make copy or not of the phone numbers
(and sort them in a certain order).

#### `rlookup`

##### Objective

Returns all the contacts that have the given phone number.

##### How to

Returns all the contacts for the given phone number.

1. Find the relevant contact node `N` using the phone number as a key in the hashmap
2. We return the content of the `from` property in `N`
3. Preferably, we can return a copy (ie. immutability)
   if the array is a direct reference to avoid the client of the library
   to mess with the internal storage (also depends on the language and collection).

##### Example

For the given example above:
```go
rlookup(z) = [x, y]
```

##### Complexity

`~O(1)`: constant.

To be really accurate, depends on whether we want to make copy or not of the phone numbers
(and sort them in a certain order).

#### `suggest`

Returns up to 10 2nd degree contacts given an initial contact phone number.

Put it otherwise:

> The friends of my friends, who aren't already my friends and who are not... myself either.

##### How to

Simpler version of trad. graph node visiting stuff:
1. Find the relevant contact node `N` using the phone number as a key in the hashmap
2. Visit the node of 1st degree (aka the phone numbers of `N`)
3. Yield only the unique phone numbers

##### Example

For the given example (way) above:
```go
suggest(p) = [y, z]
```

##### Complexity

`~O(E1 . E2)`: which is the average complexity, depends on...
- the number 1st degree phone numbers `E1` of the initial given contact has
- multiplied by the average number `E2` of contacts that first contacts have themselves

### Implementation Details

- No generic or empty interface trickery, not much to my dismay tho, plus it's Go.
- For the grand sake of saving up some cpu cycles, I haven't used sorting (ie. afaik `O(n.log(n))`) on arrays to guarantee the order,
  tbs, there is a check to test whether collection are equivalent in the unit tests.
- If Go wasn't too limited I could have created an actual hashset with the key of the hashmap acting as the node
  (by specifying which fields are relevant to the equality and hash) and using the good ol' trick of the value impl. 
  as an empty `struct{}` to save some memory.
  Instead of repassing thru the hashmap for every single key, which despite `O(1)` and being less memory intensive, 
  takes more time. But the equality system for some particular keys is a tad twisted,
  (collections and pointers most notably).

## Discussions 

Note: 
- Some are rephrased as I understand them
- In retrospect, we could have used the brand-new GitHub feature with the same name.

Occasionally asked questions to: [olivier@zen.ly](mailto:olivier@zen.ly) and [mehran@zen.ly](mailto:mehran@zen.ly)
I'm going to summarize what have been said over emails (and a rough translations), so that this can be useful if someone @ Zenly wants to reevaluate this test and / or improve it.

### Repository: GitHub, GitLab? private or public access?

> Github + Private Repo

Key Takeaway: self-explanatory.

### Commodity Hardware is a bit too vague, imho, wdym?

> A machine with 64/128 GB of RAM.

Key Takeaway: memory-wise, should be fine even for my naive (but rel. fast) implementation.

### Need to write tests, PBT, mutants?

> The bare-minimum to ensure things are working as expected in a single-threaded environment.
> It's just a tech. test and not something production-ready.
> No need for something over-engineered if we have tests plus tooling to benchmark your code on a laptop, this will be good enough.

Key Takeaway: just some tests, and a function to benchmark with enough data.

### Am I supposed to split business concerns, hex arch.?

> Ideally a library that addresses strictly the problem given in the instructions.

Key Takeaway: just a lib is fine.

### Phone Numbers Representation: Memory Requirement

> 7-8 bytes to represent an international phone number.

Key Takeaway: `uint64` is fine to a phone number.

### Phone Number Validation?

> You can use the lib phone number, this is the one we're using at work, but this is a bit far fetched (it is just a technical test)

Key Takeaway: phone validation not required, plus it adds some overhead to the benchmark.

### Suggestions based on a random selection?

> No need for a random matchmaking algorithm.

Key Takeaway: suggestions based on a deterministic algorithm.

### Programming Language Choice

> About the requirement, you can completely ignore the language overhead and GC impact since what we are looking for is the thought process, the code is just there to demonstrate that we got from the idea to a working implementation. We can always re-write the design in the another language and benchmark the result later.

Key Takeaway: you can use Go.

### Trying to avoid the in-memory caveats?

#### Clustering would be nice?

> More complex ideas like clustering would be impossible to implement in a few hours.

Key Takeaway: no clustering needed.

#### Me trying to shamelessly dodge the in-memory requirement, suggesting file memory map or half-storing stuff with SQLite solutions...

> Regarding the RAM vs SSD storage let's reason in latency percentile instead. Is it alright if the 25th percentile touches the disk? What about a 99th percentile?
> We have no constraints regarding where you read the data from as long as the service is behaving. To judge that you would need a few metrics.
> The timeout for this service is 2000 ms but this is the absolute maximum response time. Any more that that and the user will see an error screen so not acceptable even for the 99.9th percentile.

Key takeaway: algorithms must be as fast as possible.

### Overlapping contacts, can we avoid redundant information?

> However, you have mentioned that often contacts edges in a social graph are mutual (which our data confirms), so you should be able to reduce the memory footprint knowing that fact.
> We have not included this kind of very useful stats in the question since getting to that question in the first place is way more important than the implementation you do once you have asked the question.

Key takeaway: if you can come up with a smart way to organize data leveraging the high locality, that's great!

## Memory Considerations

For starters, it was asked:

> Ideally we want an algorithm that **answers these queries as fast as possible.**

### In the grand scheme of things

Worst-case scenario for naive impl:

- `100,000,000` nodes (ie. at least accounts for the all the keys of the hashmap)
- `50` edges on average
- `uint64` (ie. `8` bytes) for storing an intl. phone number
- `100,000,000 . 8 bytes = 800 megabytes`
- `50 . 2 (from and to) . 8 . 100,000,000 ~= 80 gigabytes`

Note: not even considering of the GC overhead and other memory shenanigans (ie. side allocations).

### Compression Schemes

In addition to what I'm about to mention below there are some other representations that might be worth testing
(indexed + offsets like CSR, CSC instead of just tweaking around hashmap,
each carrying their own sets of pros and cons (and sometimes, even only cons))

#### Binary Masks

The idea is described on [this "amazing" sketch of mine](./Sketches 1.jpg),
albeit needs to add two other columns for reversal functionalities.

There is a usual tradeoff speed / memory, with this method (write operations take longer, read ones too),
haven't had the time to impl. the full bulk of it.

Works better with high locality (a lot a mutual contacts given area in the social graph)

#### Roaring Bitmaps

- [Official website](https://roaringbitmap.org)
- Youtube introduction [![Roaring Bitmaps - Daniel Lemire](https://img.youtube.com/vi/ubykHUyNi_0/0.jpg)](https://www.youtube.com/watch?v=ubykHUyNi_0)
- [Arxiv pub.](https://arxiv.org/pdf/1603.06549.pdf)
- [Go lib.](https://github.com/RoaringBitmap/roaring)

Not a silver bullet per se, just avoid too much phone numbers redundancy within the same nodes playing with bits and
relevant (ie. optimized) different underlying containers.

"The lazy option" to decrease the memory footprint.

Write operations are noticeably slower, tho... the size is also much much lower.

Again, same tradeoff as mentioned above.

## My schedule

### Prep. aka the "Feasibility Study" step raising tons of stupid questions

| Date       | Duration | Description                     |
|------------|----------|---------------------------------|
| 12/04/2020 | 45 min   | Thinking about the problem      |
| 12/04/2020 | 30 min   | 1st Golang Implementation Draft |
| 12/06/2020 | 30 min   | Thinking about the problem      |
| 12/06/2020 | 20 min   | Asking questions                |
| 12/09/2020 | 30 min   | Thinking and replying email     |

### Implementing the solution (and these bits of doc)

| Date       | Duration | Description                     |
|------------|----------|---------------------------------|
| 12/14/2020 | 4 hours  | Coding                          |
| 12/14/2020 | 2 hours  | Documentation                   |


### Bonus: Documentation

I spent a bit more time writing documentation, I found that it has some relevance.

Maybe if I find the energy and the time for doing so, I will also add some bits of CI/CD.

## Some thoughts about this test:

- While I can understand that the candidate is expected to ask questions to optimize a bit the implementation.
  I think that putting right from the get-**Go** what is considered or qualified as a commodity hardware could be a nice addition to the instructions.
- There is also no standard deviation, just some vague info
  > A random distribution of contacts for each number"

  The average number of contacts per phone number is supposed to be 50 but what about the "margin"?
  I just then assumed it was 10.
- On a personal note, I found this test pretty cool, that was a long while that haven't had so much fun.


### Addendum - Goofingaroundery

In case my tech. test is a total failure, hope this can cheer you up!

Note: I just made up the second word

[![Goofingaroundery](https://img.youtube.com/vi/3MMMe1drnZY/0.jpg)](https://www.youtube.com/watch?v=3MMMe1drnZY)