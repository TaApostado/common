package specification

type ISpecification[T any] interface {
	IsSatisfiedBy(T) bool
}

type Specification[T any] struct {
	ISpecification[T]
}

type predicate[T any] struct {
	function func(object T) bool
}

func Predicate[T any](f func(object T) bool) ISpecification[T] {
	return &predicate[T]{function: f}
}

func And[T any](specifications ...ISpecification[T]) ISpecification[T] {
	return &andSpecification[T]{
		specifications: specifications,
	}
}

func Or[T any](specifications ...ISpecification[T]) ISpecification[T] {
	return &orSpecification[T]{
		specifications: specifications,
	}
}

func Not[T any](specification ISpecification[T]) ISpecification[T] {
	return &notSpecification[T]{
		specification: specification,
	}
}

type andSpecification[T any] struct {
	specifications []ISpecification[T]
}

func (self andSpecification[T]) IsSatisfiedBy(object T) bool {
	for _, specification := range self.specifications {
		if !specification.IsSatisfiedBy(object) {
			return false
		}
	}
	return true
}

type orSpecification[T any] struct {
	specifications []ISpecification[T]
}

func (self orSpecification[T]) IsSatisfiedBy(object T) bool {
	for _, specification := range self.specifications {
		if specification.IsSatisfiedBy(object) {
			return true
		}
	}
	return false
}

type notSpecification[T any] struct {
	specification ISpecification[T]
}

func (self notSpecification[T]) IsSatisfiedBy(object T) bool {
	return !self.specification.IsSatisfiedBy(object)
}

func (self *predicate[T]) IsSatisfiedBy(object T) bool {
	return self.function(object)
}
