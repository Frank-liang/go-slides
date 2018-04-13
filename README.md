# Usage

You can view the slides online or in local:

### Online

Use http://talks.godoc.org/github.com/wusuopubupt/go-slides/xx/*.slide 

eg. [http://talks.godoc.org/github.com/wusuopubupt/go-slides/goshare/go.slide](http://talks.godoc.org/github.com/wusuopubupt/go-slides/goshare/go.slide) to view

### Local

``` bash
# 1) install present
go get -u golang.org/x/tools/cmd/present 
go install golang.org/x/tools/cmd/present

# 2) get the slides
go get github.com/wusuopubupt/go-slides

# 3) start local present http server
cd ${GOPATH}/src/github.com/wusuopubupt/go-slides && ${GOPATH}/bin/present -http 127.0.0.1:8989

# 4) view the slides by accessing 127.0.0.1:8989 in your browser


```

Most of the slides are created by icexin, thanks for icexin's great work!

Powered by [present](https://godoc.org/golang.org/x/tools/present)

