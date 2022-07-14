import { baseScramble, random } from './utils'

const modifiers = ['', '2', "'"]

/**
 * @param {number} n
 */
function scrambleLength(n) {
  return (n < 6) ? [9, 25, 40, 60][n - 2] : n * 10 + 20
}

/**
 * @param {number} n
 */
function scramble(n) {
  const moves = n > 2 ? ['U', 'R', 'F', 'D', 'L', 'B'] : ['U', 'R', 'F']
  const l = scrambleLength(n)
  
  return baseScramble(moves, modifiers, l, n)
}

export default scramble
