package chaincode

//文章实体
type Article struct {
	Key string //文章存证号
	Name string //文章名称
	DataFinger string //文章数据指纹
	Author string //作者
	Submitter string //提交者
	CompletionTime string //文章完成时间
	ReadPrice string //阅读价格
	DownloadPrice string //下载价格
	CheckStatus string //审核状态
	Status string //链上文章有效状态
}

type DataFinger struct {
	FileDataFinger string //文章数据指纹
	Key string //文章存证号
}