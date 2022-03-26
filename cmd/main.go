package main

import "naka-disc/scraping-golang/internal/app/service"

func main() {
	// マイグレーション処理
	// これを先にやらないと、テーブルがないためエラー頻発する
	service.Migration()

	// Qiitaのスクレイピング実行
	service.ScrapingForQiita()
}
