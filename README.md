## wcrawl

#### Build, Config and Run

```
# Clone to this repo structure
# golang >= go1.7.4

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
$ vi sitemap.xml
```
#### Test

```shell
# Run Tests
$ cd $GOPATH/src/github.com/tokwii/crawl
$ make test
```
#### Reference (Sitemap)
[1] https://www.sitemaps.org/protocol.html

[2] https://support.google.com/webmasters/answer/183668?hl=en&ref_topic=4581190
