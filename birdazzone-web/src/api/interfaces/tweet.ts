export interface Tweet {
  text: string
  author: Author
  created_at: string
  metrics: Metrics
}

export interface Author {
  username: string
  name: string
  profile_image_url: string
}

export interface Metrics {
  like_count: number
  reply_count: number
  retweet_count: number
}
