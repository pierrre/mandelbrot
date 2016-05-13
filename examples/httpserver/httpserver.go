package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"image"
	"image/color"
	"net/http"
	"os"
	"time"

	"github.com/disintegration/gift"
	"github.com/pierrre/githubhook"
	"github.com/pierrre/imageserver"
	imageserver_cache "github.com/pierrre/imageserver/cache"
	imageserver_cache_memory "github.com/pierrre/imageserver/cache/memory"
	imageserver_http "github.com/pierrre/imageserver/http"
	imageserver_image "github.com/pierrre/imageserver/image"
	imageserver_image_gamma "github.com/pierrre/imageserver/image/gamma"
	_ "github.com/pierrre/imageserver/image/png"
	mandelbrot_image "github.com/pierrre/mandelbrot/image"
	mandelbrot_image_colorizer_rainbow "github.com/pierrre/mandelbrot/image/colorizer/rainbow"
)

var (
	flagHTTPAddr            = ":8080"
	flagGitHubWebhookSecret string
	flagCache               = int64(64 * (1 << 20))
	flagQuality             = uint(0)
	flagMaxIter             = uint(1000)
)

func main() {
	parseFlags()
	startHTTPServer()
}

func parseFlags() {
	flag.StringVar(&flagHTTPAddr, "http", flagHTTPAddr, "HTTP addr")
	flag.StringVar(&flagGitHubWebhookSecret, "github-webhook-secret", flagGitHubWebhookSecret, "GitHub webhook secret")
	flag.Int64Var(&flagCache, "cache", flagCache, "Cache")
	flag.UintVar(&flagQuality, "quality", flagQuality, "Quality")
	flag.UintVar(&flagMaxIter, "max-iter", flagMaxIter, "Max iter")
	flag.Parse()
}

func startHTTPServer() {
	http.HandleFunc("/", rootHTTPHandler)
	http.Handle("/i", newImageHTTPHandler())
	if h := newGitHubWebhookHTTPHandler(); h != nil {
		http.Handle("/github_webhook", h)
	}
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(flagHTTPAddr, nil)
	if err != nil {
		panic(err)
	}
}

func rootHTTPHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(
		`<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width">
	<title>Mandelbrot</title>
	<style type="text/css">
		html, body {
			height: 100%;
			margin: 0;
		}
		body {
			background-color: #000;
		}
		#map {
			min-height: 100%;
		}
	</style>
	<link rel="stylesheet" type="text/css" href="//cdnjs.cloudflare.com/ajax/libs/ol3/3.15.1/ol.css">
	<script type="text/javascript" src="//cdnjs.cloudflare.com/ajax/libs/ol3/3.15.1/ol.js"></script>
</head>
<body>
	<div id="map"></div>

	<script type="text/javascript">
		var extent = [-2, -2, 2, 2];
		var projection = new ol.proj.Projection({
			extent: extent
		});
		var map = new ol.Map({
			target: 'map',
			layers: [
				new ol.layer.Tile({
					source: new ol.source.XYZ({
						url: '/i?x={x}&y={y}&z={z}',
						projection: projection
					})
				})
			],
			view: new ol.View({
				projection: projection,
				center: ol.extent.getCenter(extent),
				zoom: 0,
				maxZoom: 42
			})
		});
	</script>
</body>
</html>`))
}

const (
	tileSize = 256
	maxTileZ = 42
	radius   = 2
)

func newImageHTTPHandler() http.Handler {
	var hdr http.Handler
	hdr = &imageserver_http.Handler{
		Parser:   &mandelbrotHTTPParser{},
		Server:   newImageServer(),
		ETagFunc: imageserver_http.NewParamsHashETagFunc(sha256.New),
	}
	hdr = &imageserver_http.ExpiresHandler{
		Handler: hdr,
		Expires: 7 * 24 * time.Hour,
	}
	hdr = &imageserver_http.CacheControlPublicHandler{
		Handler: hdr,
	}
	return hdr
}

type mandelbrotHTTPParser struct{}

func (prs *mandelbrotHTTPParser) Parse(req *http.Request, params imageserver.Params) error {
	for _, k := range []string{"x", "y", "z"} {
		err := imageserver_http.ParseQueryInt(k, req, params)
		if err != nil {
			return err
		}
	}
	return nil
}

func (prs *mandelbrotHTTPParser) Resolve(param string) string {
	if param == "x" || param == "y" || param == "z" {
		return param
	}
	return ""
}

func newImageServer() imageserver.Server {
	var srv imageserver.Server
	srv = &imageserver_image.Server{
		Provider:      newImageProvider(),
		DefaultFormat: "png",
	}
	srv = newImageCacheServer(srv)
	return srv
}

func newImageCacheServer(srv imageserver.Server) imageserver.Server {
	if flagCache <= 0 {
		return srv
	}
	return &imageserver_cache.Server{
		Server:       srv,
		Cache:        imageserver_cache_memory.New(flagCache),
		KeyGenerator: imageserver_cache.NewParamsHashKeyGenerator(sha256.New),
	}
}

func newImageProvider() imageserver_image.Provider {
	tileRenderSize := tileSize << flagQuality
	clr := mandelbrot_image.BoundColorizer(
		mandelbrot_image.ColorColorizer(color.Black),
		mandelbrot_image_colorizer_rainbow.Colorizer(16, 0),
	)
	var prv imageserver_image.Provider
	prv = imageserver_image.ProviderFunc(func(params imageserver.Params) (image.Image, error) {
		tsf, err := getTransformation(params, tileRenderSize)
		if err != nil {
			return nil, err
		}
		im := image.NewNRGBA(image.Rect(0, 0, tileRenderSize, tileRenderSize))
		mandelbrot_image.Render(im, tsf, int(flagMaxIter), clr)
		return im, nil
	})
	if tileRenderSize != tileSize {
		prv = &imageserver_image.ProcessorProvider{
			Provider: prv,
			Processor: imageserver_image_gamma.NewCorrectionProcessor(
				imageserver_image.ProcessorFunc(func(im image.Image, params imageserver.Params) (image.Image, error) {
					g := gift.New(gift.Resize(tileSize, tileSize, gift.CubicResampling))
					dst := image.NewNRGBA64(g.Bounds(im.Bounds()))
					g.Draw(dst, im)
					return dst, nil
				}),
				true,
			),
		}
	}
	return prv
}

func getTransformation(params imageserver.Params, tileRenderSize int) (mandelbrot_image.Transformation, error) {
	tileX, tileY, tileZ, err := getTileXYZParam(params)
	if err != nil {
		return nil, err
	}
	tileCount := 1 << uint(tileZ)
	halfPix := tileCount * tileRenderSize / 2
	tilePixOff := complex(float64(tileX*tileRenderSize-halfPix), float64(tileY*tileRenderSize-halfPix))
	invScale := float64(radius) / float64(halfPix)
	return func(c complex128) complex128 {
		c += tilePixOff
		c = complex(real(c)*invScale, -imag(c)*invScale)
		return c
	}, nil
}

func getTileXYZParam(params imageserver.Params) (tileX, tileY, tileZ int, err error) {
	tileZ, err = getTileParam(params, "z", 0, maxTileZ)
	if err != nil {
		return
	}
	maxTileXY := (1 << uint(tileZ)) - 1
	tileX, err = getTileParam(params, "x", 0, maxTileXY)
	if err != nil {
		return
	}
	tileY, err = getTileParam(params, "y", 0, maxTileXY)
	if err != nil {
		return
	}
	return
}

func getTileParam(params imageserver.Params, name string, min, max int) (int, error) {
	tile, err := params.GetInt(name)
	if err != nil {
		return 0, err
	}
	if tile < min || tile > max {
		return 0, &imageserver.ParamError{Param: name, Message: fmt.Sprintf("must be between %d and %d", min, max)}
	}
	return tile, nil
}

func newGitHubWebhookHTTPHandler() http.Handler {
	if flagGitHubWebhookSecret == "" {
		return nil
	}
	return &githubhook.Handler{
		Secret: flagGitHubWebhookSecret,
		Delivery: func(event string, deliveryID string, payload interface{}) {
			if event == "push" {
				time.AfterFunc(5*time.Second, func() {
					os.Exit(0)
				})
			}
		},
	}
}
