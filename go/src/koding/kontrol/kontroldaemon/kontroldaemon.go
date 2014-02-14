package main

import (
	"flag"
	"koding/db/mongodb"
	"koding/db/mongodb/modelhelper"
	"koding/kontrol/kontroldaemon/handler"
	"koding/kontrol/kontrolhelper"
	"koding/tools/config"
	"koding/tools/logger"

	"github.com/streadway/amqp"
)

var (
	mongo       *mongodb.MongoDB
	flagProfile = flag.String("c", "", "Configuration profile from file")
	flagDebug   = flag.Bool("d", false, "Debug mode")
	log         = logger.New("kontroldaemon")
)

func main() {
	flag.Parse()
	if *flagProfile == "" {
		log.Fatal("Please define config file with -c")
	}

	conf := config.MustConfig(*flagProfile)

	var logLevel logger.Level
	if *flagDebug {
		logLevel = logger.DEBUG
	} else {
		logLevel = logger.GetLoggingLevelFromConfig("kontroldaemon", *flagProfile)
	}
	log.SetLevel(logLevel)

	mongo = mongodb.NewMongoDB(conf.Mongo)
	modelhelper.Initialize(conf.Mongo)

	handler.Startup(conf)
	startRouting(conf)
}

func startRouting(conf *config.Config) {
	type bind struct {
		name     string
		queue    string
		key      string
		exchange string
		kind     string
	}

	streams := make(map[string]<-chan amqp.Delivery)
	bindings := []bind{
		bind{"worker", "kontrol-worker", "input.worker", "workerExchange", "topic"},
	}

	// remove old queue before we use our new auto-delete queue
	deleteConn := kontrolhelper.CreateAmqpConnection(conf)
	deleteChan := kontrolhelper.CreateChannel(deleteConn)
	_, err := deleteChan.QueueDelete("kontrol-worker", false, false, true)
	if err != nil {
		log.Info("queue deleting err:", err.Error)
	}
	deleteConn.Close()

	connection := kontrolhelper.CreateAmqpConnection(conf)
	channel := kontrolhelper.CreateChannel(connection)

	for _, b := range bindings {
		streams[b.name] = kontrolhelper.CreateStream(channel, b.kind, b.exchange, b.queue, b.key, true, false)
	}

	log.Info("kontroldaemon routing started")

	for d := range streams["worker"] {
		go handler.WorkerMessage(d.Body)
	}
}
