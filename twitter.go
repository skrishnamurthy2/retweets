package main

import (
  "io/ioutil"
  "encoding/json"
  "github.com/twitterusage/twitterlib"
  "github.com/twitterusage/filelib"
  "fmt"
)

//https://github.com/gophercises/twitter
//https://twitter.com/search-home

type twitterKey struct {
  Key    string `json:"key"`
  Secret string `json:"secret"`
}

func main() {

  users := filelib.ReadSlice("tweets.json")

  fmt.Println("Read total users ", len(users))

  key, secret := readTwitterKey()
  bearer := twitterlib.Auth(key, secret)

  twitterlib.Search(bearer, "zillow")

  retweets := twitterlib.Retweet(bearer, "995423374519259137")

  usersMerged := combineUsers(users, getUsers(retweets))

  filelib.WriteSlice("tweets.json", usersMerged)

  fmt.Println("Merged total users ", len(usersMerged))
}

func readTwitterKey() (string, string){
  jsonByte, err := ioutil.ReadFile("keys.json")

  if err != nil {
    return "", ""
  }

  var twitterInfo twitterKey
  json.Unmarshal(jsonByte, &twitterInfo)

  return twitterInfo.Key, twitterInfo.Secret
}

func getUsers(retweets twitterlib.Retweets) [] string {

  users := make([]string, 0)

  for i := 0; i < len(retweets); i++  {
    users = append(users, retweets[i].User.ScreenName)
  }

  return users
}

func combineUsers(users1 []string, users2 [] string) []string {
  users := make([]string, 0)
  usermap := make(map[string]bool)

  for _, user := range users1 {

    if usermap[user] == false {
      usermap[user] = true
    }

  }

  for _, user := range users2 {

    if usermap[user] == false {
      usermap[user] = true
    }

  }

  for key, _ := range usermap {

    users = append(users, key)

  }

  return users
}