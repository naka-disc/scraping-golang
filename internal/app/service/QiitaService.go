package service

import (
	"naka-disc/scraping-golang/internal/app/api"
	"naka-disc/scraping-golang/internal/app/dao"
)

// DBのマイグレーション
func Migration() {
	dao.Migration()
}

// Qiitaのスクレイピング
func ScrapingForQiita() {
	// Qiitaのトップページ(トレンド)をスクレイピングして、記事一覧取得
	slice, ok := api.GetQiitaTrend()
	if !ok {
		return
	}

	// DBに全部保存
	dao.SaveQiitaArticleList(slice)
}
