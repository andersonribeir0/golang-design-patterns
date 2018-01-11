package singleton

type Counter interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Counter {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
