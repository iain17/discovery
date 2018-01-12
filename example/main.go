package main

import (
	"github.com/iain17/discovery"
	"github.com/iain17/discovery/network"
	"github.com/iain17/logger"
	"time"
	"os"
	"context"
)

const testNetwork = "2d2d2d2d2d424547494e205253412050524956415445204b45592d2d2d2d2d0a4d4949456f77494241414b4341514541356d475a5a546e555647344d4d497532464e732f6d4a4d7a44526c3974487a7652443273544a71624d695a4c4363515a0a6953495879702f4e6e2b336457565a666b66616e687436616a385a633230466b4f2f6c3563316f4178305233304c625072697456514e4249674e456a41714b4b0a4b5967616f41556c556a2f4f324756447a56463564766273764d745a506b6873426c73492b46455553734d43326177326833545975596c7341447a4b366c38460a4b507658486c586132567463572b476174332b663042494d6c4576704171396a4958372b4546566b4b497247472b437661574d566f4f6c53425272594e6278720a504836596b5357424f6a6776635138397061705265592f396178554f4a506572337777306e6e4f6f6e4a456e534e6c4650585555716743434a51564234646e6b0a3372306274414b56384a63767544756862723948676b7072353447614a7a6972723545374751494441514142416f494241436c67344564524972546a617963520a414f777745505562677a4459496a375872625870436468636349474b5a54486b3264314b6c564663634864494a746954376568756a39706d67775a4e7a3448460a396658627368656f30376a327344703569305779484d6265596864592b4e7652533956337a3668734a43376f72514d7139516d636c353970624533676a432f6d0a6e42687349326c4f5473473630766b3775444f4f4a6872303973503330415649375042467645354b39716c58702b4a32416b553544484364374174766b6267300a6a4664557850474f6a79436f6d2b6c514c63323739574b5164647654794f6333584b392b3459646650395a6568376478705875427973433234726e556f70476d0a3232476e676e48624c7433796b6e66625853394e6e712f356c712b4834784674687847776a55776d6275626f72416b64483265335a526e307a385939395531420a463166326c2f3043675945412b76645845674263463959316338587453555279572b6a584b735630542f505858754c7336746e7469564c656d3173334c3377790a50662f30742f5267536f493477644b445a4c49696a6d52307a3772316c646741732f6248516e70756c616d6852694173395a6952655659702f756a46764358560a38346d6c2b3557614f55764d544a41464a33335577354459536c6841644969797a75417761384b6857735976416b702b72525a544359634367594541367743500a53506b797171525663716b2f5944722b31586d5354593341504c487a426b2f6e6d7171574366686b495a4c374f2b394d327759572b6e6669787042377a3432610a7738772b71447832624d66546b703442374f565777525538464c6f42714c48677548754771315532344e703049683541382f37304c64757945444543474975500a576762386756746a2f686b6b55646463393279424935365073554d67476e707636425a2b506c3843675945416e464b7777352f427658394b63454462577758740a6a6435744746464238414e644a64654835346d7a63685253594d6269697774376143386b79656a49696543765a647577794770464b426a6577663463747965430a324a5a673638484458436e374d4f6b643243556569457670674d5352566d3769342f3362692b6856316c616d66524a41673662586672476361454b7363326f710a7072337971307a696f4e354e72636d6c4f6a39726e635543675941593471704a732f6e6c6b426c7356766662484f5133667651374f6a4f4e4f6472655a442f470a5a53495756444e6d53735a49426f4e412f6c67596c6646787a594d4f3635506b41424479683953536d4761544e43424945644571435447666b454c30746b46780a78384c7643637352374a4133764c527349696542593635726749555554464d5632582b4c777a334866716f56384a52727278584e7939437a6d4d51686961326f0a43686d3853514b4267464e2b39726c624f52646c62517a5a696a5a4174474c6955336533694974347a624a4358445744317635624b303573364939444c3867300a494e776e445a30777451696f38546e556a4a334859654954726930796b4d3379723631397a635437624973636b6d764577724a422b5847314c77707a6b5236730a364b4b453347765642584d462f70376851544548617a7a57564d63376e2b5a333542776a4c2b496b6269425a574455314c4f65510a2d2d2d2d2d454e44205253412050524956415445204b45592d2d2d2d2d0a"

func init() {
	logger.AddOutput(logger.Stdout{
		MinLevel: logger.DEBUG, //logger.DEBUG,
		Colored:  true,
	})
}

func main(){
	newNetwork()
	ctx := context.Background()
	n, err := network.UnmarshalFromPrivateKey(testNetwork)
	if err != nil {
		panic(err)
	}
	d, err := discovery.New(ctx, n, 10, false)
	if err != nil {
		panic(err)
	}
	host, _ := os.Hostname()
	d.LocalNode.SetInfo("host", host)
	for {
		d.LocalNode.SetInfo("Updated", time.Now().Format(time.RFC822))
		logger.Info("peers:")
		for _, peer := range d.WaitForPeers(1, 0 * time.Second) {
			logger.Infof("%s", peer)
		}
		time.Sleep(5 * time.Second)
	}
}

func newNetwork() {
	n, err := network.New()
	if err != nil {
		panic(err)
	}
	logger.Info("New network:")
	logger.Infof("Private key: %s", n.MarshalFromPrivateKey())
	logger.Infof("Public key: %s", n.Marshal())

}