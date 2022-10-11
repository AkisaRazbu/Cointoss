package main

import (
  "fmt"
  "math/rand"
  "time"
)

const (
  HEADS = 1
  TAILS = 0
)

func main() {
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  headsCounter := 0;
  tailsCounter := 0;

  for i := 0; i < 1000000; i++ { 

    x := r1.Intn(2)
    if x == HEADS {
      headsCounter++;
    } else if x == TAILS {
      tailsCounter++;
    }

  }
  
  fmt.Println("Heads: ", headsCounter)
  fmt.Println("Tails: ", tailsCounter)

  if(headsCounter > tailsCounter){
    fmt.Println("Difference: ", headsCounter - tailsCounter);
  }
  if(tailsCounter > headsCounter){
    fmt.Println("Difference: ", tailsCounter - headsCounter);
  }
}
