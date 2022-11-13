import { ApiManager, ApiResponse } from "./api"
import type { ApiList } from "./interfaces/api-list"
import type HelloBird from "./interfaces/hello-bird"
import type { TvGame } from "./interfaces/tv-game"
import type { Tweet } from "./interfaces/tweet"

export default class ApiRepository {
  private static readonly _BASE_URL = "http://localhost:8080/api/v1"
  private static readonly _HELLO = "/hello"
  private static readonly _TV_GAMES = "/tv-games"
  private static readonly _TV_GAMES_ID = "/tv-games/{0}"
  private static readonly _LIST_GUESSER = "/tvgames/{0}/attempts"

  public static readonly getHelloBird = (): Promise<ApiResponse<HelloBird>> =>
    ApiManager.get<HelloBird>(this._BASE_URL + this._HELLO)

  public static readonly getTvGames = (): Promise<ApiResponse<TvGame[]>> =>
    ApiManager.get<TvGame[]>(this._BASE_URL + this._TV_GAMES)

  public static readonly getTvGameById = (id: string): Promise<ApiResponse<TvGame>> =>
    ApiManager.get<TvGame>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID, id))

  public static readonly getListOfGuesser = (id: string): Promise<ApiResponse<ApiList<Tweet>>> =>
    ApiManager.get<ApiList<Tweet>>(this.stringFormat(this._BASE_URL + this._LIST_GUESSER, id))

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
}
