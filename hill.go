package main

import (
  "log"
  "math"
  "math/rand"
)

const (
  workers int = 15
  max float64 = 10
  stepSize float64 = .001
  down float64 = -.001
)

func climb() float64 {
  seed := rand.Float64() * max
  right := seed + stepSize
  left := seed - stepSize

  for math.Sin(seed) < math.Sin(right) || math.Sin(seed) < math.Sin(left) {
    if(math.Sin(seed) < math.Sin(right)) { // climb to the right
      seed = right
    } else { // climb to the left
      seed = left
    }
    right = seed + stepSize
    left = seed - stepSize
  }

  return seed
}

func main()  {
  for i := 0; i < 10; i++ {
    maximum := climb()
    log.Println(maximum / math.Pi)
  }
}
