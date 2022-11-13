import { ApiManager, ApiResponse } from "./api"
import type HelloBird from "./interfaces/hello-bird"
import type { TweetList } from "./interfaces/tweet"
import type { TvGame } from "./interfaces/tv-game"
import type { Results } from "./interfaces/results"
import type { ChartEntry } from "./interfaces/chart-entry"
import type { WordCloudOptions } from "./interfaces/wordcloud-options"
import type { Solution } from "./interfaces/solution"

export default class ApiRepository {
  private static readonly _BASE_URL = "http://localhost:8080/api/v1"
  private static readonly _HELLO = "/hello"
  private static readonly _TWITTER = "/twitter/{0}"
  private static readonly _TV_GAMES = "/tvgames/"
  private static readonly _TV_GAMES_ID = "/tvgames/{0}"
  private static readonly _RESULTS_ID = "/tvgames/{0}/results"
  private static readonly _TV_GAMES_ID_ATTEMPTS_STATS = "/tvgames/{0}/attempts/stats"
  private static readonly _TV_GAMES_ID_SOLUTION = "/tvgames/{0}/solution"

  public static readonly getHelloBird = (): Promise<ApiResponse<HelloBird>> =>
    ApiManager.get<HelloBird>(this._BASE_URL + this._HELLO)

  public static readonly getTweetList = (keyword: string): Promise<ApiResponse<TweetList>> =>
    ApiManager.get<TweetList>(this.stringFormat(this._BASE_URL + this._TWITTER, keyword))

  public static readonly getTvGames = (): Promise<ApiResponse<TvGame[]>> =>
    ApiManager.get<TvGame[]>(this._BASE_URL + this._TV_GAMES)

  public static readonly getTvGameById = (id: string): Promise<ApiResponse<TvGame>> =>
    ApiManager.get<TvGame>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID, id))

  public static readonly getResults = (id: string): Promise<ApiResponse<Results>> =>
    ApiManager.get<Results>(this.stringFormat(this._BASE_URL + this._RESULTS_ID, id))

  public static readonly getTvGameAttemptsStat = (id: string): Promise<ApiResponse<ChartEntry[]>> =>
    ApiManager.get<ChartEntry[]>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS_STATS, id))

  public static readonly getTvGameSolutionById = (id: string): Promise<ApiResponse<Solution>> =>
    ApiManager.get<Solution>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_SOLUTION, id))

  /// Takes a string in input containing placeholders in the form of {n}, where
  /// n is a number >= 0. Then replace all the occurence of the {n} pattern with 
  /// the n-th word of args field. For example:
  ///
  /// stringFormat("Hello {0} {1}", "wonderful", "world") = "Hello wonderful world"
  ///
  private static readonly stringFormat = (str: string, ...args: string[]) =>
    str.replace(/{(\d+)}/g,
      (match, number) => typeof args[number] != 'undefined' ? args[number] : match
    )

  public static readonly postWordCloudData = async (options: WordCloudOptions): Promise<ApiResponse<string>> => {
    const url = "https://quickchart.io/wordcloud"
    const config: RequestInit = {
      method: 'POST',
      headers: { 'Content-Type': "application/json" },
      body: JSON.stringify(options)
    }

    const response = await fetch(url, config)

    if (response.ok) {
      return new ApiResponse<string>(response.status, await response.text())
    } else {
      return new ApiResponse<string>(response.status, undefined, await response.text())
    }
  }
}
