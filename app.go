package main

import (
  "net/http"
  "html/template"
  "log"
  "os"

  "github.com/rendon/tw"
)

func getTwitterClient() *tw.Client {
  var twitterClient *tw.Client
  twitterClient = tw.NewClient()

  consumerKey    := os.Getenv("TWITTER_CONSUMER_KEY")
  consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")

  if err := twitterClient.SetKeys(consumerKey, consumerSecret); err != nil {
    log.Fatalf("Failed to get credentials: %s", err)
  }

  return twitterClient
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func index(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")

  htmlPage := "views/index.html"
  t, err := template.ParseFiles(htmlPage)
  checkError(err)

  t.Execute(w, nil)
}

func getTweets(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")

  username := r.FormValue("username")

  twitterClient := getTwitterClient()
  tweets, err := twitterClient.GetTweets(username, 10)

  if err != nil {
    log.Fatalf("Failed to load tweets from username: %s", username)
  }

  htmlPage := "views/result.html"
  t, err := template.ParseFiles(htmlPage)
  checkError(err)

  t.Execute(w, tweets)
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/get_tweets", getTweets)

  log.Fatal(http.ListenAndServe(":8080", nil))
}