package main
import "testing"
import . "main/pkg/testing_helper"

func TestFindRestaurant(_t *testing.T) {
  i1, i2, o := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}, []string{"Shogun"}
  T(_t, S(i1, i2), S(findRestaurant(i1, i2)), S(o))

  i1, i2, o = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}, []string{"Shogun"}
  T(_t, S(i1, i2), S(findRestaurant(i1, i2)), S(o))
}

func TestFindRestaurantNegative(_t *testing.T) {
  i1, i2, o := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}, []string{"Shogun"}
  T(_t, S(i1, i2), S(findRestaurantNegative(i1, i2)), S(o))

  i1, i2, o = []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}, []string{"Shogun"}
  T(_t, S(i1, i2), S(findRestaurantNegative(i1, i2)), S(o))
}

