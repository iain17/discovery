package discovery

import (
	"github.com/iain17/discovery/env"
	"github.com/iain17/discovery/pb"
	"github.com/iain17/freeport"
	"github.com/iain17/logger"
	"github.com/rs/xid"
	"github.com/golang/protobuf/proto"
	"io"
	"github.com/iain17/framed"
	"cirello.io/supervisor"
	"sync"
)

type LocalNode struct {
	Node
	discovery *Discovery
	ip        string //Gets filled in by stun service.
	port      int
	wg		  *sync.WaitGroup
	lastError error
	supervisor supervisor.Supervisor
	//Services
	listenerService ListenerService
	upNpService     UPnPService
	netTableService NetTableService
	StunService     StunService
	//Peer discoveries
	discoveryDHT  DiscoveryDHT
	discoveryIRC  DiscoveryIRC
	discoveryMDNS DiscoveryMDNS
}

func newLocalNode(discovery *Discovery) (*LocalNode, error) {
	i := &LocalNode{
		Node: Node{
			id:        xid.New().String(),
			logger: logger.New("LocalNode"),
			info:   map[string]string{},
		},
		discovery: discovery,
		port:      freeport.GetPortRange("udp", PORT_RANGE),
		wg: &sync.WaitGroup{},
	}
	i.supervisor.Log = func(i interface{}) {
		//logger.Debug(i)
	}
	i.listenerService.localNode = i
	i.supervisor.Add(&i.listenerService, supervisor.Permanent)
	i.netTableService.localNode = i
	i.supervisor.Add(&i.netTableService, supervisor.Transient)
	i.upNpService.localNode = i
	i.supervisor.Add(&i.upNpService, supervisor.Temporary)
	i.StunService.localNode = i
	i.supervisor.Add(&i.StunService, supervisor.Temporary)
	i.discoveryDHT.localNode = i
	i.supervisor.Add(&i.discoveryDHT, supervisor.Permanent)
	i.discoveryIRC.localNode = i
	i.supervisor.Add(&i.discoveryIRC, supervisor.Permanent)
	i.discoveryMDNS.localNode = i
	i.supervisor.Add(&i.discoveryMDNS, supervisor.Permanent)
	numServices := len(i.supervisor.Services())
	i.wg.Add(numServices)
	go i.supervisor.Serve(discovery.ctx)
	i.wg.Wait()
	i.wg = nil
	return i, i.lastError
}

func (ln *LocalNode) sendPeerInfo(w io.Writer) error {
	ln.infoMutex.Lock()
	defer ln.infoMutex.Unlock()
	peerInfo, err := proto.Marshal(&pb.Message{
		Version: env.VERSION,
		Msg: &pb.Message_PeerInfo{
			PeerInfo: &pb.DPeerInfo{
				Network: string(ln.discovery.network.ExportPublicKey()),
				Id:      ln.id,
				Info:    ln.info,
			},
		},
	})
	if err != nil {
		return err
	}
	return framed.Write(w, peerInfo)
}

func (ln *LocalNode) String() string {
	return "Local node."
}

//Will trigger updating the clients I'm connected to
func (ln *LocalNode) SetInfo(key string, value string) {
	//if ln.info[key] == value {
	//	return
	//}
	ln.info[key] = value
	for _, peer := range ln.netTableService.GetPeers() {
		go ln.sendPeerInfo(peer.conn)
	}
}
