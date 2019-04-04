package airms_example

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	log "airman.com/airms/pkg/mslog"
	app "airman.com/airms/pkg/server"

	pb "airman.com/airmsExample/node/apis"
	"airman.com/airmsExample/node/config"
)

type AirmsExampleService struct {
	grace *app.Grace
	name  string
}

func NewAirmsExampleService(name string, deadline time.Duration) *AirmsExampleService {
	return &AirmsExampleService{
		name:  name,
		grace: app.NewGrace(deadline),
	}
}

func (s *AirmsExampleService) Run() error {
	sc := config.GetService()
	srv := app.NewGRPCServer(sc.Name, sc.Address)
	if srv == nil {
		log.Errorf("new server failed, %v:%v", sc.Name, sc.Address)
		return errors.New("new server failed")
	}
	s.grace.SetServer(srv)

	sr := config.GetRegister()
	if len(sr.Addresses) == 0 {
		log.Error("register server is nil")
		return errors.New("register server is nil")
	}

	if err := s.grace.SetRegisterWithTTL(sr.Addresses, sc.Name, sr.TTL); err != nil {
		log.Errorf("gRPC server register failed, %v:%v", sr.Addresses, err)
		return err
	}

	tr := config.GetTrace()
	gs := s.grace.GetGRPCWithTracer(tr.Url)
	if gs == nil {
		log.Errorf("gRPC server tracer error")
		return errors.New("tracer is nil")
	}
	pb.RegisterAirmsExampleServer(gs, s)
	s.grace.SetGRPCServer(gs)
	return s.grace.Run()
}

func (s *AirmsExampleService) Stop() error {
	return s.grace.StopServer()
}

func (s *AirmsExampleService) Name() string {
	return s.name
}

// SayHello
func (s *AirmsExampleService) SayHello(ctx context.Context, req *pb.AirmsExampleRequest) (*pb.AirmsExampleReply, error) {

	rrNum := strings.Split(req.Name, " ")[1]
	if n, _ := strconv.ParseInt(rrNum, 10, 64); n%6 == 0 {
		//trace.Error(errors.New("test for trace and handle time"))
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	log.Infof("SayHello incoming:%v", req.Name)
	return &pb.AirmsExampleReply{Message: "Hello " + req.Name}, nil
}
