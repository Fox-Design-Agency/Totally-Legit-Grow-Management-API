// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logic

import (
	"log"
	devicecontrol "totally-legit-grow-management/v1/internal/logic/deviceControl"
	"totally-legit-grow-management/v1/internal/persistence"
	"totally-legit-grow-management/v1/resources/config"
)

func NewLogicControl(cfg config.Config) (*Logic, error) {
	db, err := persistence.NewPersistence(cfg.Database.Dialect(), cfg.Database.Connection())
	if err != nil {
		log.Fatal("Couldn't Make DB Connection")
	}
	// device logic
	deviceLogic := devicecontrol.InitDeviceLogic(db)

	// growing groups

	//  growing levels

	// growing locations

	// growing medium
	// grow spot plants

	// grow spots

	// harvests

	// inventory

	// memebrs

	// nutrients

	// organizations

	// plant categories

	// plant

	// plant life cycle

	// plant stages

	// products

	// seeds

	//task

	return &Logic{
		DeviceLogic: deviceLogic,
		Persistence: db,
	}, nil
}
