package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"toast.cafe/x/webdd/middleware"
	"toast.cafe/x/webdd/util"
)

func main() {
	// ---- Flags
	wd, _ := os.Getwd()
	flag.BoolVar(&util.S.Asciidoc, "asciidoc", true, "preparse asciidoc files")
	flag.StringVar(&util.S.Dir, "directory", wd, "directory to serve")
	flag.BoolVar(&util.S.Log, "log", true, "log all requests")
	flag.StringVar(&util.S.LogLevel, "loglevel", "info", "log level to display")
	flag.BoolVar(&util.S.Markdown, "markdown", true, "prepare markdown files")
	flag.BoolVar(&util.S.Safe, "safe", false, "re-sanitize requests")
	flag.Parse()

	// Setup Logging
	log := logrus.New()
	loglevel, err := logrus.ParseLevel(util.S.LogLevel)
	if err != nil {
		loglevel = logrus.InfoLevel
		log.Warning("could not parse log level, defaulting to info")
	}
	log.SetLevel(loglevel)

	// ---- Report Flags
	log.WithFields(logrus.Fields{
		"asciidoc":  util.S.Asciidoc,
		"directory": util.S.Dir,
		"log":       util.S.Log,
		"loglevel":  util.S.LogLevel,
		"markdown":  util.S.Markdown,
		"safe":      util.S.Safe,
	}).Trace("compiled settings")

	// Prepare Middlewares
	var md []middleware.Middleware
	if util.S.Safe {
		md = append(md, middleware.Safe)
	}
	if util.S.Log {
		md = append(md, middleware.HTTPLog)
	}
	if util.S.Asciidoc {
		md = append(md, middleware.ASCIIDoc)
	}
	if util.S.Markdown {
		md = append(md, middleware.MarkDown)
	}
	md = append(md, middleware.File)
	handler := middleware.HandleMiddlewares(md)

	// ---- Launch Server
	s := &fasthttp.Server{
		Handler: handler,
		Logger:  log,
	}
	s.ListenAndServe(":8080")
}
