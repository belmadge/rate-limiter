package handler

import (
	"net/http"
	"testing"
	"time"
)

func TestRateLimiterByIP(t *testing.T) {
	for i := 0; i < 6; i++ {
		resp, err := http.Get("http://localhost:8080/")
		if err != nil {
			t.Fatalf("Erro na requisição: %v", err)
		}

		if i >= 5 && resp.StatusCode != http.StatusTooManyRequests {
			t.Errorf("Esperado HTTP 429 para requisição %d, mas obteve %d", i+1, resp.StatusCode)
		} else if i < 5 && resp.StatusCode != http.StatusOK {
			t.Errorf("Esperado HTTP 200 para requisição %d, mas obteve %d", i+1, resp.StatusCode)
		}
		resp.Body.Close()
	}
}

func TestRateLimiterByToken(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
	if err != nil {
		t.Fatalf("Erro ao criar a requisição: %v", err)
	}
	req.Header.Add("API_KEY", "abc123")

	for i := 0; i < 11; i++ {
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Erro na requisição: %v", err)
		}

		if i >= 10 && resp.StatusCode != http.StatusTooManyRequests {
			t.Errorf("Esperado HTTP 429 para requisição %d, mas obteve %d", i+1, resp.StatusCode)
		} else if i < 10 && resp.StatusCode != http.StatusOK {
			t.Errorf("Esperado HTTP 200 para requisição %d, mas obteve %d", i+1, resp.StatusCode)
		}
		resp.Body.Close()
		time.Sleep(100 * time.Millisecond)
	}
}
