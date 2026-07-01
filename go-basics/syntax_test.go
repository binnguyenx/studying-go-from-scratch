package gobasics

import "testing"

func TestFizzBuzz(t *testing.T) {
	if FizzBuzz(15) != "FizzBuzz" || FizzBuzz(3) != "Fizz" || FizzBuzz(7) != "7" {
		t.Fatal(FizzBuzz(15), FizzBuzz(3), FizzBuzz(7))
	}
}

func TestSumRangeEffectiveGo(t *testing.T) {
	if SumRange(9) != 45 { // 0+1+...+9
		t.Fatal(SumRange(9))
	}
}

func TestCountWhile(t *testing.T) {
	if CountWhile(8) != 4 {
		t.Fatal(CountWhile(8))
	}
}

func TestCopyMapRange(t *testing.T) {
	old := map[string]int{"a": 1, "b": 2}
	got := CopyMapRange(old)
	if len(got) != 2 || got["a"] != 1 || got["b"] != 2 {
		t.Fatalf("%v", got)
	}
}

func TestRangeStringRunes(t *testing.T) {
	// Effective Go: 日本 + invalid \x80 + 語
	s := "日本\x80語"
	runes := RangeStringRunes(s)
	if len(runes) < 4 {
		t.Fatalf("got %d runes", len(runes))
	}
	if runes[0].R != '日' || runes[0].Pos != 0 {
		t.Fatalf("first: %+v", runes[0])
	}
	if runes[2].R != '\uFFFD' {
		t.Fatalf("invalid UTF-8 should be replacement char, got %U", runes[2].R)
	}
}

func TestReverseSliceParallel(t *testing.T) {
	a := []int{1, 2, 3, 4}
	ReverseSliceParallel(a)
	want := []int{4, 3, 2, 1}
	for i := range want {
		if a[i] != want[i] {
			t.Fatalf("%v", a)
		}
	}
}

func TestBreakFromInfinite(t *testing.T) {
	if BreakFromInfinite(5) != 5 {
		t.Fatal(BreakFromInfinite(5))
	}
}

func TestReverseInPlace(t *testing.T) {
	s := []int{1, 2, 3, 4}
	ReverseInPlace(s)
	want := []int{4, 3, 2, 1}
	for i := range want {
		if s[i] != want[i] {
			t.Fatalf("%v", s)
		}
	}
}

func TestWordCount(t *testing.T) {
	m := WordCount("go is go")
	if m["go"] != 2 || m["is"] != 1 {
		t.Fatalf("%v", m)
	}
}

func TestCounterPointerReceiver(t *testing.T) {
	var c Counter
	c.Inc()
	c.Add(4)
	if c.Value() != 5 {
		t.Fatal(c.Value())
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Fatal(a, b)
	}
}

func TestDeferOrder(t *testing.T) {
	got := DeferOrder()
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Fatalf("%v", got)
	}
}

func TestDeferSum(t *testing.T) {
	if DeferSum([]int{1, 2, 3}) != 12 { // (1+2+3)*2
		t.Fatal()
	}
}

func TestExportedName(t *testing.T) {
	if ExportedName() != "exported" {
		t.Fatal()
	}
}

func TestArrayInferredLen(t *testing.T) {
	a := ArrayInferredLen()
	if len(a) != 4 || a[0] != 10 || a[3] != 40 {
		t.Fatalf("%v len=%d", a, len(a))
	}
}

func TestArrayPartialInit(t *testing.T) {
	a := ArrayPartialInit()
	want := [5]int{1, 2, 0, 0, 0}
	if a != want {
		t.Fatalf("got %v want %v", a, want)
	}
}

func TestArrayIndexInit(t *testing.T) {
	a := ArrayIndexInit()
	if a[2] != 99 || a[4] != 1 || a[0] != 0 {
		t.Fatalf("%v", a)
	}
}

func TestArrayCopyDemo(t *testing.T) {
	orig, copy := ArrayCopyDemo()
	if orig[0] != 1 || copy[0] != 99 {
		t.Fatalf("orig=%v copy=%v", orig, copy)
	}
}

func TestNilVsEmptySlice(t *testing.T) {
	nilS := NilSlice()
	emptyLit := EmptySliceLiteral()
	emptyMake := EmptySliceMake()

	if !SliceIsNil(nilS) {
		t.Fatal("nil slice should be nil")
	}
	if SliceIsNil(emptyLit) || SliceIsNil(emptyMake) {
		t.Fatal("empty slices should not be nil")
	}
	if len(nilS) != 0 || len(emptyLit) != 0 || len(emptyMake) != 0 {
		t.Fatal("all should have len 0")
	}
}

func TestEmptySliceMakeCap(t *testing.T) {
	s := EmptySliceMakeCap(8)
	if len(s) != 0 || cap(s) != 8 || SliceIsNil(s) {
		t.Fatalf("len=%d cap=%d nil=%v", len(s), cap(s), s == nil)
	}
}

func TestAppendNilAndEmpty(t *testing.T) {
	nilOut := AppendToEither(NilSlice(), 1)
	emptyOut := AppendToEither(EmptySliceLiteral(), 1)
	if len(nilOut) != 1 || len(emptyOut) != 1 || nilOut[0] != 1 || emptyOut[0] != 1 {
		t.Fatalf("nilOut=%v emptyOut=%v", nilOut, emptyOut)
	}
}

func TestLoopSliceSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	if SumSliceIndex(nums) != 10 || SumSliceRange(nums) != 10 {
		t.Fatal()
	}
}

func TestRangeValueCopyDoesNotMutate(t *testing.T) {
	nums := []int{1, 2, 3}
	FillSliceWrong(nums, 100)
	if nums[0] != 1 || nums[1] != 2 {
		t.Fatalf("wrong fill changed slice? %v", nums)
	}
	FillSliceCorrect(nums, 100)
	for _, v := range nums {
		if v != 100 {
			t.Fatalf("correct fill failed: %v", nums)
		}
	}
}

func TestSliceOfUserValue(t *testing.T) {
	users := UsersByValue()
	if len(users) != 2 || users[0].Name != "John" {
		t.Fatalf("%v", users)
	}
	RenameUserAt(users, 0, "Johnny")
	if users[0].Name != "Johnny" {
		t.Fatal(users[0].Name)
	}
}

func TestSliceOfUserPointer(t *testing.T) {
	ptrs := UsersByPointer()
	empty := EmptyUserPointers()
	if len(empty) != 0 || empty == nil {
		t.Fatal("empty pointer slice should be non-nil with len 0")
	}
	RenameUserPtr(ptrs[0], "Johnny")
	if ptrs[0].Name != "Johnny" {
		t.Fatal(ptrs[0].Name)
	}
	a, b, same := SharedPointerDemo()
	if !same || a.Name != b.Name {
		t.Fatalf("same=%v a=%v b=%v", same, a, b)
	}
	RenameUserPtr(a, "Shared")
	if b.Name != "Shared" {
		t.Fatal(b.Name)
	}
}
