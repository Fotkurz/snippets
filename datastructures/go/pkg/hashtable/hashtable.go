package hashtable

type HashTable struct {
	// arr is the underlying array implementation of the hash table
	arr []any
	len int
}

func NewHashTable() *HashTable {
	return &HashTable{arr: []any{}}
}

// hashingFunc produces the hashes which translates to the arr index
func (h *HashTable) hashingFunc(v any) int {

	h.arr = append(h.arr)
}

// Set adds a new key value pair to the hashtable
func (h *HashTable) Set(key, value any) {
	hashed := h.hashingFunc(key)

	if len(h.arr) >= hashed {
		h.arr[hashed] = value
	} else {
		h.arr = append(h.arr, value)
	}
}

// Get returns the value of the key or nil if no value
func (h *HashTable) Get(key any) (value any) {
	hashed := h.hashingFunc(key)
	if len(h)
}
