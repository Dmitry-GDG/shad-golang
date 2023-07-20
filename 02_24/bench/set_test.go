package bench

import (
	"strconv"
	"sync"
	"testing"
)

type Set struct {
	set map[interface{}]struct{}
	mu sync.Mutex
}

func (s *Set) Add(x interface{}) {
	s.mu.Lock()
	s.set[x] = struct{}{}
	s.mu.Unlock()
}

func (s *Set) Delete(x interface{}) {
	s.mu.Lock()
	delete(s.set, x)
	s.mu.Unlock()
}

func BenchmarkSetDelete(b *testing.B) {
	var testSet []string
	for i := 0; i < 1024; i++ { // заполняем новый сет 1024 элементами
		testSet = append(testSet, strconv.Itoa(i))
	}

	b.ResetTimer() // чтоб время создания сета не входило в бенчмарк

	for i := 0; i < b.N; i++ { // в цикле добавляем до b.N и удаляем
		b.StopTimer() // чтоб время считать отдельно для каждой итерации
		set := Set{set: make(map[interface{}]struct{})} // создаём пустой мап
		for _, elem := range testSet {
			set.Add(elem)
			// set.Add("1qc") // Ключ мб любого типа, обязательно сравниваемый
			// set.Add(1)
		}
		b.StartTimer() // чтоб время считать отдельно для каждой итерации

		for _, elem := range testSet {
			set.Delete(elem)
		}
	}
}

// go test -v -count=10 -bench=BenchmarkSetDelete$ ./set_test.go