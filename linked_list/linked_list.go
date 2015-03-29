package linked_list

// Each element in the List
type ListElement struct {
	// Pointers to next and prev elements.
	next *ListElement

	// Pointer to the list where the element belongs
	list *List

	// The actual value stored in the list
	Value interface{}
}

// Get the next element
func (e *ListElement) Next() *ListElement {
	return e.next
}

// Adds a new value right next to the current element
func (e *ListElement) AddNext(v interface{}) bool {
	oldNext := e.next
	newNext := ListElement{oldNext, e.list, v}
	e.next = &newNext

	return true
}

// List represents a linked list
type List struct {
	// The root element
	root *ListElement

	// The current lenght of the list
	length int
}

// Instantiate a new List
func New() *List {
	l := new(List)
	l.length = 0
	return l
}

// Returns the length of the list
func (l *List) Length() int {
	return l.length
}

// Inserts a new value into the list
func (l *List) Insert(v interface{}) bool {
	e := new(ListElement)
	e.Value = v
	e.list = l

	if l.root == nil {
		l.root = e
	} else {
		l.lastElement().AddNext(v)
	}

	l.length++

	return true
}

// Returns the first ListElement of the list in order to be able to iterate
// over all elements
// @example
// for e := l.Iterate(); e != nil; e = e.Next() {
//     // Do something with e.Value
// }
func (l *List) Iterate() *ListElement {
	return l.root
}

// Returns the first value of the list
func (l *List) First() interface{} {
	if l.root != nil {
		return l.root.Value
	}

	return nil
}

// Return the last ListElement pointer of the list
func (l *List) lastElement() *ListElement {
	el := l.root
	for el.Next() != nil {
		el = el.Next()
	}
	return el
}

// Return the last value of the list
func (l *List) Last() interface{} {
	el := l.lastElement()

	if el != nil {
		return el.Value
	}

	return nil
}
