package main

import (
  "bytes";
  "container/vector";
  "fmt";
  "io";
  "os";
  "strings";
);

func separateTweets(tweetString string) []string {
  return strings.Split(tweetString, "\n\n", 0);
}

func separateRetweets(tweets []string) ([]string, []string) {
  originalTweets := vector.NewStringVector(0);
  retweets := vector.NewStringVector(0);
  for i := 0; i < len(tweets); i++ {
    if strings.Index(tweets[i], "RT") == -1 {
      // Add to original tweets list
      originalTweets.Push(tweets[i]);
    } else {
      // Add to retweets list
      retweets.Push(tweets[i]);
    }
  }

  return originalTweets.Data(), retweets.Data();
}

func main() {
  fileBytes, error := io.ReadFile("timoreilly_tweets.txt");
  if fileBytes == nil {
    fmt.Printf("Can't open file. Error: %s\n", error.String());
    os.Exit(1);
  }

  fileContents := bytes.NewBuffer(fileBytes).String();
  tweets := separateTweets(fileContents);

  originalTweets, retweets := separateRetweets(tweets);
  fmt.Printf("%d tweets are original.\n%d tweets are retweets.\n\n",
             len(originalTweets), len(retweets));
}