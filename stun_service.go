package discovery

import (
	"context"
	"github.com/ccding/go-stun/stun"
	"time"
	"fmt"
	"github.com/iain17/logger"
)

type StunService struct {
	client *stun.Client
	localNode *LocalNode
	logger  *logger.Logger
	context context.Context
}

func (d *StunService) String() string {
	return "Stun"
}

func (s *StunService) init(ctx context.Context) error {
	defer func() {
		if s.localNode.wg != nil {
			s.localNode.wg.Done()
		}
	}()
	s.logger = logger.New(s.String())
	s.context = ctx
	s.client = stun.NewClientWithConnection(s.localNode.listenerService.socket)
	return nil
}

func (s *StunService) Serve(ctx context.Context) {
	defer s.Stop()
	if err := s.init(ctx); err != nil {
		s.localNode.lastError = err
		panic(err)
	}
	s.localNode.waitTilReady()
	for {
		select {
		case <-s.context.Done():
			return
		default:
			err := s.process()
			if err != nil {
				s.logger.Debugf("error on forwarding process, %v", err)
			}
			time.Sleep(time.Minute)
		}
	}
}

func (s *StunService) Stop() {

}

func (s *StunService) process() (err error) {
	nat, host, err := s.client.Discover()
	if err != nil {
		return err
	}

	if host != nil {
		s.logger.Debugf("processed, family %d, host %q, port %d", host.Family(), host.IP(), host.Port())
		s.localNode.ip = host.IP()
	}

	switch nat {
	case stun.NATError:
		return fmt.Errorf("test failed")
	case stun.NATUnknown:
		return fmt.Errorf("unexpected response from the STUN server")
	case stun.NATBlocked:
		return fmt.Errorf("UDP is blocked")
	case stun.NATFull:
		return fmt.Errorf("full cone NAT")
	case stun.NATSymetric:
		return fmt.Errorf("symetric NAT")
	case stun.NATRestricted:
		return fmt.Errorf("restricted NAT")
	case stun.NATPortRestricted:
		return fmt.Errorf("port restricted NAT")
	case stun.NATNone:
		return fmt.Errorf("not behind a NAT")
	case stun.NATSymetricUDPFirewall:
		return fmt.Errorf("symetric UDP firewall")
	}
	s.logger.Info("NAT open!")
	return nil
}

