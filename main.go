package main

import (
  "fmt"
  "math/rand"
  "time"
)

const (
  HEADS = 1
  TAILS = 0
  RUNTIME = 1000000
)

func main() {
  // Generates the seed using nanoseconds and current time
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  // Init the counters
  headsCounter := 0;
  tailsCounter := 0;
  // For loop which runs (constant) RUNTIME times
  for i := 0; i < RUNTIME; i++ { 

    x := r1.Intn(2)
    // Checks if it's heads or tails
    if x == HEADS {
      // Increments the counter if it's heads
      headsCounter++;
    } else if x == TAILS {
      // Increments the counter if it's tails
      tailsCounter++;
    }

  }
  // Prints the counter 
  fmt.Println("Heads: ", headsCounter)
  fmt.Println("Tails: ", tailsCounter)

  // Checks to see the difference and makes sure it's always a positive integer
  if(headsCounter > tailsCounter){
    fmt.Println("Difference: ", headsCounter - tailsCounter);
  }

  if(tailsCounter > headsCounter){
    fmt.Println("Difference: ", tailsCounter - headsCounter);
  }

  // This is nigh-impossible to get on higher amounts of coinflips
  if(tailsCounter == headsCounter){
    fmt.Println("Difference: None");
  }
}
