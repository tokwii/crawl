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
glide install
```

```shell
# Compile the App
make build
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
