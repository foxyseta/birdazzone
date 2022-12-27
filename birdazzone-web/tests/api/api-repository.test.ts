import ApiRepository from '/src/api/api-repository';
import { test, expect } from 'vitest';

test('stringFormat works without additional arguments', () => {
  let testVar = 'var test';
  expect(ApiRepository.stringFormat(testVar)).equals(testVar);
  expect(ApiRepository.stringFormat(testVar, 'word1', 'word2')).equals(testVar);
});

test('stringFormat works with some arguments', () => {
  let testVar = 'Ciao {0}';
  let bellezze = 'bellezze';
  let solution = 'Ciao bellezze';
  expect(ApiRepository.stringFormat(testVar, bellezze)).equals(solution);
});
