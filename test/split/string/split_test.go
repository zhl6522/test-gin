package string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		str		string
		sep		string
		want	[]string
	}
	testGroup := map[string]testCase{
		"case 1":{"babcdbef", "b", []string{"", "a", "cd", "ef"}},
		"case 2":{"a:b:c", ":", []string{"a", "b", "c"}},
		"case 3":{"abcef", "bc", []string{"a", "ef"}},
		"case 4":{"中国有钱还有人", "有", []string{"中国", "钱还", "人"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {		// 子测试:go test -run=TestSplit/case_1
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v but got:%#v\n", tc.want, got)
			}
		})
	}
	// 将覆盖率相关的信息输出到当前文件夹下面的cover.out文件中:go test -cover -coverprofile=cover.out
	// 使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告:go tool cover -html=cover.out
	/*testGroup := []testCase{
		{"babcdbef", "b", []string{"", "a", "cd", "ef"}},
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"abcef", "bc", []string{"a", "ef"}},
		{"中国有钱还有人", "有", []string{"中国", "钱还", "人"}},
	}
	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%#v but got:%#v\n", tc.want, got)
		}
	}*/
}