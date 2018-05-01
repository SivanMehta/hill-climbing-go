package main

import (
  "log"
  "math"
  "math/rand"
)

// define what types of things we can climb
type mountain func(float64) float64

const (
  workers int = 100
  max float64 = 10
  stepSize float64 = .001
  down float64 = -.001
)

func climb(done chan float64, f mountain) {
  seed := rand.Float64() * max
  right := seed + stepSize
  left := seed - stepSize

  for f(seed) < f(right) || f(seed) < f(left) {
    if(f(seed) < f(right)) { // climb to the right
      seed = right
    } else { // climb to the left
      seed = left
    }

    right = seed + stepSize
    left = seed - stepSize
  }

  done <- seed
}

func arrangeWorkers(name string, fn mountain) {
  localMaxes := make(chan float64, workers)
  for i := 0; i < workers; i++ {
    go climb(localMaxes, fn)
  }

  globalMax := float64(-100)
  localMax := float64(-99)

  for i := 0; i < workers; i ++ {
    localMax = <- localMaxes
    if(localMax > globalMax) {
      globalMax = localMax
    }
  }

  log.Println("(probable) global max for", name, "at x =", globalMax)

  close(localMaxes)
}

func sin(x float64) float64 {
  return math.Sin(x)
}

func cos(x float64) float64 {
  return math.Cos(x)
}

func tan(x float64) float64 {
  return math.Tan(x)
}

func main() {
  arrangeWorkers("math.Sin", sin)
  arrangeWorkers("math.Cos", cos)
  arrangeWorkers("math.Tan", tan)
}
