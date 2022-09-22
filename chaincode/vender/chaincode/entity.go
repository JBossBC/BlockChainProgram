package chaincode

//文章实体
type Article struct {
	Key string `jason:"key"` //文章存证号
	Name string `jason:"name"` //文章名称
	DataFinger string `jason:"dataFinger"`//文章数据指纹
	Author string `jason:"author"`//作者
	Submitter string `jason:"submitter"`//提交者
	CompletionTime string `jason:"completionTime"`//文章完成时间
	ReadPrice string `jason:"readPrice"`//阅读价格
	DownloadPrice string `jason:"downloadPrice"`//下载价格
	CheckStatus string `jason:"checkStatus"`//审核状态
	Status string `jason:"status"`//链上文章有效状态
}

type DataFinger struct {
	FileDataFinger string `jason:"fileDataFinger"`//文章数据指纹
	Key string `jason:"key"`//文章存证号
}
