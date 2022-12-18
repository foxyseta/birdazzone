export interface Tweet {
  text: string;
  author: Author;
  created_at: string;
  coordinates?: Coordinates;
  metrics: Metrics;
  medias: string[];
}

export interface Author {
  username: string;
  name: string;
  profile_image_url: string;
}

export interface Metrics {
  like_count: number;
  reply_count: number;
  retweet_count: number;
}

export interface Coordinates {
  latitude: number;
  longitude: number;
}
