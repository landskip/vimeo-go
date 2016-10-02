[![CircleCI](https://circleci.com/gh/landskip/vimeo-go/tree/master.svg?style=shield)](https://circleci.com/gh/landskip/vimeo-go/tree/master)

# Go Vimeo 

Golang library for [vimeo][0] APIs.


## Prepared

1. Create your [new vimeo app][1].
1. Generate an access token
2. Set VIMEO_TOKEN in your environment

## Test

Run `go test` command.

## Usages

### Get own videos

```
vimeo.SetTokenFromEnv()
cli := vimeo.VideoClient{}

options := vimeo.ListOptions{
    Page: 2,
    Sort: vimeo.SortComments,
}
videos, err := cli.ListMyVideos(options)

fmt.Println("length: " + videos.Total)
```

[0]: https://developer.vimeo.com/
[1]: https://developer.vimeo.com/apps/new
