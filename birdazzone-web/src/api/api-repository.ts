import { ApiManager } from "./api"

interface HelloBird{
    message: string
}

export const getHelloBird =  () =>
    ApiManager.get<HelloBird>("http://localhost:8080/api/v1/hello")
