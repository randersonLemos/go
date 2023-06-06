package main

import (
  "fmt"
  "log"
  "example.com/greetings"
)

func main() {
  // Set properties of the predefined Logger, including
  // the log entry prefix and a flag to disable priting
  // the time, source file, and line number
  log.SetPrefix("greetings: ")
  log.SetFlags(0)


  // Request a greeting message.
  message, err := greetings.Hello("Jana")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(message)


  // A slice og names.
  names := []string{"Victor", "Cesar", "Alan"}

  messages, err := greetings.Hellos(names)

  for _, name := range names {
    fmt.Println(messages[name])
  }

}
