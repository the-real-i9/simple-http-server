# simple-http-server

`simple-http-server` is a tool, built in Go, that allows you serve any directory in your local file system over the network.

## Setup

> I'm in Linux, so if you use a different os/arch, ensure you perform the equivalent action.

1. Clone this Repo
2. `go build` the source (in your os/arch)

    ```bash
    go build -o $GOBIN/my-http-server
    ```

3. Place the binary in your system's `/bin` directory

    ```bash
    sudo cp $GOBIN/my-http-server /bin
    ```

## Usage

### Without flags

Execute the binary in the directory you want to serve.

```bash
user@pc:~/website$ my-http-server
```

By default, it serves the current directory where it is run (i.e. `.`) as root, and it runs on port `5442`

### With flags

The two flags available are `-p`, which you can use to specify a custom port, and `-d`, which you can use to specify the directory you want to serve.

#### Specify a Custom Port

```bash
user@pc:~/website$ my-http-server -p 6544
```

#### Specify Directory

```bash
user@pc:~/website$ my-http-server -d ./about
```

After running the application you get the message:

``` bash
Live @ http://localhost:5442
```
