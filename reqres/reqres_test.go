package tests

import (
	"fmt"
	"sync"
	"testing"
	"net/http"
	"io"
	"time"
	"encoding/json"
)

type UserData struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type ReqResResponse struct {
	Data []UserData `json:"data"`
}

type InMemoryDB struct {
    mu sync.RWMutex
	storage sync.Map
}

func (db *InMemoryDB) Save(id int, email string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.storage.Store(id, email)
}

func (db *InMemoryDB) PrintAll() {
	db.mu.RLock()
	defer db.mu.RUnlock()

	fmt.Println("\n--- Данные в InMemory DB (Только ID и Email) ---")
	db.storage.Range(func(key, value any) bool {
		fmt.Printf("ID: %d -> Email: %s\n", key.(int), value.(string))
		return true
	})
}

func start() {
    db := &InMemoryDB{}
	pagesChannel := make(chan int, 10)
	var wg sync.WaitGroup
	pages := 10
	workers := 3

    wg.Add(workers)

	for page := 1; page <= pages; page++ {
		pagesChannel <- page
	}

    for id := 1; id <= workers; id++ {
		go worker(id, db, pagesChannel, &wg)
	}

    close(pagesChannel)

    wg.Wait()

    db.PrintAll()
}

func worker(id int, db *InMemoryDB, pagesChannel <- chan int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("%d worker start\n", id)

    for page := range pagesChannel {
        // 1. Создаем клиент с таймаутом
        client := &http.Client{
            Timeout: 10 * time.Second,
        }

        // 2. Создаем сам GET-запрос
        url := fmt.Sprintf("https://reqres.in/api/users?page=%d", page)
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            fmt.Println("Ошибка создания запроса:", err)
            return
        }

        // 3. Добавляем ваш заголовок x-api-key
        req.Header.Set("x-api-key", "free_user_3LPKSC2FBk3hO0KRS7RLg8TXMpw")

        // 4. Отправляем запрос через клиент
        resp, err := client.Do(req)
        if err != nil {
            fmt.Println("Ошибка отправки запроса:", err)
            return
        }

        // 5. Читаем ответ
        body, err := io.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Println("Ошибка чтения ответа:", err)
            return
        }

        // Парсим полученный JSON
        var apiResponse ReqResResponse
        if err := json.Unmarshal(body, &apiResponse); err != nil {
            fmt.Printf("[Воркер %d] Ошибка парсинга JSON страницы %d: %v\n", id, page, err)
            return
        }

        // Сохраняем выбранные атрибуты в in-memory DB
        for _, u := range apiResponse.Data {
            db.Save(u.ID, u.Email)
        }

        fmt.Printf("%d worker process page %d\n", id, page)
    }
}

func TestStart(t *testing.T) {
	start()
}

/*
free_user_3LPKSC2FBk3hO0KRS7RLg8TXMpw

{"page":2,"per_page":6,"total":12,"total_pages":2,"data":[{"id":7,"email":"michael.lawson@reqres.in","first_name":"Michael","last_name":"Lawson","avatar":"https://reqres.in/img/faces/7-image.jpg"},{"id":8,"email":"lindsay.ferguson@reqres.in","first_name":"Lindsay","last_name":"Ferguson","avatar":"https://reqres.in/img/faces/8-image.jpg"},{"id":9,"email":"tobias.funke@reqres.in","first_name":"Tobias","last_name":"Funke","avatar":"https://reqres.in/img/faces/9-image.jpg"},{"id":10,"email":"byron.fields@reqres.in","first_name":"Byron","last_name":"Fields","avatar":"https://reqres.in/img/faces/10-image.jpg"},{"id":11,"email":"george.edwards@reqres.in","first_name":"George","last_name":"Edwards","avatar":"https://reqres.in/img/faces/11-image.jpg"},{"id":12,"email":"rachel.howell@reqres.in","first_name":"Rachel","last_name":"Howell","avatar":"https://reqres.in/img/faces/12-image.jpg"}],"support":{"url":"https://reqres.in/go/posthog","text":"PostHog ingests and stores your analytics, errors, replays, and business data so you and your agents can query and act on it. Get started.","sponsored_by":"PostHog"},"_meta":{"powered_by":"ReqRes","docs_url":"https://app.reqres.in/documentation","upgrade_url":"https://app.reqres.in/upgrade","example_url":"https://app.reqres.in/examples/notes-app","variant":"key_v1","message":"You are on the shared demo (no personal key), which has a low daily limit. A free account gives you your own key with 250 requests a day, plus request logs and your own collections.","cta":{"label":"Get a free key","url":"https://app.reqres.in/?utm_source=reqres_api&utm_medium=json&utm_campaign=demo_meta_key"},"context":"keyless_success"}}
{"page":2,"per_page":6,"total":12,"total_pages":2,"data":[{"id":7,"email":"michael.lawson@reqres.in","first_name":"Michael","last_name":"Lawson","avatar":"https://reqres.in/img/faces/7-image.jpg"},{"id":8,"email":"lindsay.ferguson@reqres.in","first_name":"Lindsay","last_name":"Ferguson","avatar":"https://reqres.in/img/faces/8-image.jpg"},{"id":9,"email":"tobias.funke@reqres.in","first_name":"Tobias","last_name":"Funke","avatar":"https://reqres.in/img/faces/9-image.jpg"},{"id":10,"email":"byron.fields@reqres.in","first_name":"Byron","last_name":"Fields","avatar":"https://reqres.in/img/faces/10-image.jpg"},{"id":11,"email":"george.edwards@reqres.in","first_name":"George","last_name":"Edwards","avatar":"https://reqres.in/img/faces/11-image.jpg"},{"id":12,"email":"rachel.howell@reqres.in","first_name":"Rachel","last_name":"Howell","avatar":"https://reqres.in/img/faces/12-image.jpg"}],"support":{"url":"https://reqres.in/go/posthog","text":"PostHog ingests and stores your analytics, errors, replays, and business data so you and your agents can query and act on it. Get started.","sponsored_by":"PostHog"},"_meta":{"powered_by":"ReqRes","docs_url":"https://app.reqres.in/documentation","upgrade_url":"https://app.reqres.in/upgrade","example_url":"https://app.reqres.in/examples/notes-app","variant":"key_v1","message":"You are on the shared demo (no personal key), which has a low daily limit. A free account gives you your own key with 250 requests a day, plus request logs and your own collections.","cta":{"label":"Get a free key","url":"https://app.reqres.in/?utm_source=reqres_api&utm_medium=json&utm_campaign=demo_meta_key"},"context":"keyless_success"}}
{"page":2,"per_page":6,"total":12,"total_pages":2,"data":[{"id":7,"email":"michael.lawson@reqres.in","first_name":"Michael","last_name":"Lawson","avatar":"https://reqres.in/img/faces/7-image.jpg"},{"id":8,"email":"lindsay.ferguson@reqres.in","first_name":"Lindsay","last_name":"Ferguson","avatar":"https://reqres.in/img/faces/8-image.jpg"},{"id":9,"email":"tobias.funke@reqres.in","first_name":"Tobias","last_name":"Funke","avatar":"https://reqres.in/img/faces/9-image.jpg"},{"id":10,"email":"byron.fields@reqres.in","first_name":"Byron","last_name":"Fields","avatar":"https://reqres.in/img/faces/10-image.jpg"},{"id":11,"email":"george.edwards@reqres.in","first_name":"George","last_name":"Edwards","avatar":"https://reqres.in/img/faces/11-image.jpg"},{"id":12,"email":"rachel.howell@reqres.in","first_name":"Rachel","last_name":"Howell","avatar":"https://reqres.in/img/faces/12-image.jpg"}],"support":{"url":"https://reqres.in/go/posthog","text":"PostHog ingests and stores your analytics, errors, replays, and business data so you and your agents can query and act on it. Get started.","sponsored_by":"PostHog"},"_meta":{"powered_by":"ReqRes","docs_url":"https://app.reqres.in/documentation","upgrade_url":"https://app.reqres.in/upgrade","example_url":"https://app.reqres.in/examples/notes-app","variant":"key_v1","message":"You are on the shared demo (no personal key), which has a low daily limit. A free account gives you your own key with 250 requests a day, plus request logs and your own collections.","cta":{"label":"Get a free key","url":"https://app.reqres.in/?utm_source=reqres_api&utm_medium=json&utm_campaign=demo_meta_key"},"context":"keyless_success"}}
*/

