package constant

const (

	//group admin
	//	OrdinaryMember = 0
	//	GroupOwner     = 1
	//	Administrator  = 2
	//group application
	//	Application      = 0
	//	AgreeApplication = 1

	//friend related
	BlackListFlag         = 1
	ApplicationFriendFlag = 0
	FriendFlag            = 1
	RefuseFriendFlag      = -1

	//Websocket Protocol
	WSGetNewestSeq     = 1001
	WSPullMsg          = 1002
	WSSendMsg          = 1003
	WSPullMsgBySeqList = 1004
	WSPushMsg          = 2001
	WSKickOnlineMsg    = 2002
	WSDataError        = 3001

	///ContentType
	//UserRelated
	Text           = 101
	Picture        = 102
	Voice          = 103
	Video          = 104
	File           = 105
	AtText         = 106
	Merger         = 107
	Card           = 108
	Location       = 109
	Custom         = 110
	Revoke         = 111
	HasReadReceipt = 112
	Typing         = 113
	Quote          = 114
	Common         = 200
	GroupMsg       = 201

	//SysRelated

	FriendApplicationProcessedNotification = 1201 //AcceptFriendApplicationTip = 201
	FriendApplicationAddedNotification     = 1202 //AddFriendTip               = 202
	FriendAddedNotification                = 1203
	FriendDeletedNotification              = 1204
	FriendInfoChangedNotification          = 1205
	BlackAddedNotification                 = 1206
	BlackDeletedNotification               = 1207

	SelfInfoUpdatedNotification = 1303 //SetSelfInfoTip             = 204

	GroupCreatedNotification         = 1501 //CreateGroupTip            = 502
	JoinApplicationNotification      = 1502 //JoinGroupTip              = 504
	ApplicationProcessedNotification = 1503 //AcceptGroupApplicationTip = 507 RefuseGroupApplicationTip = 508
	MemberInvitedNotification        = 1504 //InviteUserToGroupTip      = 510
	MemberKickedNotification         = 1505 //KickGroupMemberTip        = 509
	GroupInfoChangedNotification     = 1506 //SetGroupInfoTip           = 506  TransferGroupOwnerTip     = 501
	MemberLeaveNotification          = 1507 //QuitGroupTip              = 505
	MemberEnterNotification          = 1508

	//MsgFrom
	UserMsgType = 100
	SysMsgType  = 200

	//SessionType
	SingleChatType = 1
	GroupChatType  = 2
	//token
	NormalToken  = 0
	InValidToken = 1
	KickedToken  = 2
	ExpiredToken = 3

	//MultiTerminalLogin
	//Full-end login, but the same end is mutually exclusive
	AllLoginButSameTermKick = 1
	//Only one of the endpoints can log in
	SingleTerminalLogin = 2
	//The web side can be online at the same time, and the other side can only log in at one end
	WebAndOther = 3
	//The PC side is mutually exclusive, and the mobile side is mutually exclusive, but the web side can be online at the same time
	PcMobileAndWeb = 4

	OnlineStatus  = "online"
	OfflineStatus = "offline"
	Registered    = "registered"
	UnRegistered  = "unregistered"

	//MsgReceiveOpt
	ReceiveMessage          = 0
	NotReceiveMessage       = 1
	ReceiveNotNotifyMessage = 2

	//OptionsKey
	IsHistory            = "history"
	IsPersistent         = "persistent"
	IsOfflinePush        = "offlinePush"
	IsUnreadCount        = "unreadCount"
	IsConversationUpdate = "conversationUpdate"
	IsSenderSync         = "senderSync"
)

var ContentType2PushContent = map[int64]string{
	Picture:  "[图片]",
	Voice:    "[语音]",
	Video:    "[视频]",
	File:     "[文件]",
	Text:     "你收到了一条文本消息",
	AtText:   "[有人@你]",
	GroupMsg: "你收到一条群聊消息",
	Common:   "你收到一条新消息",
}

const (
	AppOrdinaryUsers = 1
	AppAdmin         = 2

	GroupOrdinaryUsers = 1
	GroupOwner         = 2
	GroupAdmin         = 3

	GroupResponseAgree  = 1
	GroupResponseRefuse = -1

	Male   = 1
	Female = 2
)

const FriendAcceptTip = "You have successfully become friends, so start chatting"
