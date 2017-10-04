package main

import (
  "errors"
)

/*
  Triangle Project Code.

  Triangle analyzes the lengths of the sides of a triangle
  (represented by a, b and c) and returns the type of triangle.

  It returns:
    'equilateral'  if all sides are equal
    'isosceles'    if exactly 2 sides are equal
    'scalene'      if no sides are equal

  The tests for this method can be found in
    about_triangle_project.go
  and
    about_triangle_project_2.go
*/

func Triangle(a, b, c int) (result string, err error) {
  if a <= 0 || b <= 0 || c <= 0 {
    return "error", errors.New("hello world")
  }
  if a + b <= c || a + c <= b || b + c <= a {
    return "error", errors.New("hello world")
  }
  if a == b && b == c {
    return "equilateral", nil
  }
  if a != b && b != c && a != c {
    return "scalene", nil
  }
  return "isosceles", nil
}
