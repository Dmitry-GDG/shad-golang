package bench

import (
	"testing"
	// "encoding/json"
	"strconv"
)

type testStruct struct {
	X int
	Y string
}

// func (t *testStruct) ToJSON() ([]byte,error) {
// 	return json.Marshal(t)
// }

// // альтернативная реализация: мы вручную записали JSON через конкатенацию строк
func (t *testStruct) ToJSON() ([]byte,error) {
	return []byte(`{"X": ` + strconv.Itoa(t.X) + `, "Y": ` + t.Y + `}`), nil
}

func BenchmarkToJSON(b *testing.B) {
	tmp := &testStruct{X: 1, Y: "string"}
	js, err := tmp.ToJSON()
	if err != nil {
		b.Fatal(err)
	}

	b.SetBytes(int64(len(js)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := tmp.ToJSON(); err != nil {
			b.Fatal(err)
		}
	}
}

// НЕ работает запись в файл в терминале VS Code!!!, записывать в файлы из штатного Терминала!
// go test -v -bench=BenchmarkToJSON$ -count=10 ./tojson_test.go > /tmp/concat.txt
// go test -v -bench=BenchmarkToJSON$ -count=10 ./tojson_test.go > /tmp/json.txt
// cat /tmp/concat.txt

