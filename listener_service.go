package discovery

import (
	"github.com/anacrolix/utp"
	"net"
	"context"
	"github.com/iain17/discovery/pb"
	"github.com/iain17/logger"
	"errors"
	"fmt"
)

type ListenerService struct {
	localNode *LocalNode
	listener net.PacketConn
	context context.Context
	socket    *utp.Socket

	logger *logger.Logger
}

func (d *ListenerService) String() string {
	return "listener"
}

func (l *ListenerService) init(ctx context.Context) error {
	defer func() {
		if l.localNode.wg != nil {
			l.localNode.wg.Done()
		}
	}()
	l.logger = logger.New(l.String())
	l.context = ctx

	//Open a new socket on a free UDP port.
	var err error
	l.listener, err = net.ListenPacket("udp4", ":0")
	if err != nil {
		return fmt.Errorf("could not listen: %s", err.Error())
	}
	l.socket, err = utp.NewSocketFromPacketConn(l.listener)
	if err == nil {
		addr := l.socket.Addr().(*net.UDPAddr)
		l.localNode.port = addr.Port
		l.logger.Infof("listening on %d", l.localNode.port)
	}
	return err
}

func (l *ListenerService) Serve(ctx context.Context) {
	defer l.Stop()
	if err := l.init(ctx); err != nil {
		l.localNode.lastError = err
		panic(err)
	}

	for {
		select {
		case <-l.context.Done():
			return
		default:
			l.logger.Info("Listing for new connections")
			conn, err := l.socket.Accept()
			if err != nil {
				logger.Warning(err)
				break
			}

			l.logger.Debugf("new connection from %q", conn.RemoteAddr().String())

			if err = l.process(conn); err != nil {
				conn.Close()
				if err.Error() == "peer reset" || err.Error() == "we can't add ourselves" {
					continue
				}
				l.logger.Errorf("error on process, %v", err)
			}
		}
	}
}

func (l *ListenerService) Stop() {
	l.logger.Info("Stopping...")
	l.socket.Close()
}

//We receive a connection from a possible new peer.
func (l *ListenerService) process(c net.Conn) error {
	rn := NewRemoteNode(c, l.localNode)

	rn.logger.Debug("Waiting for peer info...")
	peerInfo, err := pb.DecodePeerInfo(c, string(l.localNode.discovery.network.ExportPublicKey()))
	if err != nil {
		return err
	}
	if peerInfo.Id == l.localNode.id {
		return errors.New("we can't add ourselves")
	}
	if l.localNode.netTableService.isConnected(peerInfo.Id) {
		logger.Debugf("We are already connected to %s", peerInfo.Id)
		return nil
	}
	rn.logger.Debug("Received peer info...")

	rn.logger.Debug("Sending our peer info...")
	err = l.localNode.sendPeerInfo(c)
	if err != nil {
		return err
	}
	rn.logger.Debug("Sent our peer info...")

	rn.Initialize(peerInfo)
	l.localNode.netTableService.AddRemoteNode(rn)
	return nil
}
