import camelcaseKeys from "camelcase-keys"

export class ApiResponse<T> {
  public data?: T
  public error?: any
  private _esit: boolean
  public statusCode: number

  public get esit(): boolean {
    return this._esit
  }

  constructor(statusCode: number, data?: T, error?: any) {
    this.data = data
    this.error = error
    this._esit = statusCode >= 200 && statusCode < 300
    this.statusCode = statusCode
  }
}

export class ApiManager {
  private static async generalRequest<T>(url: string, config: RequestInit): Promise<ApiResponse<T>> {
    const response = await fetch(url, config)

    const responseStatusCode = response.status
    if (response.ok) {            // success
      const responseBody = (camelcaseKeys(await response.json())) as T
      return new ApiResponse<T>(responseStatusCode, responseBody)
    }
    else {
      console.log(response.json)
      const responseError = await response.json()      // error
      return new ApiResponse<T>(responseStatusCode, undefined, responseError)
    }
  }

  private static readonly camelToSnakeCase = (str: string) => str.replace(/[A-Z]/g, letter => `_${letter.toLowerCase()}`);

  public static async get<T>(url: string): Promise<ApiResponse<T>> {
    const config: RequestInit = {
      method: 'GET',
      headers: { 'Accept': 'Application/Json' },
    }

    return await this.generalRequest<T>(url, config)
  }

  public static async post<T>(url: string, data: any): Promise<ApiResponse<T>> {
    const config: RequestInit = {
      method: 'POST',
      headers: {
        'Accept': 'Application/Json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }

    return await this.generalRequest<T>(url, config)
  }
}
