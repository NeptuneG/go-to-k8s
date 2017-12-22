package version

// 下記情報はgo linkerでコンパイルの時に設定する
var (
	// ビルドタイムスダンプ
	BuildTimestamp = ""
	// gitコッミトハッシュ
	Commit = ""
	// リリースバージョン
	Release = ""
)
