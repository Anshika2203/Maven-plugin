// Copyright 2024 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package main

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"maven-plugin/plugin"
)

func main() {
	logrus.SetFormatter(new(formatter))

	var args plugin.Plugin
	if err := envconfig.Process("", &args); err != nil {
		logrus.Fatalln(err)
	}

	switch args.LogLevel {
	case "debug":
		logrus.SetFormatter(textFormatter)
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetFormatter(textFormatter)
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	if err := args.Exec(context.Background()); err != nil {
		logrus.Fatalln(err)
	}
}

type formatter struct{}

func (*formatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

var textFormatter = &logrus.TextFormatter{
	DisableTimestamp: true,
}
