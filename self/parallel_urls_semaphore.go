/*
	Параллельный обход URL / Fan-Out, Fan-In

• Суть: Дан список из 1000 URL. Нужно опросить их, делая не более 10 одновременных HTTP-запросов, собрать результаты и отменить все оставшиеся запросы, если хотя бы один вернул критическую ошибку (или если получен таймаут).
• Что проверяют: Паттерн семафора (буферизированный канал chan struct{}), errgroup (или ручную реализацию через context.WithCancel).
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func fetchAll(ctx context.Context, urls []string, limit int) ([]int, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		wg      sync.WaitGroup
		sem     = make(chan struct{}, limit)
		errCh   = make(chan error, 1)
		results = make([]int, len(urls))
	)

	for i, url := range urls {
		select {
		case <-ctx.Done():
			break
		case sem <- struct{}{}:
		}

		if ctx.Err() != nil {
			break
		}

		wg.Add(1)
		go func(idx int, url string) {
			defer wg.Done()
			defer func() { <-sem }()

			code, err := fetchUrl(ctx, url)

			if err != nil {
				select {
				case errCh <- err:
				default:
				}
				return
			}

			results[idx] = code

		}(i, url)
	}
	wg.Wait()

	select {
	case err := <-errCh:
		return nil, err
	default:
		return results, ctx.Err()
	}

}

func fetchUrl(ctx context.Context, url string) (code int, err error) {
	return 200, errors.New("fg")
}

func main() {
	ctx := context.Background()

	urls := make([]string, 10)
	for i := range urls {
		urls[i] = fmt.Sprintf("%d", i)
	}

	results, err := fetchAll(ctx, urls, 2)

	fmt.Println(results, err)
}
