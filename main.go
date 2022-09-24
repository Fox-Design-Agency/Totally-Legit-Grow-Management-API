// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"

	"totally-legit-grow-management/v1/app"
	"totally-legit-grow-management/v1/internal/server"
	"totally-legit-grow-management/v1/resources/config"
)

func main() {
	/**********************************************************************
	/
	/	Configuration
	/
	/**********************************************************************/
	ctx, cancel := context.WithCancel(context.Background())

	cfg := config.LoadConfig()

	svr := server.NewServer(cfg, true)

	app.Run(ctx, &cfg, svr)

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}
