package main

import (
  "encoding/json"
  "fmt"
)

type Matrix struct {
  rows int
  cols int
  elems [][]int
}

func (m *Matrix) rowsCnt() int {
  return m.rows
}

func (m *Matrix) colsCnt() int {
  return m.cols
}

func (m *Matrix) setVal(i, j, val int) {
  m.elems[i][j] = val
}

func (m1 *Matrix) add(m2 *Matrix) Matrix {
  elems := make([][]int, m1.rows)
  for i := range elems {
    elems[i] = make([]int, m1.cols)
  }
  sumMatrix := Matrix{m1.rows, m2.rows, elems}
  for i, rows := range m1.elems {
    for j, val := range rows {
      sum := val + m2.elems[i][j]
      sumMatrix.setVal(i, j, sum)
    }
  }
  return sumMatrix
}

func (m *Matrix) toJson() (string, error) {
  if j, err := json.Marshal(m.elems); err != nil {
    return "", err
  } else {
    return string(j), nil
  }
}

func main() {
  m1_elems := [][]int{
    {1, 2, 3},
    {4, 5, 6},
  }
  m1 := Matrix{2, 3, m1_elems}
  m2_elems := [][]int{
    {1, 1, 1},
    {2, 2, 2},
  }
  m2 := Matrix{2, 3, m2_elems}
  fmt.Println("M1 Rows count:", m1.rowsCnt())
  fmt.Println("M2 Cols count:", m2.colsCnt())
  m3 := m1.add(&m2)
  if json, err := m3.toJson(); err != nil {
    fmt.Println("Error converting to Json:", err)
  } else {
    fmt.Println("Sum Matrix:", json)
  }
}
