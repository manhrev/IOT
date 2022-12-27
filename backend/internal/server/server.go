package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	hook "github.com/manhrev/IOT/backend/internal/hooks"
	"github.com/mochi-co/mqtt/v2"
	"github.com/mochi-co/mqtt/v2/hooks/auth"
	"github.com/mochi-co/mqtt/v2/listeners"
)

func Run() {
	server := newServer()
	Serve(server)
}

func newServer() *mqtt.Server {
	server := mqtt.New(nil)
	return server
}

func Serve(server *mqtt.Server) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	err := server.AddHook(new(auth.AllowHook), nil)
	if err != nil {
		log.Fatalf("error adding auth hook: %v", err)
	}

	tcp := listeners.NewTCP("t1", ":1883", nil)
	err = server.AddListener(tcp)
	if err != nil {
		log.Fatalf("error adding listener: %v", err)
	}

	err = server.AddHook(new(hook.CustomHook), map[string]any{})
	if err != nil {
		log.Fatalf("error adding custom hook: %v", err)
	}

	// Start the server
	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Demonstration of directly publishing messages to a topic via the
	// `server.Publish` method. Subscribe to `direct/publish` using your
	// MQTT client to see the messages.
	// go func() {
	// 	cl := mqtt.NewInlineClient("inline", "local")
	// 	for range time.Tick(time.Second * 10) {
	// 		server.InjectPacket(cl, packets.Packet{
	// 			FixedHeader: packets.FixedHeader{
	// 				Type: packets.Publish,
	// 			},
	// 			TopicName: "direct/publish",
	// 			Payload:   []byte("scheduled message"),
	// 		})
	// 		server.Log.Info().Msgf("main.go issued direct message to direct/publish")
	// 	}
	// }()

	<-done
	server.Log.Warn().Msg("caught signal, stopping...")
	server.Close()
	server.Log.Info().Msg("main.go finished")
}
