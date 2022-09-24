// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"log"

	"totally-legit-grow-management/v1/internal/logic"
	"totally-legit-grow-management/v1/resources/config"
)

func NewServer(cfg config.Config, shouldMigrate bool) *Server {

	// Device Handlers
	// Growing Group Handlers

	// GrowingLocation Handlers

	// Growing Medium Handlers

	// Grow Spot Handlers

	// Harvest Handlers

	// Inventory Handlers

	// Member Handlers

	// Nutrient Handlers

	// Organization Handlers

	// Plant Category Handlers

	// Plant Handlers

	// Plant Life Cycle Handlers

	// Plant Stage Handlers

	// Product Handlers

	// Seed Handlers

	// Task Handlers

	control, err := logic.NewLogicControl(cfg)
	if err != nil {
		log.Fatal("Logic Layer Not Setup!")
	}

	if shouldMigrate {
		err := control.Persistence.MigrateDBUP()
		if err != nil {
			// migrations couldn't happen
			log.Println(err)
		}
	}

	return &Server{}
}
