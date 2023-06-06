package greetings

import (
  "time"
  "errors"
  "fmt"
  "math/rand"
)

func Hello( name string ) (string, error) {
  // If no name was given, return a value that embeds the name
  // in a greeting message.
  if name == "" {
    return "", errors.New("empty name")
  }
  // Return a greeting that embeds the name in a message.
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

  
// Hellos returns a map that associates each of the named people
// with a greeting message
func Hellos( names []string ) ( map[string]string, error ) {
  messages := make(map[string]string)
  // Loop through the received slice of names, calling
  // the Hello function to get a message for each name.
  for _, name := range names {
    message, err := Hello(name)
    if err != nil {
      return nil, err
    }
    // In the map, associate the retrieved message with
    // the name.
    messages[name] = message
  }
  return messages, nil
}


func randomFormat() string {

  rand.Seed(time.Now().UnixNano())

  formats := []string{
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
  }

  return formats[rand.Intn(len(formats))]

}
