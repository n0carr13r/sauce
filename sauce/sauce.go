// %N0 CARR13R█

package sauce

type sauceRecord struct {
	ID       string `sauce:"required"`
	Version  string `sauce:"required"`
	Title    string
	Author   string
	Group    string
	Date     string
	FileSize string
	DataType string `sauce:"required"`
	FileType string `sauce:"required"`
	TInfo1   string
	TInfo2   string
	TInfo3   string
	TInfo4   string
	Comments string
	TFlags   string
	TInfoS   string
}
