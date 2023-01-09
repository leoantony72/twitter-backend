package model

import "time"

type Tweets struct{
    Id string `json:"id"`
    UserId string `json:"user_id"`
    Content string `json:"tweet_content"`
    CreatedAt time.Time `json:"created_at"`
    LikeCout int `json:"like_count"`
    RetweetCount int `json:"retweet_count"`
}

type Retweet struct{
    TweetId string `json:"tweet_id"`
    UserId string `json:"user_id"`
}

type Like struct{
    TweetId string `json:"tweet_id"`
    UserId  string `json:"user_id"`
}
