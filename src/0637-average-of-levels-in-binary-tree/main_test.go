package main
import "testing"
import . "main/pkg/testing_helper"
import . "main/pkg/tree_node"

func TestAverageOfLevels(_t *testing.T) {
  i, o := NewTreeNode("3,9,20,15,7"), []float64{3.00000,14.50000,11.00000}
  T(_t, S(i), S(averageOfLevels(i)), S(o))
}

