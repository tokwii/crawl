## wcrawl

#### Build, Config and Run

```
# Clone to this repo structure

$GOPATH/
└── src
    └── github.com
        └── tokwii
            └── crawl
```

```shell
# Install dependencies
$ cd $GOPATH/src/github.com/tokwii/crawl
$ glide install
```

```shell
# Compile the App
$ cd $GOPATH/src/github.com/tokwii/crawl
$ make build
```

```shell
# Config the App
$ cd $GOPATH/src/github.com/tokwii/crawl
$ vi config/settings.toml
```

```shell
# Run App
$ cd $GOPATH/src/github.com/tokwii/crawl
$ ./crawl 
```
#### Test

```shell
# Run Tests
$ cd $GOPATH/src/github.com/tokwii/crawl
$ make test
```
