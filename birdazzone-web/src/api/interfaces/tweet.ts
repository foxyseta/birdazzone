export interface TweetList {
  tweets: { [key: string]: Tweet };
}

export interface Tweet {
  AttachmentMedia: any[];
  AttachmentPolls: any[];
  InReplyUser: any | null;
  Media: any | null;
  Mentions: any[];
  Place: any | null;
  Poll: any | null;
  ReferencedTweets: any[];
  Tweet: TweetClass;
  User: any | null;
}

export interface TweetClass {
  attachments: Attachments;
  author_id: string;
  context_annotations: any | null;
  conversation_id: string;
  created_at: Date;
  entities: Entities;
  geo: Geo;
  id: string;
  in_reply_to_user_id: string;
  lang: Lang;
  non_public_metrics: Metrics;
  organic_metrics: Metrics;
  possiby_sensitive: boolean;
  promoted_metrics: Metrics;
  public_metrics: Metrics;
  referenced_tweets: any | null;
  source: string;
  text: string;
  withheld: Withheld;
}

export interface Attachments {
  media_keys: any | null;
  poll_ids: any | null;
}

export interface Entities {
  annotations: any | null;
  cashtags: any | null;
  hashtags: any | null;
  mentions: any | null;
  urls: any | null;
}

export interface Geo {
  coordinates: Coordinates;
  place_id: string;
}

export interface Coordinates {
  coordinates: any | null;
  type: string;
}

export enum Lang {
  It = "it",
  Pt = "pt",
}

export interface Metrics {
  impression_count: number;
  like_count: number;
  quote_count: number;
  reply_count: number;
  retweet_count: number;
  url_link_clicks: number;
  user_profile_clicks: number;
}

export interface Withheld {
  copyright: boolean;
  country_codes: any | null;
}

