import { ApiManager, ApiResponse } from "./api"
import type HelloBird from "./interfaces/hello-bird"
import type { TweetList } from "./interfaces/tweet"

export default class ApiRepository {
  private static readonly _BASE_URL = "http://localhost:8080/api/v1"
  private static readonly _HELLO = "/hello"
  private static readonly _TWITTER = "/twitter/{0}"

  public static readonly getHelloBird = (): Promise<ApiResponse<HelloBird>> =>
    ApiManager.get<HelloBird>(this._BASE_URL + this._HELLO)

  public static readonly getTweetList = (keyword: string): Promise<ApiResponse<TweetList>> =>
    ApiManager.get<TweetList>(this.stringFormat(this._BASE_URL + this._TWITTER, keyword))

  /// Takes a string in input containing placeholders in the form of {n}, where
  /// n is a number > 0. Then replace all the occurence of the {n} pattern with 
  /// the n-th word of args field. For example:
  ///
  /// stringFormat("Hello {0} {1}", "wonderful", "world") = "Hello wonderful world"
  ///
  private static readonly stringFormat = (str: string, ...args: string[]) =>
    str.replace(/{(\d+)}/g,
      (match, number) => typeof args[number] != 'undefined' ? args[number] : match
    )
}
