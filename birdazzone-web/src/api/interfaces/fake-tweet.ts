export interface FakeTweet {
  user: string
  text: string
  metrics: Metrics
  created_at: Date
}

export interface Metrics {
  impression_count?: number;
  like_count?: number;
  quote_count?: number;
  reply_count?: number;
  retweet_count?: number;
  url_link_clicks?: number;
  user_profile_clicks?: number;
}
