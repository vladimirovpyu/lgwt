package main

import (
	"context"
	"errors"
	"sync"
)

type Result struct{}

type SearchFunc func(ctx context.Context, query string) (Result, error)

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
	if len(sfs) == 0 {
		return Result{}, errors.New("no search functions provided")
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Небуферизованные каналы
	resCh := make(chan Result)
	errCh := make(chan error)

	var wg sync.WaitGroup
	wg.Add(len(sfs))

	for _, sf := range sfs {
		go func(fn SearchFunc) {
			defer wg.Done()

			res, err := fn(ctx, query)
			if err != nil {
				// Если контекст уже отменен (кто-то другой нашел ответ), не зависаем на отправке ошибки
				select {
                    case errCh <- err:
                    case <-ctx.Done():
				}
				return
			}

			// Если основной цикл уже вышел по первому результату или таймауту, не зависаем на отправке
			select {
                case resCh <- res:
                case <-ctx.Done():
			}
		}(sf)
	}

	// Отдельная горутина закрывает каналы, когда все поисковики отработали или отменились
	go func() {
		wg.Wait()
		close(resCh)
		close(errCh)
	}()

	var lastErr error

	for {
		select {
		case <-ctx.Done():
			return Result{}, ctx.Err()
		case res, ok := <-resCh:
			if ok {
				return res, nil // defer cancel() отменит контекст и разблокирует остальные горутины
			}
		case err, ok := <-errCh:
			if !ok { // это сработает только если канал закрыт и мы вернем последнюю полученную ошибку
				// Канал errCh закрыт — значит, wg.Wait() завершился и ни одна горутина не отдала успех
				return Result{}, lastErr
			}
			lastErr = err
		}
	}
}
