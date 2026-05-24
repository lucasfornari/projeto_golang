package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type helloResponse struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("failed to encode response: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "name is required"})
		return
	}

	writeJSON(w, http.StatusOK, helloResponse{Message: "Olá, " + name + "!"})
}

func findClientDir() (string, error) {
	roots := []string{}

	if cwd, err := os.Getwd(); err == nil {
		roots = append(roots, cwd)
	}

	if exe, err := os.Executable(); err == nil {
		roots = append(roots, filepath.Dir(exe))
	}

	for _, root := range roots {
		dir := root
		for i := 0; i < 6; i++ {
			clientDir := filepath.Join(dir, "client")
			if info, err := os.Stat(clientDir); err == nil && info.IsDir() {
				return clientDir, nil
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				break
			}
			dir = parent
		}
	}

	return "", fmt.Errorf("pasta client não encontrada; execute a partir da raiz do repositório")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", helloHandler)

	clientDir, err := findClientDir()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Servindo front-end em %s", clientDir)
	mux.Handle("/", http.FileServer(http.Dir(clientDir)))

	addr := ":8080"
	log.Printf("Servidor iniciado em http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
