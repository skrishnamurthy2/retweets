package filelib

import (
  "io/ioutil"
  "encoding/json"
)

type ArrayString []string

func ReadSlice(filename string) ArrayString {

  bytes, err := ioutil.ReadFile(filename)

  if err != nil {
    return make([]string, 0)
  }

  var strings ArrayString
  json.Unmarshal(bytes, &strings)

  return strings
}

func WriteSlice(filename string, strings ArrayString) {

  bytes, err := json.Marshal(strings)

  if  err != nil {
    panic(err)
  }

  ioutil.WriteFile(filename, bytes, 0644)
}
