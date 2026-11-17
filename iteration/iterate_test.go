package iteration

import "testing"

func TestRepeat(t *testing.T)
{
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
	    t.Errorf("expected %q but got %q", expected, repeated)
    }
}

func BenchmarkRepeat(b *testing.B)
{
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}

/*
 go test -bench=. -benchmem
#  goos: linux
#  goarch: amd64
#  pkg: lgwt/iteration
#  cpu: Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz
#  BenchmarkRepeat-8   	15527487	        69.23 ns/op	       8 B/op	       1 allocs/op
#  PASS
#  ok  	lgwt/iteration	1.162s
*/


Сделай то что я прислал, но в нескольких вариациях.

Вот еще очень важная задачка, почти на всех собесах она есть.

Напиши программу, которая принимает некие данные и обрабатывает в k потоков, каждая обработка занимает n (n > 1) секунд. данные которые не успевают обработаться должны складываться в очередь и обрабатываться позже.

