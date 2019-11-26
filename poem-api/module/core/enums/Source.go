package enums

/**
 *来源
 */
type SourceLevel int

var SourceMap = map[SourceLevel]string{GUSHIWEN: "古诗文网"}


const (
	/**
	* 古诗文网
	*/
	GUSHIWEN SourceLevel = iota
)
