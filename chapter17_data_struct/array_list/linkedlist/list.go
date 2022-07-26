package linkedlist

type any interface{}

type Element struct {
	prev, next *Element
	Value      any
	list       *List
}
type List struct {
	// sentinel element without logic
	root Element
	len  int
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0

	return l
}

// lazyInit lazily initializes a zero List value.
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) PushFront(v any) *Element {
	//l.lazyInit()
	return l.insert(&Element{Value: v}, &l.root)
}

func (l *List) PushBack(v any) *Element {
	//l.lazyInit()
	return l.insert(&Element{Value: v}, l.root.prev)
}

// insert e after at element
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++

	return e
}

func (l *List) remove(e *Element) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	e.list = nil
	l.len--
}
