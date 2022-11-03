import { ApiResponse, ApiManager } from '/src/api/api'
import { test, assert, expect} from 'vitest'
import { createRouter, createWebHistory } from 'vue-router'
import HelloBird from '/src/components/HelloBird.vue'

test('Successfull ApiResponse', () => {
  const STATUS_CODE = 200, DATA = 42
  let ar = new ApiResponse<number>(STATUS_CODE, DATA, null)
  expect(ar.data).toBe(DATA)
  expect(ar.error).toBeNull()
  expect(ar.statusCode).toBe(STATUS_CODE)
  assert(ar.esit)
})

test('Unsuccessfull ApiResponse', () => {
  const STATUS_CODE = 404, DATA = "boh", ERROR = "Not found!"
  let ar = new ApiResponse<string>(STATUS_CODE, DATA, ERROR)
  expect(ar.data).toBe(DATA)
  expect(ar.error).toBe(ERROR)
  expect(ar.statusCode).toBe(STATUS_CODE)
  assert(!ar.esit)
})
