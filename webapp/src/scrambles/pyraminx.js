import { random, baseScramble } from './utils'

function scramble() {
    // Step 1 - Get number of tips to rotate
    const nbTips = random(4)
    const length = 12 - nbTips
    const modifiers = ['', "'"]

    let res = baseScramble(['U', 'B', 'L', 'R'], modifiers, length)
    let tips = ['u', 'l', 'r', 'b']
    for (let t = 0; t < nbTips; t++) {
        let i = random(tips.length)
        res += ` ${tips[i]}${modifiers[random(2)]}`
        tips.splice(i, 1)
    }
    return res
}

export default scramble