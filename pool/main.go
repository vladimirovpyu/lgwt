package main

import (
	"bytes"
	"encoding/json"
	"sync"
	"time"
)

type ResponsePayload struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

const maxBufferCap = 64 * 1024

var bufferPool = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 4096))
	},
}

// Обертка, объединяющая буфер и энкодер
type pooledEncoder struct {
	buf *bytes.Buffer
	enc *json.Encoder
}

var encoderPool = sync.Pool{
	New: func() any {
		buf := bytes.NewBuffer(make([]byte, 0, 4096))
		return &pooledEncoder{
			buf: buf,
			enc: json.NewEncoder(buf), // Создаем энкодер один раз вместе с буфером!
		}
	},
}

// Вариант 1: Без использования пула (классический json.Marshal)
func encodeWithoutPool(p ResponsePayload) ([]byte, error) {
	return json.Marshal(p)
}

// Вариант 2: С использованием sync.Pool
func encodeWithPool(p ResponsePayload) ([]byte, error) {
	// Достаем готовую связку
	item := encoderPool.Get().(*pooledEncoder)
	defer func() {
		if item.buf.Cap() <= maxBufferCap {
			item.buf.Reset()
			encoderPool.Put(item)
		}
	}()

	// Используем уже созданный энкодер — 0 новых аллокаций!
	if err := item.enc.Encode(p); err != nil {
		return nil, err
	}

	return item.buf.Bytes(), nil
}

func makePayload() ResponsePayload {
	return ResponsePayload{
		Status:    "ok",
		Timestamp: time.Now().Unix(),
		Message:   "Тестовое сообщение для проверки аллокаций памяти при сериализации JSON",
	}
}
