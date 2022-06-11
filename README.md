# macli
Unofficial CLI-Based MyAnimeList Client

## How to use
Currently, macli is heavily under development so there are no pre-built binaries available.

### To compile macli, 

1. Clone this repo
2. Create a .env file with the following data:
```
ACCESS_TOKEN={your token here}
```
(a token can be generated easily with [this python script.](https://github.com/MikunoNaka/mal-authtoken-generator))

3. Run macli with `./macli`

NOTE: currently, it looks for the .env file in the PWD, but a token can also be passed with
``` shell
export ACCESS_TOKEN={your token here}
```
for use inside other directories. (this way of handling the token is going to change)

## Licence
Licenced under GNU General Public Licence V3

GNU GPL License: [LICENSE](LICENSE)

Copyright (c) 2022 Vidhu Kant Sharma
