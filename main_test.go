package main

import "testing"

func TestStripLines(t *testing.T) {
	in := []string{"  a ", " b", "  ", "c  "}
	got := stripLines(in, false)
	want := []string{"a", "b", "", "c"}
	if len(got) != 4 {
		t.Fatalf("期望 4 行, 得到 %v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("第 %d 行期望 %q 得到 %q", i, want[i], got[i])
		}
	}
}

func TestStripLinesDropBlank(t *testing.T) {
	in := []string{"  a ", "  ", "c  "}
	got := stripLines(in, true)
	if len(got) != 2 || got[0] != "a" || got[1] != "c" {
		t.Errorf("删空行期望 [a c], 得到 %v", got)
	}
}
