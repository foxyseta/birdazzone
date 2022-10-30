import { ApiManager, ApiResponse } from "./api"
import type HelloBird from "./interfaces/hello-bird"

export default class ApiRepository {
  private static readonly _BASE_URL = "http://localhost:8080/api/v1"
  private static readonly _HELLO = "/hello"

  public static readonly getHelloBird = (): Promise<ApiResponse<HelloBird>> =>
    ApiManager.get<HelloBird>(this._BASE_URL + this._HELLO)
}
