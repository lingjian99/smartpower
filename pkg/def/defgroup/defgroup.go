package defgroup

const (
	GroupMemCntMin = 3  //群成员最小数量
	GroupMemCntMax = 30 //群成员最大数量
	GroupMchCntMax = 3  //每个群最大管理棋牌室数量
	GroupMax       = 10 //每个人群最大数量
)

type GroupMemberReq struct {
	UserId int64
	Lord   bool
}
