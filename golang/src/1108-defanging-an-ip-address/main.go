package main
import "bytes"
import "strings"

// T: O(1)
// M: O(1)
// -- start --

func defangIPaddr(address string) string {
  rtn, i := make([]rune, len(address)+6), 0
  for _, v := range address {
    if v == '.' {
      rtn[i] = '['
      i++
      rtn[i] = '.'
      i++
      rtn[i] = ']'
    } else {
      rtn[i] = v
    }
    i++
  }
  return string(rtn)
}

func defangIPaddrConcat(address string) string {
  for i := 0; i < len(address); i++ {
    if address[i] == '.' {
      address = address[:i] + "[.]" + address[i+1:]
      i += 3
    }
  }
  return address
}

func defangIPaddrBuilder(address string) string {
  var str strings.Builder

  for _, v := range address {
    if v == '.' {
      str.WriteString("[.]")
    } else {
      str.WriteRune(v)
    }
  }

  return str.String()
}


func defangIPaddrBuffer(address string) string {
  var buffer bytes.Buffer
  for _, char := range address {
    if string(char) == "." {
      buffer.WriteString("[.]")
    } else {
      buffer.WriteString(string(char))
    }
  }
  return buffer.String()
}

func defangIPaddrStrings(address string) string {
  return strings.Replace(address, ".", "[.]", -1)
}


// -- end --

