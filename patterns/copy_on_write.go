/*
В Go этот паттерн чаще всего называют Copy-on-Write (COW) или Read-Copy-Update (RCU) с использованием atomic.Pointer (в старом коде — atomic.Value).
Суть паттерна заключается в том, что читатели получают доступ к данным без блокировок (lock-free), а писатели создают новую копию структуры, изменяют её и атомарно подменяют указатель.

Как это работает на практике

• Чтение (Lock-free): Читающие горутины вызывают метод ptr.Load(), получают актуальный указатель на map и свободно читают из неё. Никакие мьютексы при этом не захватываются.
• Запись (Copy-on-Write): Пишущая горутина блокирует мьютекс (чтобы исключить гонку между несколькими писателями), делает поверхностную копию текущей мапы, вносит изменения в новую мапу и вызывает ptr.Store(newMap).
• Безопасность: Поскольку старая мапа больше никем не модифицируется, горутины, которые успели получить на неё указатель до подмены, безопасно дочитывают данные, после чего старую мапу забирает сборщик мусора (GC).
*/

package main

import (
	"sync"
	"sync/atomic"
)

type ConfigCache struct {
	mu   sync.Mutex
	data atomic.Pointer[map[string]string]
}

func NewConfigCache() *ConfigCache {
	c := &ConfigCache{}
	initial := make(map[string]string)
	c.data.Store(&initial)
	return c
}

// Get работает без блокировок (очень быстро)
func (c *ConfigCache) Get(key string) (string, bool) {
	m := *c.data.Load()
	val, ok := m[key]
	return val, ok
}

// Set копирует мапу и атомарно подменяет указатель
func (c *ConfigCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	oldMap := *c.data.Load()
	newMap := make(map[string]string, len(oldMap)+1)
	for k, v := range oldMap {
		newMap[k] = v
	}
	newMap[key] = value

	c.data.Store(&newMap)
}
