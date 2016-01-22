package main

import (
  "net/http"
  "html/template"
  "log"
)

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
  checkError(err)

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