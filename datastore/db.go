package datastore

import (
	"context"
	"log"
	"sync"

	"github.com/goava/di"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"blog/config"
)

func Start(ctx context.Context, cfg *config.Config, dic *di.Container) error {
	log.Println("Start connect to database...")
	orm, err := mySQLConnect(cfg)
	if err != nil {
		return errors.Wrap(err, "database connect filed")
	}

	if err := dic.Provide(func() *gorm.DB { return orm }); err != nil {
		return errors.Wrap(err, "provide orm filed")
	}
	var wg *sync.WaitGroup
	if err := dic.Resolve(&wg); err != nil {
		return errors.Wrap(err, "resolve application wait group filed")
	}
	createContextHandler(ctx, orm, wg)
	autoMigrate(orm)

	log.Println("Connect to database success.")
	return nil
}

func Close(orm *gorm.DB) error {
	if orm != nil {
		db, err := orm.DB()
		if err != nil {
			return err
		}

		return db.Close()
	}

	return nil
}

func createContextHandler(ctx context.Context, db *gorm.DB, wg *sync.WaitGroup) {
	cc, cancel := context.WithCancel(ctx)
	wg.Add(1)
	go func() {
		for {
			<-cc.Done()
			log.Println("Stop database service...")
			if err := Close(db); err != nil {
				log.Println("Close database connection: ", err)
			} else {
				log.Println("Database Service stopped.")
			}
			cancel()
			wg.Done()
			return
		}
	}()

}
