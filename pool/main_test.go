package main

import (
	"testing"
)

// Однопоточный бенчмарк БЕЗ пула
func BenchmarkEncodeWithoutPool(b *testing.B) {
	payload := makePayload()
	b.ReportAllocs() // Включаем подсчет аллокаций памяти
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := encodeWithoutPool(payload)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Однопоточный бенчмарк С пулом
func BenchmarkEncodeWithPool(b *testing.B) {
	payload := makePayload()
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := encodeWithPool(payload)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Многопоточный бенчмарк БЕЗ пула (имитация HTTP-сервера)
func BenchmarkParallelWithoutPool(b *testing.B) {
	payload := makePayload()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := encodeWithoutPool(payload)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// Многопоточный бенчмарк С пулом (имитация HTTP-сервера)
func BenchmarkParallelWithPool(b *testing.B) {
	payload := makePayload()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := encodeWithPool(payload)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}


