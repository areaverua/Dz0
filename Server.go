package main

import (
 "fmt"
 "net"
 "os"
)

func main() {
 // Создаем слушатель на порту 4545
 listener, err := net.Listen("tcp", ":4545")
 if err != nil {
  fmt.Println("Ошибка при создании слушателя:", err)
  os.Exit(1)
 }
 defer listener.Close()

 fmt.Println("Сервер запущен на порту 4545")

 for {
  // Принимаем входящее соединение
  conn, err := listener.Accept()
  if err != nil {
   fmt.Println("Ошибка при принятии соединения:", err)
   continue
  }

  go handleConnection(conn)
 }

}

func handleConnection(conn net.Conn) {
 defer conn.Close()

 // Отправляем ответ "ок"
 _, err := conn.Write([]byte("ok"))
 if err != nil {
  fmt.Println("Ошибка при отправке данных:", err)
  return
 }

 fmt.Println("Ответ отправлен, соединение закрыто.")
}
