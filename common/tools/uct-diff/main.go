package main

import (
	"bufio"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"os"
	"uct/common/model"
)

var (
	app      = kingpin.New("model-diff", "An application to filter unchanged objects")
	format   = app.Flag("format", "choose file input format").Short('f').HintOptions(model.PROTOBUF, model.JSON).PlaceHolder("[protobuf, json]").Required().String()
	old      = app.Arg("old", "the first file to compare").Required().File()
	new      = app.Arg("new", "the second file to compare").File()
	logLevel = app.Flag("log-level", "Log level").Short('l').Default("debug").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if lvl, err := log.ParseLevel(*logLevel); err != nil {
		log.WithField("loglevel", *logLevel).Fatal(err)
	} else {
		log.SetLevel(lvl)
	}

	if *format != model.JSON && *format != model.PROTOBUF {
		log.Fatalln("Invalid format:", *format)
	}

	var firstFile = bufio.NewReader(*old)
	var secondFile *bufio.Reader

	if *new != nil {
		secondFile = bufio.NewReader(*new)
	} else {
		secondFile = bufio.NewReader(os.Stdin)
	}

	var oldUniversity model.University

	if err := model.UnmarshallMessage(*format, firstFile, &oldUniversity); err != nil {
		log.WithError(err).Fatalf("Failed to unmarshall message")
	}

	var newUniversity model.University

	if err := model.UnmarshallMessage(*format, secondFile, &newUniversity); err != nil {
		log.WithError(err).Fatalf("Failed to unmarshall message")
	}

	filteredUniversity := model.DiffAndFilter(oldUniversity, newUniversity)

	buf := model.MarshalMessage(*format, filteredUniversity)

	io.Copy(os.Stdout, buf)
}
