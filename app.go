package main

import (
  "net/http"
  "html/template"
  "log"
  "os"
  "fmt"

  "github.com/rendon/tw"
)

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

func getTweets(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")

  twitterClient := getTwitterClient()
  user, err := twitterClient.GetUsersShow("lccezinha")
  if err != nil {
    log.Fatalf("Failed to get user: %s", err)
  }

  fmt.Printf("User ID: %d\n", user.ID)
  fmt.Printf("User name: %s\n", user.ScreenName)
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/get_tweets", getTweets)

  log.Fatal(http.ListenAndServe(":8080", nil))
}