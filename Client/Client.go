package main

import (
 "fmt"
 "net"
 "os"
)

func main() {
 // Подключаемся к серверу на порту 4545
 conn, err := net.Dial("tcp", "localhost:4545")
 if err != nil {
  fmt.Println("Ошибка при подключении к серверу:", err)
  os.Exit(1)
 }
 defer conn.Close()

 // Читаем ответ от сервера
 buffer := make([]byte, 2) // Размер буфера равен 2, так как мы ожидаем "ок"
 _, err = conn.Read(buffer)
 if err != nil {
  fmt.Println("Ошибка при чтении данных:", err)
  return
 }

 // Проверяем полученный ответ
 response := string(buffer)
 if response == "ok" {
  fmt.Println("Получен ответ от сервера:", response)
 } else {
  fmt.Println("Получен неожиданный ответ:", response)
 }
}
