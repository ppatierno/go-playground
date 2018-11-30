# Go playground

## Install on Fedora

Just use the `dnf` package manager running:

    sudo dnf install golang

Go binaries are installed in `/usr/lib/golang`.
Check the installation running:

    go version

Running `go env` you get the list of environment variables for Go like `GOROOT="/usr/lib/golang"`.

Running `whereis go` you get `/usr/bin/go` (which is on the `PATH`). This `go` file is a link to `etc/alternatives/go` which is then a link to the actual Go binary `/usr/lib/golang/bin/go`.
So you have : `/usr/bin/go` -> `etc/alternatives/go` -> `/usr/lib/golang/bin/go`

Another important environment variable is `GOPATH="/home/<user>/go"` (if `go` folder doesn't exist you have to create it or choose a different one). This is the path where source codes of your application live and where Go search for custom packages (not the system ones).
