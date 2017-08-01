/*
go test -v -bench=. -benchtime=10s
http://joxi.ru/Dr8jWGzHkzzll2 - golang
http://joxi.ru/GrqxdQ8tNkkWeA - nginx proxy
go test -v -bench=. -benchtime=10s -benchmem
http://joxi.ru/Dr8jWGzHkzzll2 - golang
http://joxi.ru/xAe6NpZcYggDOr - nginx proxy

https://habrahabr.ru/post/268585/ Бенчмарки в Go

по результатам тестов видно, что nginx в 1,5-2 раза быстрее отдает статику, а вот
просто rest api быстрее отрабатывает golang примерно в 1,5-2 раза.
по выделению памяти все одинаково у голанг и nginx.
вывод - использовать nginx для статики только.

BenchmarkHello    10000000    282 ns/op
means that the loop ran 10000000 times at a speed of 282 ns (nanosecond) per loop.

BenchmarkSample 10000000 208 ns/op 32 B/op 2 allocs/op
количество байт и аллокаций памяти за итерацию

go test -bench=. -benchmem bench_test.go > new.txt
git stash
go test -bench=. -benchmem bench_test.go > old.txt

http://godoc.org/golang.org/x/tools/cmd/benchcmp
go get golang.org/x/tools/cmd/benchcmp

b.StopTimer()
b.StartTimer()
b.ResetTimer()

cpu и memory профили во время выполнения бенчмарков:
go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out bench_test.go
https://blog.golang.org/profiling-go-programs
 */
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
	"github.com/Zhanat87/go/app"
	//"github.com/Zhanat87/go/db"
	"encoding/json"
)

func TestLoadEnvFile(t *testing.T) {
	assert.Equal(t, "8080", os.Getenv("PORT"))
}

// ping
func TestPing(t *testing.T) {
	resp, err := http.Get(os.Getenv("API_BASE_URL") + "ping")
	if err != nil {
		t.Fatalf("test ping error: %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("test ping read body error: %s", err.Error())
	}
	assert.Equal(t, fmt.Sprintf("OK %s", app.Version), string(body))
}

func getPingResponse() (res []byte) {
	resp, err := http.Get(os.Getenv("API_BASE_URL") + "ping")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

func BenchmarkRestApi(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getPingResponse()
	}
}

// static
func TestStaticDefaultAvatar(t *testing.T) {
	resp, err := http.Get(os.Getenv("API_BASE_URL") + "static/img/default-avatar.jpg")
	if err != nil {
		t.Fatalf("test StaticDefaultAvatar error: %s", err.Error())
	}
	defer resp.Body.Close()
	assert.Equal(t, "image/jpeg", resp.Header.Get("Content-Type"))
}

func getStaticDefaultAvatarHeader() (string, error) {
	resp, err := http.Get(os.Getenv("API_BASE_URL") + "ping")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Header.Get("Content-Type"), nil
}

func BenchmarkStatic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getStaticDefaultAvatarHeader()
	}
}

// db
//func TestDb(t *testing.T) {
//	var email string
//	row := db.Connection.QueryRow("SELECT email FROM users WHERE username = '$1' LIMIT 1", "test")
//	err := row.Scan(&email)
//	assert.Equal(t, "test@test.com", email)
//	assert.Nil(t, err)
//}

func TestDbApi(t *testing.T) {
	resp, err := http.Get(os.Getenv("API_BASE_URL") + "user-email/test")
	if err != nil {
		t.Fatalf("test user-email error: %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("test user-email read body error: %s", err.Error())
	}

	var dat map[string]interface{}
	err = json.Unmarshal(body, &dat)

	assert.Equal(t, "test@test.com", dat["email"].(string))
	assert.Nil(t, err)
}

func getDbApiResponse() ([]byte, error) {
	resp, err := http.Get(os.Getenv("API_BASE_URL") + "user-email/test")
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func BenchmarkDbApi(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getDbApiResponse()
	}
}
