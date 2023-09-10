package main

import (
	"net/http"
	"os"

	pb "github.com/maitaken/protobuf_sample/proto"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		logrus.Debugf("get request: [header][%v]", r.Header)
		person := pb.Person{
			Name:  "maitaken",
			Id:    1,
			Email: "hogepiyoapp@gmail.com",
			Phones: []*pb.Person_Number{
				{Number: "090-1234-5678", Type: pb.Person_WORK},
				{Number: "03-1234-5678", Type: pb.Person_HOME},
			},
		}
		p, err := proto.Marshal(&person)
		if err != nil {
			log.Errorf("failed to marshal: %v", err)
		}
		w.Write([]byte(p))
	})

	http.ListenAndServe(":8080", nil)
}
