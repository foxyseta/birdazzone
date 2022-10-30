import { ApiManager, ApiResponse } from "./api"
import type HelloBird from "./interfaces/hello-bird"

const _BASE_URL = "http://localhost:8080/api/v1"
const _HELLO = "/hello"

export const getHelloBird = (): Promise<ApiResponse<HelloBird>> =>
  ApiManager.get<HelloBird>(_BASE_URL + _HELLO)
