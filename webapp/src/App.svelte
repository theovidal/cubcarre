<script>
  import "./telegram-web-app"
  import config from './config'
  import scramble from './scrambles/index'

  let Telegram = window.Telegram.WebApp
  let validData = false
  
  let waiting = false
  let ready = false
  let solving = false
  let timeout = null
  let label = 'Waiting'
  let color = ''
  
  let startTime = 0
  let time = 0

  Telegram.ready()
  const params = new URLSearchParams(window.location.search)
  const cube = params.get('cube')
  $: data = time + (cube === null ? '' : `:${cube}`)
  $: sc = scramble(cube)

  function startWait() {
    if (!ready && !solving) {
      color = 'red'
      waiting = true
      timeout = setTimeout(() => {
        color = 'green'
        waiting = false
        ready = true
      }, 400)
    }
  }

  function startSolve() {
    color = ''
    if (waiting) {
      waiting = false
      clearTimeout(timeout)
    }
    if (ready) {
      label = 'Solving'
      ready = false
      startTime = performance.now()
      setTimeout(() => solving = true, 20)
    }
  }

  function stopSolve() {
    if (solving) {
      time = Math.floor(performance.now() - startTime)
      let count = time
      let display = ''

      let mill = count % 1000
      count = Math.floor(count/1000)

      let seconds = count % 60
      count = Math.floor(count/60)

      let minutes = count

      if (minutes > 0) {
        if (minutes < 10) display += '0'
        display += `${minutes}:`
      }

      if (seconds < 10) display += '0'
      display += `${seconds}.`

      if (mill < 10) display += '0'
      if (mill < 100) display += '0'
      display += `${mill}`

      label = display
      solving = false

      Telegram.HapticFeedback.notificationOccurred('success')

      let url = new URL(`http://localhost.com/?${Telegram.initData}`)

      let body = new FormData()
      body.append('cube', cube)
      body.append('scramble', sc)
      body.append('userID', JSON.parse(url.searchParams.get('user'))["id"])
      body.append('queryID', url.searchParams.get('query_id'))
      body.append('time', time.toString())
      
      fetch(`${config.botURL}/save`, {
        method: 'POST',
        body
      })
    }
  }

</script>

<svelte:window
  on:keydown|stopPropagation|preventDefault={e => e.code !== 'Space' || solving ? stopSolve() : startWait()}
  on:keyup={startSolve}/>

<main>
  <div id="controller"
    on:touchstart|stopPropagation|preventDefault={_ => solving ? stopSolve() : startWait()}
    on:touchend={startSolve}

    on:mousedown={startWait}
    on:mouseup={startSolve}
    on:click={stopSolve}
    />
  <div class="container">
    <span
      id="label"
      style:color={color}>
      {label}
  </span>
    {#if cube !== null && label == 'Waiting'}
      <p id="sublabel">{@html sc.replaceAll(/_(\d+)_/g, '<sub>$1</sub>')}</p>
    {/if}
  </div>
</main>

<style>
main {
  background-color: var(--tg-theme-bg-color);
}

#controller {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Old versions of Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none;
}

.container {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  min-height: 100vh;
}

#label {
  font-family: "Lucida Grande", "Lucida Sans Unicode", Arial, Helvetica, Verdana, sans-serif;
  font-size: 20vw;
  font-weight: bold;
  color: var(--tg-theme-text-color);
}

#sublabel {
  padding: 20px;
  font-size: 20px;
  text-align: center;
  color: var(--tg-theme-hint-color);
}
</style>
