<script>
  import "./telegram-web-app"
  let Telegram = window.Telegram.WebApp
  
  let waiting = false
  let ready = false
  let solving = false
  let timeout = null
  let label = 'Waiting'
  let sublabel = ''
  let color = ''
  
  let startTime = 0
  let time = 0

  Telegram.ready()
  const params = new URLSearchParams(window.location.search)
  const cube = params.get('cube')
  $: data = time + (cube === null ? '' : `:${cube}`)

  if (cube !== null) sublabel = `Selected cube: ${cube}`

  Telegram.MainButton.setParams({ is_visible: true })

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
      sublabel = ''
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

      Telegram.MainButton.setParams({ text: display })
      Telegram.sendData(data)
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
    <span id="sublabel">{sublabel}</span>
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
  height: 100vh;
}

#label {
  font-family: "Lucida Grande", "Lucida Sans Unicode", Arial, Helvetica, Verdana, sans-serif;
  font-size: 20vw;
  font-weight: bold;
  color: var(--tg-theme-text-color);
}

#sublabel {
  font-size: 5vw;
  color: var(--tg-theme-hint-color);
}
</style>
