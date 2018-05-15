package twitterlib

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "net/url"
  "encoding/json"
  "strings"
)

const (
  TWITTER_BASE = "https://api.twitter.com"
  TWITTER_AUTH = "/oauth2/token"
  TWITTER_SEARCH = "/1.1/search/tweets.json"
  TWITTER_RETWEET = "/1.1/statuses/retweets/%s.json"
  TWITTER_MAX = "100"
)

type AccessToken struct {
  TokenType    string `json:"token_type"`
  AccessToken string `json:"access_token"`
}

type Retweets []struct {
  User struct {
	ScreenName string `json:"screen_name"`
  } `json:"user"`
}

func Auth(key string, secret string) string {

  httpClient := &http.Client{}
  req, _ := http.NewRequest("POST", TWITTER_BASE + TWITTER_AUTH, strings.NewReader("grant_type=client_credentials"))

  req.SetBasicAuth(key, secret)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8.")

  resp, err := httpClient.Do(req)

  if err != nil || resp.StatusCode != 200{
	fmt.Println(err.Error())
  }

  defer resp.Body.Close()

  bodyBytes, err := ioutil.ReadAll(resp.Body)

  if err != nil {
	panic(err)
  }

  var twitterToken AccessToken
  json.Unmarshal(bodyBytes, &twitterToken)

  return twitterToken.AccessToken

}

func Search(bearer string, searchString string) []string {

  httpClient := &http.Client{}
  req, _ := http.NewRequest("GET", TWITTER_BASE + TWITTER_SEARCH, nil)

  req.Header.Set("Authorization", "Bearer " + bearer)

  query := req.URL.Query()
  query.Add("q", url.QueryEscape(searchString))
  query.Add("count", TWITTER_MAX)
  req.URL.RawQuery = query.Encode()

  resp, err := httpClient.Do(req)

  if err != nil || resp.StatusCode != 200{
	fmt.Println(err.Error())
  }

  defer resp.Body.Close()
  bodyBytes, err := ioutil.ReadAll(resp.Body)

  if err != nil {
	panic(err)
  }

  fmt.Println(string(bodyBytes))

  return make([]string, 0)
}

func Retweet(bearer string, tweetId string) Retweets {

  searchPath := fmt.Sprintf(TWITTER_RETWEET, tweetId)
  httpClient := &http.Client{}
  req, _ := http.NewRequest("GET", TWITTER_BASE + searchPath, nil)

  req.Header.Set("Authorization", "Bearer " + bearer)
  query := req.URL.Query()
  query.Add("count", TWITTER_MAX)
  req.URL.RawQuery = query.Encode()

  resp, err := httpClient.Do(req)

  if err != nil || resp.StatusCode != 200{
	fmt.Println(err.Error())
  }

  defer resp.Body.Close()
  bodyBytes, err := ioutil.ReadAll(resp.Body)

  if err != nil {
	panic(err)
  }

  var retweets Retweets
  err = json.Unmarshal(bodyBytes, &retweets)

  if err != nil {
    return nil;
  }

  return retweets
}
