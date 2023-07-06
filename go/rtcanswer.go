

// 说明 reflect track  ；web端视频到服务端 服务端再返回 
 
package rtc

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/pion/webrtc/v3"
)

var answerpeerConnectionCache = make(map[string]*webrtc.PeerConnection)

var ansmutex = sync.Mutex{}

// 缓存 answer peer
func getAnswerrPeerCache(key string) (*webrtc.PeerConnection, bool) {
	ansmutex.Lock()
	defer ansmutex.Unlock()
	value, ok := answerpeerConnectionCache[key]
	return value, ok
}

func InitAndCreateAnswer(id string, offersdp string) string {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// 配置了这个后会自动产生 空的 track  后续的流要自己手动替换
	// init := webrtc.RTPTransceiverInit{
	// 	Direction: webrtc.RTPTransceiverDirectionSendrecv,
	// }
	// senderVideo, err := peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeVideo, init)
	// senderAudio, err := peerConnection.AddTransceiverFromKind(webrtc.RTPCodecTypeAudio, init)
	// if err != nil {
	// 	panic(err)
	// }

	// outputVideoTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, "video", "pion-v")
	// senderVideo.Sender().ReplaceTrack(outputVideoTrack)
	// outputAudioTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, "audio", "pion-a")
	// senderAudio.Sender().ReplaceTrack(outputAudioTrack)
	// if err != nil {
	// 	panic(err)
	// }
	outputVideoTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, "video", "pion-v")
	outputAudioTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, "audio", "pion-a")
	if err != nil {
		panic(err)
	}

	// Add this newly created track to the PeerConnection
	_, err = peerConnection.AddTrack(outputVideoTrack)
	_, err = peerConnection.AddTrack(outputAudioTrack)

	if err != nil {
		panic(err)
	}
	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		fmt.Println("remote OnTrack")
		codec := track.Codec()
		fmt.Printf("远程媒体编码信息    %s ------> %d", codec.MimeType, codec.PayloadType)
		go func() {
			for {
				// Read RTP packets being sent to Pion
				rtp, _, readErr := track.ReadRTP()
				if readErr != nil {
					panic(readErr)
				}
				if track.Kind() == webrtc.RTPCodecTypeVideo {
					fmt.Println("Track is video")
					if writeErr := outputVideoTrack.WriteRTP(rtp); writeErr != nil {
						panic(writeErr)
					}
				} else if track.Kind() == webrtc.RTPCodecTypeAudio {
					fmt.Println("Track is audio")
					if writeErr := outputAudioTrack.WriteRTP(rtp); writeErr != nil {
						panic(writeErr)
					}
				}

			}
		}()

	})

	//暂存 连接信息
	answerpeerConnectionCache[id] = peerConnection

	dataChannel, err := peerConnection.CreateDataChannel("data-answer", nil)
	if err != nil {
		panic(err)
	}
	onAnswerDataChannelListener(id, dataChannel)

	onAnswerPcListener(id, peerConnection)

	setRemoteOffer(peerConnection, offersdp)

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}

	gatherComplete := make(chan struct{})

	go func() {
		// ICE收集
		peerConnection.OnICEGatheringStateChange(func(state webrtc.ICEGathererState) {
			fmt.Printf("Answer ICE Gathering 状态变更: %s\n", state.String())
			if state == webrtc.ICEGathererStateComplete {
				close(gatherComplete)
			}
		})
	}()

	<-gatherComplete

	payload, err := json.Marshal(answer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("answer %s", string(payload))
	return answer.SDP
}

// set remote offer
func setRemoteOffer(pc *webrtc.PeerConnection, offersdp string) {
	err := pc.SetRemoteDescription(CreateOfferSessionDescription(offersdp))
	if err != nil {
		panic(err)
	}
}

// pc  listener
func onAnswerPcListener(id string, peerConnection *webrtc.PeerConnection) {

	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			fmt.Println("Answer ICE Candidate ", candidate.ToJSON())
		}
	})

	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		fmt.Printf("Answer ICE Connection State -----------------> %s\n", connectionState.String())
	})

	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		fmt.Printf("Answer  Peer Connection 状态变更: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			fmt.Println("Answer Peer Connection 连接失败 退出")
			delete(answerpeerConnectionCache, id) //从answerpeerConnectionCache中移除id对应的信息
		}
	})

	// 监听对方datachannel消息
	peerConnection.OnDataChannel(func(dataChannel *webrtc.DataChannel) {
		fmt.Printf("remote DataChannel info '%s' \n", dataChannel.Label())
		dataChannel.OnOpen(func() {
			fmt.Printf("Data channel '%s'-'%d' open. \n", dataChannel.Label(), dataChannel.ID())
		})
		dataChannel.OnClose(func() {
			fmt.Println("remote data channel closed")
		})

		dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
			fmt.Printf("receive remote msg '%s': '%s'\n", dataChannel.Label(), string(msg.Data))
		})
	})

}

// listener datachannel
func onAnswerDataChannelListener(id string, dataChannel *webrtc.DataChannel) {
	dataChannel.OnOpen(func() {
		fmt.Printf("answer client create datachannel open  --------> Data channel '%s'-'%d' open. \n", dataChannel.Label(), dataChannel.ID())
		for range time.NewTicker(5 * time.Second).C {
			message := "你好 我是 answer"
			// Send the message as text
			sendTextErr := dataChannel.SendText(message)
			if sendTextErr != nil {
				fmt.Printf("数据通道发送数据异常 -------->%s", sendTextErr.Error())
				return
			}
		}
	})

	dataChannel.OnClose(func() {
		fmt.Println("answer data channel closed")
	})

	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Printf("answer client create datachannel receive msg '%s': '%s'\n", dataChannel.Label(), string(msg.Data))
	})

}
