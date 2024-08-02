# PolarPrint

🖨️ Printer Manager for [SwitchCraft3](https://sc3.io). Store and queue printing jobs with ease! Licensed under [GNU AGPLv3](https://raw.githubusercontent.com/Snowflake-Software/PolarPrint/main/LICENSE).

> [!IMPORTANT]
> Snowflake-Software no longer has a precence on SwitchCraft3, as SwitchCraft3 is shutting down.
> As such this repository has been archived.

## Development

PolarPrint is written in Golang and handlebars.
It's recommended that you are familiar with Golang if you wish to work on the backend, and required that you have it installed.
See the install [offical install guide](https://go.dev/doc/install).

Once you are setup with golang, run the following to install the dependencies

```bash
go mod tidy
```

For easier development we recommend air to reload the program once you change something.

```bash
go install github.com/cosmtrek/air@latest
air
```

Behind the scenes it uses the awesome [Fiber](https://gofiber.io/) framework for the web server.

## API Documentation

The API is documented with the awesome [Bruno](https://www.usebruno.com/), a "Fast and Git-Friendly Opensource API client".
