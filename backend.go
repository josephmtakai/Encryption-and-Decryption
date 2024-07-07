package main

import (
  "crypto/aes"
  "crypto/cipher"
  "encoding/base64"
  "fmt"
  "io"
  "net/http"
)

// EncryptAES and DecryptAES functions as previously discussed...

func handleEncryptDecrypt(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  text := r.FormValue("text")
  key := []byte(r.FormValue("key"))

  if text == "" || len(key) == 0 {
    http.Error(w, "Text and key must be provided", http.StatusBadRequest)
    return
  }

  algorithm := r.FormValue("algorithm")
  var result string
  var err error

  switch algorithm {
  case "aes":
    result, err = EncryptAES([]byte(text), key)
    if err != nil {
      http.Error(w, fmt.Sprintf("Encryption error: %s", err), http.StatusInternalServerError)
      return
    }
  case "decrypt_aes":
    result, err = DecryptAES(text, key)
    if err != nil {
      http.Error(w, fmt.Sprintf("Decryption error: %s", err), http.StatusInternalServerError)
      return
    }
  default:
    http.Error(w, "Invalid algorithm", http.StatusBadRequest)
    return
  }

  w.Write([]byte(result))
}

func main() {
  http.HandleFunc("/encrypt-decrypt", handleEncryptDecrypt)
  fmt.Println("Server listening on port 8080...")
  http.ListenAndServe(":8080", nil)
}
