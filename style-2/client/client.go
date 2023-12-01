package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println("Erro ao criar a requisição:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)

	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("Context Error:", err)
	}

	if err != nil {
		log.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Erro ao ler a resposta:", err)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		return
	}

	cotacao, ok := data["bid"]
	if !ok {
		log.Println("Valor de cotação inválido na resposta.")
		return
	}

	fileContent := fmt.Sprintf("Dólar: %s\n", cotacao)
	err = os.WriteFile("cotacao.txt", []byte(fileContent), 0644)
	if err != nil {
		log.Println("Erro ao salvar arquivo:", err)
		return
	}

	log.Println("Cotação salva com sucesso em cotacao.txt")
}
