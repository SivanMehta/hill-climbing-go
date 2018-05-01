package main

import (
  "log"
  "math"
  "math/rand"
)

const (
  workers int = 15000
  max float64 = 10
  stepSize float64 = .001
  down float64 = -.001
)

func climb(done chan float64) {
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

  done <- seed
}

func main()  {
  localMaxes := make(chan float64, workers)
  for i := 0; i < workers; i++ {
    go climb(localMaxes)
  }

  globalMax := float64(-100)
  localMax := float64(-99)

  for i := 0; i < workers; i ++ {
    localMax = <- localMaxes
    if(localMax > globalMax) {
      globalMax = localMax
    }
  }

  log.Println(globalMax)

  close(localMaxes)
}
