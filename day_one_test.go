package main

import "testing"

func TestDayOne(t *testing.T) {

	t.Run("part 1", func(t *testing.T) {
		t.Helper()
		got := getAnswer("xdesign_day_one.txt", 1)
		want := 250

		assertEquals(got, want, t)
	})

	t.Run("part 2", func(t *testing.T) {
		t.Helper()
		got := getAnswer("xdesign_day_one.txt", 2)
		want := 151

		assertEquals(got, want, t)
	})

}

func benchmarkDayOne(b *testing.B, part int) {
	for i := 0; i < b.N; i++ {
		getAnswer("xdesign_day_one.txt", part)
	}
}

func BenchmarkDayOnePartOne(b *testing.B) { benchmarkDayOne(b, 1) }
func BenchmarkDayOnePartTwo(b *testing.B) { benchmarkDayOne(b, 2) }

func assertEquals(got, want int, t *testing.T) {
	if got != want {
		t.Errorf("Got %d but want %d", got, want)
	}
}
