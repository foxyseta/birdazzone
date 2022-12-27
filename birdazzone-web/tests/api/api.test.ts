import { ApiResponse } from '/src/api/api';
import { test, assert, expect } from 'vitest';

test('Successfull ApiResponse', () => {
  // Arrange
  const STATUS_CODE = 200,
    DATA = 42;

  // Act
  let ar = new ApiResponse<number>(STATUS_CODE, DATA, null);

  // Assert
  expect(ar.data).toBe(DATA);
  expect(ar.error).toBeNull();
  expect(ar.statusCode).toBe(STATUS_CODE);
  assert(ar.esit);
});

test('Unsuccessfull ApiResponse', () => {
  // Arrange
  const STATUS_CODE = 404,
    DATA = 'boh',
    ERROR = 'Not found!';

  // Act
  let ar = new ApiResponse<string>(STATUS_CODE, DATA, ERROR);

  // Assert
  expect(ar.data).toBe(DATA);
  expect(ar.error).toBe(ERROR);
  expect(ar.statusCode).toBe(STATUS_CODE);
  assert(!ar.esit);
});
