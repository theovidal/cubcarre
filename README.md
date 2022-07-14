<div align="center">
    <img src="webapp/public/cubcarre.png" alt="CubCarré logo" width="30%">
    <h1>CubCarré</h1>
    <h3>Rubik's cube timer on a Telegram bot</h3>
    <a href="https://t.me/cubcarre_bot">Add the bot</a> — <a href="./LICENSE">License</a>
</div>

## 💻 Development

Thanks for participating in CubCarré's improvement and/or debugging! First, check the following requirements:

- Git, for version control
- Golang 1.18 or higher with go-modules for dependencies
- Node.js (tested on v14+) with npm or yarn for dependencies
- A SQLite database file (whatever location)

Clone the project on your local machine:

```bash
git clone https://github.com/theovidal/cubcarre  # HTTP
git clone git@github.com:theovidal/cubcarre      # SSH
```

### Telegram bot

The bot itself is a Go program that uses the [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) library.

- Database models are located under the `database` directory
- Commands and callback actions are in `handlers`
- Some helpers can be found in `lib`

Set up some environment variables described in the [.env.example file](./.env.example), either by adding them in the shell or by creating a .env file at the root of the project. To run and test the bot, simply use `go run .` in the working directory. To build an executable, use `go build .`.

### Web app

The web app associated with the bot is located under the [webapp folder](./webapp). It is written with [Svelte](https://svelte.dev) and built with [Vite](https://vitejs.dev). This combination ensures a light and instant experience to the user.

To start, go in the folder and install required dependencies :

```bash
npm i          # Install the dependencies
yarn i
```

While developing, you can run a local test server that you plug to the Telegram bot. You can either open your ports on your router, or use services such as [ngrok](https://ngrok.com/) that creates a tunnel to your machine.

```bash
npm run dev    # Run the development server
yarn dev
```

Once done, to build for production, use the script below and serve the output directory through a web server, like Apache or nginx. You can also deploy on GitHub Pages, Netlify....

```bash
npm run build  # Build for production
yarn build
```

## 📜 Credits

- Maintainer: [Théo Vidal](https://github.com/theovidal)
- Libraries: [check go.mod](./go.mod)

## 🔐 License

[GNU GPL v3](./LICENSE)
