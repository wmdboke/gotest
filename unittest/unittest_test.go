package unit

import "testing"

// 展示测试执行信息
// go test -v
// 测试单个文件
// go test -v unittest_test.go  unittest.go
// 测试单个函数
// go test -v -run="(?i)user"
// go test -v -run=TestAdd
// 测试文件夹嵌套
// go test -v ./...
// 性能测试
// go test -bench="."
// go test -bench="." -cpu=3 -benchtime=5s
// 测试代码覆盖率
// go test -v -coverprofile cover.out unittest_test.go unittest.go
// go tool cover -html=cover.out -o cover.html

func TestAdd(t *testing.T) {
	add(1, 2)
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}
