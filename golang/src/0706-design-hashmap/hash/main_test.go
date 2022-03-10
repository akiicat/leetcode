package main
import "testing"
import . "main/pkg/testing_helper"
import "fmt"

func TestMyHashMap(_t *testing.T) {
  i1 := []string{"MyHashMap","put","put","get","get","put","get", "remove", "get"}
  i2 := [][]int{[]int{},[]int{1,1},[]int{2,2},[]int{1},[]int{3},[]int{2,1},[]int{2},[]int{2},[]int{2}}
  o := []string{"null","null","null","1","-1","null","1","null","-1"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

  i1 = []string{"MyHashMap","remove","put","remove","remove","get","remove","put","get","remove","put","put","put","put","put","put","put","put","put","put","put","remove","put","put","get","put","get","put","put","get","put","remove","remove","put","put","get","remove","put","put","put","get","put","put","remove","put","remove","remove","remove","put","remove","get","put","put","put","put","remove","put","get","put","put","get","put","remove","get","get","remove","put","put","put","put","put","put","get","get","remove","put","put","put","put","get","remove","put","put","put","put","put","put","put","put","put","put","remove","remove","get","remove","put","put","remove","get","put","put"}
  i2 = [][]int{[]int{},[]int{27},[]int{65,65},[]int{19},[]int{0},[]int{18},[]int{3},[]int{42,0},[]int{19},[]int{42},[]int{17,90},[]int{31,76},[]int{48,71},[]int{5,50},[]int{7,68},[]int{73,74},[]int{85,18},[]int{74,95},[]int{84,82},[]int{59,29},[]int{71,71},[]int{42},[]int{51,40},[]int{33,76},[]int{17},[]int{89,95},[]int{95},[]int{30,31},[]int{37,99},[]int{51},[]int{95,35},[]int{65},[]int{81},[]int{61,46},[]int{50,33},[]int{59},[]int{5},[]int{75,89},[]int{80,17},[]int{35,94},[]int{80},[]int{19,68},[]int{13,17},[]int{70},[]int{28,35},[]int{99},[]int{37},[]int{13},[]int{90,83},[]int{41},[]int{50},[]int{29,98},[]int{54,72},[]int{6,8},[]int{51,88},[]int{13},[]int{8,22},[]int{85},[]int{31,22},[]int{60,9},[]int{96},[]int{6,35},[]int{54},[]int{15},[]int{28},[]int{51},[]int{80,69},[]int{58,92},[]int{13,12},[]int{91,56},[]int{83,52},[]int{8,48},[]int{62},[]int{54},[]int{25},[]int{36,4},[]int{67,68},[]int{83,36},[]int{47,58},[]int{82},[]int{36},[]int{30,85},[]int{33,87},[]int{42,18},[]int{68,83},[]int{50,53},[]int{32,78},[]int{48,90},[]int{97,95},[]int{13,8},[]int{15,7},[]int{5},[]int{42},[]int{20},[]int{65},[]int{57,9},[]int{2,41},[]int{6},[]int{33},[]int{16,44},[]int{95,30}}
  o = []string{"null","null","null","null","null","-1","null","null","-1","null","null","null","null","null","null","null","null","null","null","null","null","null","null","null","90","null","-1","null","null","40","null","null","null","null","null","29","null","null","null","null","17","null","null","null","null","null","null","null","null","null","33","null","null","null","null","null","null","18","null","null","-1","null","null","-1","35","null","null","null","null","null","null","null","-1","-1","null","null","null","null","null","-1","null","null","null","null","null","null","null","null","null","null","null","null","null","-1","null","null","null","null","87","null","null"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))

  i1 = []string{"MyHashMap","put","get","remove","get"}
  i2 = [][]int{[]int{},[]int{0,42},[]int{0},[]int{0},[]int{0}}
  o = []string{"null","null","42","null","-1"}
  T(_t, S(i1, i2), S(run(i1, i2)), S(o))
}

func run(intrs []string, values [][]int) []string {
  var obj MyHashMap
  rtn := make([]string, len(intrs))

  for i, intr := range intrs {
    switch intr {
    case "MyHashMap":
      obj = Constructor()
      rtn[i] = "null"
    case "put":
      obj.Put(values[i][0], values[i][1])
      rtn[i] = "null"
    case "get":
      rtn[i] = fmt.Sprintf("%d", obj.Get(values[i][0]))
    case "remove":
      obj.Remove(values[i][0])
      rtn[i] = "null"
    default:
      fmt.Printf("Command Not Found\n")
      rtn[i] = "null"
    }
  }
  return rtn
}

