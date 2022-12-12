import { ApiManager, ApiResponse } from "./api"
import type { ApiList } from "./interfaces/api-list"
import type { TvGame } from "./interfaces/tv-game"
import type { Tweet } from "./interfaces/tweet"
import type { Results } from "./interfaces/results"
import type { ChartEntry } from "./interfaces/chart-entry"
import type { WordCloudOptions } from "./interfaces/wordcloud-options"
import type { Solution } from "./interfaces/solution"
import type { Politician } from "./interfaces/politician"

export default class ApiRepository {
  private static readonly _BASE_URL = `http://${import.meta.env.VITE_SERVER_URL}/api/v1`
  private static readonly _TV_GAMES = "/tvgames/"
  private static readonly _TV_GAMES_ID = "/tvgames/{0}"
  private static readonly _RESULTS_ID = "/tvgames/{0}/results"
  private static readonly _TV_GAMES_ID_ATTEMPTS = "/tvgames/{0}/attempts?pageLength={1}&pageIndex={2}"
  private static readonly _RESULTS_ID_FILTERED = "/tvgames/{0}/results?from={1}&to={2}&each={3}"
  private static readonly _TV_GAMES_ID_ATTEMPTS_FILTERED = "/tvgames/{0}/attempts?from={1}&to={2}&pageLength={3}&pageIndex={4}"
  private static readonly _TV_GAMES_ID_ATTEMPTS_STATS = "/tvgames/{0}/attempts/stats"
  private static readonly _TV_GAMES_ID_ATTEMPTS_STATS_FILTERED = "/tvgames/{0}/attempts/stats?from={1}&to={2}"
  private static readonly _TV_GAMES_ID_SOLUTION = "/tvgames/{0}/solution"

  private static readonly _FANTACITORIO_POLITICIANS = "/fantacitorio/politicians"

  public static readonly getPoliticians = (): Promise<ApiResponse<Politician[]>> =>
    ApiManager.get<Politician[]>(this._BASE_URL + this._FANTACITORIO_POLITICIANS)

  private static readonly _TV_GAMES_ID_SOLUTION_FILTERED = "/tvgames/{0}/solution?date={1}"


  public static readonly getTvGames = (): Promise<ApiResponse<TvGame[]>> =>
    ApiManager.get<TvGame[]>(this._BASE_URL + this._TV_GAMES)

  public static readonly getTvGameById = (id: string): Promise<ApiResponse<TvGame>> =>
    ApiManager.get<TvGame>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID, id))

  public static readonly getListOfGuesser = (id: string, index: string, itemPerPage: string): Promise<ApiResponse<ApiList<Tweet>>> =>
    ApiManager.get<ApiList<Tweet>>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS, id, itemPerPage, index))

  public static readonly getListOfGuesserFiltered = (id: string, from: string, to: string, index: string, itemPerPage: string): Promise<ApiResponse<ApiList<Tweet>>> =>
    ApiManager.get<ApiList<Tweet>>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS_FILTERED, id, from, to, itemPerPage, index))

  public static readonly getResults = (id: string): Promise<ApiResponse<Results[]>> =>
    ApiManager.get<Results[]>(this.stringFormat(this._BASE_URL + this._RESULTS_ID, id))

  public static readonly getResultsFiltered = (id: string, from: string, to: string): Promise<ApiResponse<Results[]>> =>
    ApiManager.get<Results[]>(this.stringFormat(this._BASE_URL + this._RESULTS_ID_FILTERED, id, from, to, '604800'))    // 604800: secondi in una settimana

  public static readonly getTvGameAttemptsStat = (id: string): Promise<ApiResponse<ChartEntry[]>> =>
    ApiManager.get<ChartEntry[]>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS_STATS, id))

  public static readonly getTvGameAttemptsStatFiltered = (id: string, from: string, to: string): Promise<ApiResponse<ChartEntry[]>> =>
    ApiManager.get<ChartEntry[]>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS_STATS_FILTERED, id, from, to))

  public static readonly getTvGameSolutionById = (id: string): Promise<ApiResponse<Solution>> =>
    ApiManager.get<Solution>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_SOLUTION, id))

  public static readonly getTvGameSolutionByIdFiltered = (id: string, date: string): Promise<ApiResponse<Solution>> =>
    ApiManager.get<Solution>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_SOLUTION_FILTERED, id, date))

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
      return new ApiResponse<string>(response.status, undefined, { code: 400, message: await response.text() })
    }
  }
}
