package entity

// Qiita記事エンティティ
type QiitaArticle struct {
	Id           uint   `gorm:"primarykey"`
	ScrapingDate string // スクレイピング日時
	Title        string // 記事タイトル
	Url          string // 記事URL
}

// コンストラクタ用処理。
// この関数を実行してインスタンスを生成すること。
func NewQiitaArticle() *QiitaArticle {
	e := new(QiitaArticle)
	return e
}
