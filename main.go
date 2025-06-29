package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"io"
	"fmt"
	"bytes"
    "github.com/joho/godotenv"
)

type RequestBody struct {
	Name   string `json:"name"`
	Taijyu string `json:"taijyu"`
}

func recordHandler(w http.ResponseWriter, r *http.Request) {
	secretToken := os.Getenv("SECRET_TOKEN")
	authHeader := r.Header.Get("Authorization")

	if authHeader != "Bearer "+secretToken {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	var body RequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
		return
	}

	log.Printf("受け取った: %s, %s", body.Name, body.Taijyu)
	json.NewEncoder(w).Encode(map[string]string{"message": "受け取ったっちゃ！"})
}

func sendToNotion(w http.ResponseWriter, r *http.Request) {
	url := "https://api.notion.com/v1/pages"
	notionToken := os.Getenv("NOTION_TOKEN")
	dbId := os.Getenv("NOTION_DATABASE_CAT_WEIGHT_ID")

	// JSONボディから名前と体重を取得
	var input struct {
		Name   string  `json:"name"`
		Weight float64 `json:"weight"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Notion API 送信用ペイロード作成
	payload := map[string]interface{}{
		"parent": map[string]string{"database_id": dbId},
		"properties": map[string]interface{}{
			"名前": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"text": map[string]string{
							"content": input.Name,
						},
					},
				},
			},
			"体重": map[string]interface{}{
				"number": input.Weight,
			},
		},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+notionToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send request to Notion", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		resBody, _ := io.ReadAll(resp.Body)
		http.Error(w, "Notion API error: "+string(resBody), resp.StatusCode)
		return
	}

	fmt.Fprintln(w, "✅ Notionに登録できたっちゃ〜！")
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("😱 .envファイルが読み込めなかったっちゃ…")
    }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s...", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/record", recordHandler)
	mux.HandleFunc("/cat-weight", sendToNotion)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}