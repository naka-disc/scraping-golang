package api

import (
	"log"
	"naka-disc/scraping-golang/internal/app/entity"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Qiitaのトップページからトレンド記事を取得
// @return Qiita記事のエンティティスライス
// @return 成功／失敗
func GetQiitaTrend() ([]entity.QiitaArticle, bool) {
	var slice []entity.QiitaArticle

	// GETリクエスト送信
	u := "https://qiita.com/"
	res, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
		return slice, false
	}
	defer res.Body.Close()

	// ステータスコードチェック
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return slice, false
	}

	// レスポンスデータをGoQueryに読み込ませる
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return slice, false
	}

	// スクレイピング日時として、現在日時を取得しておく
	nowDatetime := time.Now().Format("2006-01-02 15:04:05")

	// articleタグを抽出し、そこからデータをさらに抽出
	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		// スクレイピングデータからの抽出
		href, _ := s.Find("a").Attr("href")
		title := s.Find("h2").Text()

		// エンティティ生成
		e := entity.NewQiitaArticle()
		e.ScrapingDate = nowDatetime
		e.Title = title
		e.Url = href
		slice = append(slice, *e)
	})

	return slice, true
}
