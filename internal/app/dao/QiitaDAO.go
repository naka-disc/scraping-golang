package dao

import (
	"log"
	"naka-disc/scraping-golang/internal/app/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DBのマイグレーション
func Migration() {
	db := getDatabaseObject()
	// マイグレーション実行
	db.AutoMigrate(entity.NewQiitaArticle())
}

// GORMのデータベースオブジェクトを取得
// daoパッケージ内のみでの使用に留めるため、外には公開しない
func getDatabaseObject() *gorm.DB {
	// TODO: 接続エラーは発生しない想定で組んであるため、エラーハンドリングしていない
	db, _ := gorm.Open(sqlite.Open("database/database.sqlite"), &gorm.Config{})
	return db
}

// Qiita記事を一括保存
// @param slice 保存記事のエンティティスライス
func SaveQiitaArticleList(slice []entity.QiitaArticle) {
	db := getDatabaseObject()

	// スライス分繰り返し処理
	for _, v := range slice {
		db.Create(&v)
		log.Printf("SaveQiitaArticleList complete: ID = %d", v.Id)
	}
}
