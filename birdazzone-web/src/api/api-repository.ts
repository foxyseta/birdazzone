import { ApiManager, ApiResponse } from './api';
import type { ApiList } from './interfaces/api-list';
import type { TvGame } from './interfaces/tv-game';
import type { Tweet, User } from './interfaces/tweet';
import type { Results } from './interfaces/results';
import type { ChartEntry } from './interfaces/chart-entry';
import type { WordCloudOptions } from './interfaces/wordcloud-options';
import type { Solution } from './interfaces/solution';
import type { Politician } from './interfaces/politician';
import type { FantaTeam } from './interfaces/fanta-team';

export default class ApiRepository {
  private static readonly _BASE_URL = `http://${import.meta.env.VITE_SERVER_URL}/api/v1`;
  private static readonly _TV_GAMES = '/tvgames/';
  private static readonly _TV_GAMES_ID = '/tvgames/{0}';
  private static readonly _RESULTS_ID = '/tvgames/{0}/results?each={1}';
  private static readonly _TV_GAMES_ID_ATTEMPTS = '/tvgames/{0}/attempts?pageLength={1}&pageIndex={2}';
  private static readonly _RESULTS_ID_FROM_TO = '/tvgames/{0}/results?from={1}&to={2}&each={3}';
  private static readonly _RESULTS_ID_FROM = '/tvgames/{0}/results?from={1}&each={2}';
  private static readonly _TV_GAMES_ID_ATTEMPTS_FILTERED =
    '/tvgames/{0}/attempts?from={1}&to={2}&pageLength={3}&pageIndex={4}';
  private static readonly _TV_GAMES_ID_ATTEMPTS_STATS = '/tvgames/{0}/attempts/stats';
  private static readonly _TV_GAMES_ID_ATTEMPTS_STATS_FILTERED = '/tvgames/{0}/attempts/stats?from={1}&to={2}';
  private static readonly _TV_GAMES_ID_SOLUTION = '/tvgames/{0}/solution';

  private static readonly _FANTACITORIO_POLITICIANS = '/fantacitorio/politicians';
  private static readonly _FANTACITORIO_TEAMS = '/fantacitorio/teams';
  private static readonly _FANTACITORIO_TEAMS_USERNAME = '/fantacitorio/teams?username={0}';

  private static readonly _HELLO_USER = '/hello/user/{0}';

  private static readonly _CHESS_USER_GAME_TURN = '/chess/{0}/{1}/{2}';

  public static readonly getPoliticians = (): Promise<ApiResponse<Politician[]>> =>
    ApiManager.get<Politician[]>(this._BASE_URL + this._FANTACITORIO_POLITICIANS);

  private static readonly _TV_GAMES_ID_SOLUTION_FILTERED = '/tvgames/{0}/solution?date={1}';

  public static readonly getTvGames = (): Promise<ApiResponse<TvGame[]>> =>
    ApiManager.get<TvGame[]>(this._BASE_URL + this._TV_GAMES);

  public static readonly getTvGameById = (id: string): Promise<ApiResponse<TvGame>> =>
    ApiManager.get<TvGame>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID, id));

  public static readonly getListOfGuesser = (
    id: string,
    index: string,
    itemPerPage: string
  ): Promise<ApiResponse<ApiList<Tweet>>> =>
    ApiManager.get<ApiList<Tweet>>(
      this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS, id, itemPerPage, index)
    );

  public static readonly getListOfGuesserFiltered = (
    id: string,
    from: string,
    to: string,
    index: string,
    itemPerPage: string
  ): Promise<ApiResponse<ApiList<Tweet>>> =>
    ApiManager.get<ApiList<Tweet>>(
      this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS_FILTERED, id, from, to, itemPerPage, index)
    );

  public static readonly getResults = (id: string, each = `${60 * 60 * 24}`): Promise<ApiResponse<Results[]>> =>
    ApiManager.get<Results[]>(this.stringFormat(this._BASE_URL + this._RESULTS_ID, id, each));

  public static readonly getResultsFiltered = (
    id: string,
    from: string,
    to: string | null,
    each: string
  ): Promise<ApiResponse<Results[]>> =>
    to
      ? ApiManager.get<Results[]>(this.stringFormat(this._BASE_URL + this._RESULTS_ID_FROM_TO, id, from, to, each))
      : ApiManager.get<Results[]>(this.stringFormat(this._BASE_URL + this._RESULTS_ID_FROM, id, from, each));

  public static readonly getTvGameAttemptsStat = (
    id: string,
    from: string | null = null,
    to: string | null = null
  ): Promise<ApiResponse<ChartEntry[]>> =>
    from && to
      ? ApiManager.get<ChartEntry[]>(
          this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS_STATS_FILTERED, id, from, to)
        )
      : ApiManager.get<ChartEntry[]>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_ATTEMPTS_STATS, id));

  public static readonly getTvGameSolutionById = (id: string): Promise<ApiResponse<Solution>> =>
    ApiManager.get<Solution>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_SOLUTION, id));

  public static readonly getFantacitorioTeams = (username?: string) =>
    ApiManager.get<FantaTeam[]>(
      username
        ? this.stringFormat(this._BASE_URL + this._FANTACITORIO_TEAMS_USERNAME, username)
        : this.stringFormat(this._BASE_URL + this._FANTACITORIO_TEAMS)
    );

  public static readonly getTvGameSolutionByIdFiltered = (id: string, date: string): Promise<ApiResponse<Solution>> =>
    ApiManager.get<Solution>(this.stringFormat(this._BASE_URL + this._TV_GAMES_ID_SOLUTION_FILTERED, id, date));

  public static readonly getHelloUser = (username: string): Promise<ApiResponse<User>> =>
    ApiManager.get<User>(this.stringFormat(this._BASE_URL + this._HELLO_USER, username));

  public static readonly getChessMoves = (user: string, gameId: string, turn: string) =>
    ApiManager.get<string>(
      this.stringFormat(
        this._BASE_URL + this._CHESS_USER_GAME_TURN,
        user,
        new Date(parseInt(gameId)).toISOString(),
        turn
      )
    );

  /// Takes a string in input containing placeholders in the form of {n}, where
  /// n is a number >= 0. Then replace all the occurence of the {n} pattern with
  /// the n-th word of args field. For example:
  ///
  /// stringFormat("Hello {0} {1}", "wonderful", "world") = "Hello wonderful world"
  ///
  private static readonly stringFormat = (str: string, ...args: string[]) =>
    str.replace(/{(\d+)}/g, (match, number) => (typeof args[number] != 'undefined' ? args[number] : match));

  public static readonly postWordCloudData = async (options: WordCloudOptions): Promise<ApiResponse<string>> => {
    const url = 'https://quickchart.io/wordcloud';
    const config: RequestInit = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(options)
    };

    const response = await fetch(url, config);

    if (response.ok) {
      return new ApiResponse<string>(response.status, await response.text());
    } else {
      return new ApiResponse<string>(response.status, undefined, {
        code: 400,
        message: await response.text()
      });
    }
  };

  public static readonly getTwitterTimestamp = async (): Promise<Date> => {
    /*return fetch('https://api.twitter.com/1.1/help/languages.json', { method: 'HEAD', })
    .then(response => {
      const date = response.headers.get('date');
      if (!response.ok) {
        const error = response.status;
        return Promise.reject(error);
      }
      return date ? new Date(date): new Date(Date.now());
    })*/

    return new Date();
  };
}
