/*
For a given set of strings, find out the frequency of each letter in all the strings parallelly.
For example, if given the following input.,

[“quick”, “brown”, “fox”, “lazy”, “dog”]

The output should be.,
{
“a”: 1,
“b”: 1,
“c”: 1,
...
“z”: 1
}

*/

package main

import (
  "fmt"
  "sync"
  "encoding/json"
)

type Result struct {
  mu          sync.Mutex
  letterFreqs map[string]int
}

func (r *Result) inc(letter rune) {
  r.mu.Lock()
  defer r.mu.Unlock()
  r.letterFreqs[string(letter)]++
}

func main() {
  //strings := []string{"quick", "brown", "fox", "lazy", "dog"}
  strings := []string{"harry", "joe", "john"}
  fmt.Println("Input:", strings)
  strChan := make(chan string, 5)
  r := Result{letterFreqs: make(map[string]int)}
  var wg sync.WaitGroup
  wg.Add(len(strings))

  for _, str := range strings {
    strChan <- str
    // ideal to run in a worker group to handle larger arrays
    go func() {
      defer wg.Done()
      for _, letter := range <-strChan {
        r.inc(letter)
      }
    }()
  }

  wg.Wait()
  if j, err := json.Marshal(r.letterFreqs); err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", string(j))
  }
}
