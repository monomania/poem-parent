package enums

/**
 *诗词种类
 */
type PoemTypeLevel int

var PoemTypeMap = map[PoemTypeLevel]string{SHI: "诗", CI: "词", QU: "曲", WEN_YAN_WEN: "文言文", CI_FU: "辞赋"}

const (
	/**
	* 诗
	*/
	SHI PoemTypeLevel = iota
	/**
	 * 词
	 */
	CI
	/**
	 * 曲
	 */
	QU
	/**
	 * 文言文
	 */
	WEN_YAN_WEN
	/**
	 * 辞赋
	 */
	CI_FU
)
