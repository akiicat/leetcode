package main
import "testing"
import . "main/pkg/testing_helper"

func TestFizzBuzz(_t *testing.T) {
  i, o := 15, []string{
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"}
  T(_t, S(i), S(fizzBuzz(i)), S(o))
}

func TestFizzBuzzMethodOne(_t *testing.T) {
  i, o := 15, []string{
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"}
  T(_t, S(i), S(fizzBuzzMethodOne(i)), S(o))
}
