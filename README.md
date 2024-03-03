# GasDBUpdater

![Size](https://img.shields.io/github/repo-size/appuchias/gasdbupdater?color=orange&style=flat-square)
[![Author](https://img.shields.io/badge/Project%20by-Appu-9cf?style=flat-square)](https://github.com/appuchias)

## How it works

It fetches the gas station prices from a Spanish government API
([API guide here](https://sedeaplicaciones.minetur.gob.es/ServiciosRESTCarburantes/PreciosCarburantes/help))
and then saves the current prices to a SQLite database (named `db.sqlite3`)

I made this to learn a bit of Go and try to improve the efficiency of the price updates of [my website](https://appu.ltd).

Current version should work but I'm not currently using it.

## Setup

1. Navigate to the desired folder to store the code: `cd <path>`
1. Clone the repo: `git clone https://github.com/appuchias/gasdbupdater.git`
1. Navigate into the repo folder: `cd gasdbupdater`
1. Compile the code: `go build`
1. Run it! (`./gasdbupdater` on Unix-like systems)

## License

This code is licensed under the [GPLv3 license](https://github.com/appuchias/gasdbupdater/blob/master/LICENSE).

Coded with 🖤 by Appu
