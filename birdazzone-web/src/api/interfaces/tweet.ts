export interface Tweet {
  text: string;
  author: User;
  created_at: string;
  coordinates?: Coordinates;
  metrics: Metrics;
  medias: string[];
}

export interface User {
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
