// Provides a simple hashtable data structure that uses separate
// chaining.
// @example
//     h := hashtable.New()
//     h.Set("PHP", "7.0")
//     h.Set("Ruby", "4")
//     h.Set("Python", "3.0")
//     fmt.Println(h.Get("PHP")) // 7.0
//     fmt.Println(h.Get("Ruby")) // 4
//     fmt.Println(h.Get("Python")) // 3.0
//     fmt.Println(h.Get("Elixir")) // <nil>
//
package hashtable

import "github.com/Zizaco/go-abc/linked_list"

// The Hashtable itself
type Hashtable struct {
	// The bucket where the values will be stored. Its is an array
	// that will point to multiple linked lists. These linked lists
	// will contain hashtableItems
	bucket [100]*linked_list.List
}

// Instantiate a new Hashtable
func New() *Hashtable {
	return new(Hashtable)
}

// Stores a new hashtableItem into the Hashtable bucket.
func (h *Hashtable) Set(key string, value interface{}) bool {
	hashedKey := Hash(key)

	if h.bucket[hashedKey] == nil {
		h.bucket[hashedKey] = linked_list.New()
	}

	for e := h.bucket[hashedKey].Iterate(); e != nil; e = e.Next() {
		item := e.Value.(*hashtableItem)
		if item.key == key {
			item.value = value
			return true
		}
	}

	h.bucket[hashedKey].Insert(&hashtableItem{key, value})

	return true
}

// Retrieves the value for the given key
func (h *Hashtable) Get(key string) interface{} {
	list := h.bucket[Hash(key)]

	if list == nil {
		return nil
	}

	for e := list.Iterate(); e != nil; e = e.Next() {
		item := e.Value.(*hashtableItem)
		if item.key == key {
			return item.value
		}
	}

	return nil
}

// Each item in the linked lists of the bucket. This way we can
// store the original key that was assigned to that value.
type hashtableItem struct {
	key   string
	value interface{}
}

// Function that will resolve the given key to a bucket position.
func Hash(k string) int {
	return int(k[0]) % 100
}
