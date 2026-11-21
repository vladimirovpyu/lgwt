type Set[T comparable] struct {
    m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(val T) {
    s.m[val] = struct{}{}
}

func (s *Set[T]) Has(val T) bool {
    _, exists := s.m[val]
    return exists
}

func (s *Set[T]) Delete(val T) {
    delete(s.m, val)
}

func (s *Set[T]) Size() int {
    return len(s.m)
}
