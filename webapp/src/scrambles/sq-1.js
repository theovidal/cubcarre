function randomMove() {
    return Math.floor(Math.random() * 12 - 6)
}

function scramble() {
    let res = ''
    for (let i = 0; i < 16; i++) {
        let one = 0
        let two = 0
        do {
            one = randomMove()
            two = randomMove()
        } while (one === 0 && two === 0)

        res += `${one},${two} / `
    }
    return res
}

export default scramble
