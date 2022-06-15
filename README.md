# macli
Unofficial CLI-Based MyAnimeList Client

### Notice:
macli is currently highly under development and only 3 commands work as of now.
It can search for anime, update episodes and set status.

## How to install
macli is available on the AUR. But I can't guarantee if it works right now.

To compile macli, simply clone this repo, and install the dependencies `go` and `gnome-keyring` (linux only)

Then, run `go build` to build the binary, and move the macli executable to your path for global usage.

## Logging in
Currently there is no support to generate an access token,

but it can be generated with [this python script.](https://github.com/MikunoNaka/mal-authtoken-generator)

simply generate one using my script, and then run `macli login` then paste the token.

## Licence
Licenced under GNU General Public Licence V3

GNU GPL License: [LICENSE](LICENSE)

Copyright (c) 2022 Vidhu Kant Sharma
