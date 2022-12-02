package account

type region string

const (
	cn_qingdao     region = "cn-qingdao"
	cn_beijing     region = "cn-beijing"
	cn_zhangjiakou region = "cn-zhangjiakou"
	cn_hangzhou    region = "cn-hangzhou"
	cn_shanghai    region = "cn-shanghai"
)

var regionNameMapping = map[region]string{
	cn_qingdao:     "青岛",
	cn_beijing:     "北京",
	cn_zhangjiakou: "张家口",
	cn_hangzhou:    "杭州",
	cn_shanghai:    "上海",
}
