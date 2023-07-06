package bot

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
	"log"
	"strings"
	"time"
)
var (
	_self *openwechat.Self
)
// ConsoleQrCode 二维码控制台打印
func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	log.Println(q.ToString(true))
}

// StartBot bot 启动类
func StartBot()  {
	//bot := openwechat.DefaultBot() //默认模式
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = onMessage
	// 注册登陆二维码回调
	bot.UUIDCallback = ConsoleQrCode

	// 登陆
	if err := bot.Login(); err != nil {
		log.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	log.Printf("当前登录用户 %v \n", self)
	if err != nil {
		log.Println(err)
		return
	}
	_self = self
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()

}

func onMessage(msg *openwechat.Message)  {
	sender, err := msg.Sender()
	if err != nil {
		log.Println("获取发送人异常",err)
	}
	// 文本消息 &&  长度大于0 && 非自己消息 &&
	if msg.IsText() && len(msg.Content) > 0 && !msg.IsSendBySelf()  {
		var reply = ""
		if msg.IsAt() && msg.IsSendByGroup() {
			groupName := sender.NickName
			onGroupMsg(groupName,msg)
		}
		//朋友单聊
		if msg.IsSendByFriend() {
			log.Printf("发送人: 【%s】  【%s】",sender.NickName,msg.Content)
			reply = getMessageReplyByKeyWord(msg.Content)
			if reply != "" {
				msg.ReplyText(reply)
			}
		}
	}
}

func onGroupMsg(groupName string,msg *openwechat.Message) () {
	content := msg.Content
	//具体发送用户
	sender, err := msg.SenderInGroup()
	if err != nil {
		log.Println(err)
	}
	//群里具体消息发送者
	senderUserName := sender.NickName
	//用户艾特空格分割：@ - <0X2005> 全角空格--->  
	index := strings.Index(content, " ")
	splitContent := content[index+3:]
	splitContent = strings.Replace(splitContent," ","",-1)
	log.Printf("群组名：【%s】 发消息用户【%s】 消息【%s】 ",groupName,senderUserName,splitContent)
	if splitContent == "" {
		return
	}
	//关键词查找
	reply := getMessageReplyByKeyWord(splitContent)
	if reply != "" {
		reply = "@"+senderUserName + " " + reply
		msg.ReplyText(reply)
	}

}


func getMessageReplyByKeyWord(keys string)  string{
	info := SelectRecordForKeysInfo(keys)
	time.Sleep(time.Second * 1)
	if len(info) > 0 {
		return info[0].Details
	}
	return ""
}

