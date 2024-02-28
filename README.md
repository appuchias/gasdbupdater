# WebDBUpdater

![Size](https://img.shields.io/github/repo-size/appuchias/webdbupdater?color=orange&style=flat-square)
[![Author](https://img.shields.io/badge/Project%20by-Appu-9cf?style=flat-square)](https://github.com/appuchias)

## How it works

It fetches the gas station prices from a Spanish government API
([API guide here](https://sedeaplicaciones.minetur.gob.es/ServiciosRESTCarburantes/PreciosCarburantes/help))
and then saves the current prices to a SQLite database (named `db.sqlite3`)

I made this to learn a bit of Go and try to improve the efficiency of the price updates of [my website](https://appu.ltd).

## Setup

1. Navigate to the desired folder to store the code: `cd <path>`
1. Clone the repo: `git clone https://github.com/appuchias/webdbupdater.git`
1. Navigate into the repo folder: `cd webdbupdater`
1. Compile the code: `go build`
1. Run it! (`./webdbupdater` on Unix-like systems)

## License

This code is licensed under the [GPLv3 license](https://github.com/appuchias/webdbupdater/blob/master/LICENSE).

Coded with 🖤 by Appu
