package main

import (
  "os"
  "log"

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