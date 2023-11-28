import { decodeMorseCode } from '../src/decodeMorseCode'

describe('decodeMorseCode', () => {
  test('converts morse code into natural language', () => {
    expect(decodeMorseCode('.... . -.--   .--- ..- -.. .')).toEqual('HEY JUDE');
    expect(decodeMorseCode('...---...')).toEqual('SOS');
    expect(decodeMorseCode('...   ---   ...')).toEqual('S O S');
  })
})