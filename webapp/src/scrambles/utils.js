/**
 * Returns a random number between 0 and (n - 1)
 * 
 * @param {number} n
 */
function random(n) {
  return Math.floor(Math.random() * n)
}

/**
 * Gets a random move from a list, different from the precedent
 * 
 * @param {string} pred
 * @param {Array<String>} moves
 */
 function getRandomMove(pred, moves) {
  let res = ''
  do {
    res = moves[random(moves.length)]
  } while (res === pred)

  return res
}

/**
 * Generates a flexible scramble, for Cube and Pyraminx puzzles
 * 
 * @param {Array<String>} moves
 * @param {Array<String>} modifiers
 * @param {number} l
 */
function baseScramble(moves, modifiers, l, n = 2) {
  const nbMoves = moves.length
  let res = [moves[random(nbMoves)]]

  // Step 1 - generate base moves
  for (let i = 1; i < l; i++) {
    const pred = res[i - 1]
    res.push(getRandomMove(pred, moves))
  }

  // Step 2 (optional, if cube size is more than 3) - Get layer number
  if (n > 3) {
    res = res.map(el => {
      const layer = random(Math.floor(n/2)) + 1
      return layer > 1 ? el + `_${layer}_` : el
    })
  }

  // Step 3 - Apply a modifier or not (clockwise/counter-clockwise, and double for some puzzles)
  res = res.map(el => {
    const mod = modifiers[random(modifiers.length)]
    return el + mod
  })

  return res.join(' ')
}

export { random, baseScramble }
