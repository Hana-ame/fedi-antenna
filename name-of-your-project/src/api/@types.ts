export type Visibility = "public" | "unlisted" | "private" | "direct"

export interface IEmojis {
  shortcode: string;
  url: string;
  static_url: string;
  visible_in_picker: boolean;
}

export interface IFields {
  name: string;
  value: string;
  verified_at: null;
}

export interface IAccount {
  id: string;
  username: string;
  acct: string;
  display_name: string;
  locked: boolean;
  bot: boolean;
  discoverable: boolean;
  group: boolean;
  created_at: string;
  note: string;
  url: string;
  uri: string;
  avatar: string;
  avatar_static: string;
  header: string;
  header_static: string;
  followers_count: number;
  following_count: number;
  statuses_count: number;
  last_status_at: string;
  noindex: boolean;
  emojis: IEmojis[];
  roles: unknown[];
  fields: IFields[];
}

export interface IMediaAttachment {
  id: string;
  type: string;
  url: string;
  preview_url: string;
  remote_url: null;
  preview_remote_url: null;
  text_url: null;
  meta: {
    original: {
      width: number;
      height: number;
      size: string;
      aspect: number;
    };
    small: {
      width: number;
      height: number;
      size: string;
      aspect: number;
    };
  };
  description: null;
  blurhash: string;
}

export interface ICard {
  url: string;
  title: string;
  description: string;
  language: string;
  type: string;
  author_name: string;
  author_url: string;
  provider_name: string;
  provider_url: string;
  html: string;
  width: number;
  height: number;
  image: string;
  image_description: string;
  embed_url: string;
  blurhash: string;
  published_at: string;
}

export interface IPoll {
  id: string;
  expires_at: string;
  expired: boolean;
  multiple: boolean;
  votes_count: number;
  voters_count: number;
  voted: boolean;
  own_votes: unknown[];
  options: {
    title: string;
    votes_count: number;
  }[];
  emojis: IEmojis[];
}

export interface IStatus {
  id: string;
  created_at: string;
  in_reply_to_id: string | null;
  in_reply_to_account_id: string | null;
  sensitive: boolean;
  spoiler_text: string;
  visibility: Visibility;
  language: string;
  uri: string;
  url: string;
  replies_count: number;
  reblogs_count: number;
  favourites_count: number;
  edited_at: string | null;
  favourited: boolean;
  reblogged: boolean;
  muted: boolean;
  bookmarked: boolean;
  content: string;
  filtered: unknown[];
  reblog: IStatus;
  application: {
    name: string;
    website: string | null;
  } | null;
  account : IAccount;
  media_attachments: IMediaAttachment[];
  mentions: {
    id: string;
    username: string;
    url: string;
    acct: string;
  }[];
  tags: {
    name: string;
    url: string;
  }[];
  emojis: IEmojis[];
  card: ICard[] | null;
  poll: null;
}
