package enums

import "strings"

/**
 *朝代
 */
type DynastyLevel int

var DynastyMap = map[DynastyLevel]string{QIN: "先秦", HAN: "两汉", WEI_JIN: "魏晋", NAN_BEI: "南北朝", SUI: "隋代", TANG: "唐代", WU: "五代", SONG: "宋代", JIN: "金朝", YUAN: "元代", MING: "明代", QING: "清代",JIN_XIAN:"近现代"}

func GetDynastyLevel(word string) DynastyLevel {
	if len(word) <= 0 {
		return -1
	}

	var level DynastyLevel
	for k, v := range DynastyMap {
		if strings.EqualFold(v, word) {
			level = k
			break;
		}
	}
	return level
}

const (
	//先秦
	QIN DynastyLevel = iota
	//两汉
	HAN
	//魏晋
	WEI_JIN
	//南北朝
	NAN_BEI
	//隋代
	SUI
	//唐代
	TANG
	//五代
	WU
	//宋代
	SONG
	//金朝
	JIN
	//元代
	YUAN
	//明代
	MING
	//清代
	QING
	//近现代
	JIN_XIAN
)
