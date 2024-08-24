# HashTable

HashTable (also known as hashmaps, dictionaries, map, associative array, etc...) is
one of the fundamentals data structures in computer science.

It works with a system of key / value pairs, and allows us to access values from 
an array like structure using something different than its index.

An example:

In the array below

```
arr = [["apple", 5], ["orange", 10], ["watermelon", 15]]

for n in arr { // O(n)
    if n[0] == "orange" {
        print(n[1])
        // output 10
    }
}

```

For finding the price of the orange, we would have to go
over all the array with a time complexity of O(n).

Hashtables defines a structure where we could simple ask
for the value of the orange key:

```
hashtable = {
    "apple": 5,
    "orange": 10,
    "watermelon": 15
}

print(hashtable["orange"])
// output 10
```

## How it works

HashTables works by implementing the so called hashing function.
The hashing function will get our key "orange" and tell us
which index the orange is in. 

This means the hash table implements a simple array below but
with all the index having their own hash.