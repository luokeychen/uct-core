package main

import (
	_ "net/http/pprof"
	"os"
	"time"
	"uct/common/conf"
	"uct/common/model"

	log "github.com/Sirupsen/logrus"
	"github.com/lib/pq"
	"github.com/pquerna/ffjson/ffjson"
	"gopkg.in/alecthomas/kingpin.v2"
	"uct/redis"
	"encoding/json"
	"uct/notification"
)

var (
	app           = kingpin.New("julia", "An application that queue messages from the database")
	configFile    = app.Flag("config", "configuration file for the application").Short('c').File()
	config        = conf.Config{}
)

var redisWrapper *redishelper.RedisWrapper

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Parse configuration file
	config = conf.OpenConfig(*configFile)
	config.AppName = app.Name

	redisWrapper = redishelper.New(config, app.Name)

	// Start profiling
	go model.StartPprof(config.GetDebugSever(app.Name))

	log.Debugln(config)

	// Open connection to postgresql
	listener := pq.NewListener(config.GetDbConfig(app.Name), 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.WithError(err).Fatalln()
		}
	})

	// Listen on channel
	if err := listener.Listen("status_events"); err != nil {
		panic(err)
	}

	postgresNotify := NewNotifier(listener)

	log.Infoln("Start monitoring PostgreSQL...")
	for {
		waitForNotification(postgresNotify, func(uctNotification *model.UCTNotification) {
			log.Debugln("pushing", uctNotification.TopicName)

			if b, err := uctNotification.Marshal(); err != nil {
				log.WithError(err).Fatalln()
			} else if _, err := redisWrapper.Client.Set(notification.MainQueueData + uctNotification.TopicName, b, 0).Result(); err != nil {
				log.WithError(err).Fatalln()
			} else if redisWrapper.RPush(notification.MainQueue, uctNotification.TopicName); err != nil {
				log.WithError(err).Fatalln()
			}
		})
	}
}

var sem = make(chan int, 10)

func waitForNotification(l Notifier, onNotify func (notification *model.UCTNotification)) {
	for {
		select {
		case message := <-l.Notify():
			if message == "" {
				continue
			}
			select {
			case sem <- 1:
				go func() {
					var notification model.UCTNotification
					err := ffjson.Unmarshal([]byte(message), &notification)
					model.CheckError(err)

					onNotify(&notification)
					// Process and send notification, free workRoutine when done.
					<- sem
				}()
			case <-time.After(time.Minute * 10):
				log.Fatalln("Routines blocked for too long")
			}
			// Received no notification from the database for 60 seconds.
		case <-time.After(1 * time.Minute):
			go l.Ping()
		}
	}
}

type pgNotifier struct {
	l *pq.Listener
	ch chan string
}

func NewNotifier(listener *pq.Listener) *pgNotifier {
	pg := &pgNotifier{l:listener, ch:make(chan string)}
	go func() {
		for n := range pg.l.Notify {
			pg.ch <- n.Extra
		}
	}()

	return pg
}

func (pg *pgNotifier) Notify() <-chan string {
	return pg.ch
}

func (pg *pgNotifier) Ping() {
	if err := pg.l.Ping(); err != nil {
		log.WithError(err).Fatalln("Failed to ping server")
	}
}

type MockNotifier struct {
	notifications []string
	ch chan string
}

func (pg *MockNotifier) send() {
	go func() {
		for _, val := range pg.notifications {
			fakeNoti := model.UCTNotification{TopicName:val}
			b, _ := json.Marshal(fakeNoti)
			pg.ch <- string(b)
		}
	}()
}

func (pg *MockNotifier) Notify() <-chan string {
	return pg.ch
}

func (pg *MockNotifier) Ping() {
	log.Debugln("Pinging notifier")
}

type Notifier interface {
	Notify() <-chan string
	Ping()
}

