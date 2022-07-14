import { random } from './utils'

function scramble() {
  let res = []
  for (let row = 0; row < 6; row++) {
    for (let i = 0; i < 10; i++) {
      let move = i % 2 === 0 ? 'R' : 'D'
      move += random(2) === 0 ? '‒‒' : '++'
      res.push(move)
    }
    res.push(`U${random(2) === 0 ? "'" : '&nbsp;'}<br/>`)
  }
  return res.join(' ')
}

export default scramble
