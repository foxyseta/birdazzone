export interface Tweet {
  text: string;
  author: User;
  createdAt: string;
  coordinates?: Coordinates;
  metrics: Metrics;
  medias: string[];
}

export interface User {
  username: string;
  name: string;
  profileImageUrl: string;
}

export interface Metrics {
  likeCount: number;
  replyCount: number;
  retweetCount: number;
}

export interface Coordinates {
  latitude: number;
  longitude: number;
}
