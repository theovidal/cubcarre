import cubeScramble from './cube'
import megaminxScramble from './megaminx'
import pyraminxScramble from './pyraminx'
import skewbScramble from './skewb'
import squareOneScramble from './sq-1'

const scrambles = {
    'Megamx': megaminxScramble,
    'Sq-1': squareOneScramble,
    'Pyramx': pyraminxScramble,
    'Skewb': skewbScramble
}

/**
 * If needed, extracts the cube size ; if not a cube, finds the function 
 * @param {string} puzzle 
 * @returns {string} can be empty string if unknown puzzle
 */
export default function(puzzle) {
    const regex = /(\d+)x\1|(\d+)BLD/
    const cube = puzzle.match(regex)
    if (cube !== null) {
        const size = parseInt(cube[1] === undefined ? cube[2] : cube[1])
        return cubeScramble(size)
    }
    else {
        const fn = scrambles[puzzle]
        return fn === undefined ? '' : fn()
    }
}