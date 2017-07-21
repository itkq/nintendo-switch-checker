package nschecker

var SetSources = []Source{
	{
		Name:        "7net - Nintendo Switch スプラトゥーン2セット",
		URL:         "http://7net.omni7.jp/general/010007/170518splatoon",
		SoldOutText: `<button class="linkBtn" type="button" style="color: #fff;">SOLD OUT</button>`,
	},
	{
		Name:          "Amazon - Nintendo Switch スプラトゥーン2セット",
		URL:           "https://www.amazon.co.jp/dp/B072FGP433/ref=sr_1_11?ie=UTF8&qid=1495076407&tag=gamepedia-22&sr=8-11&keywords=switch+splatoon2",
		AvailableText: `この商品は、<a href="/gp/help/customer/display.html?ie=UTF8&amp;nodeId=643004">Amazon.co.jp</a> が販売、発送します。`,
	},
	{
		Name:        "楽天ブックス: Nintendo Switch スプラトゥーン2セット",
		URL:         "http://ts.books.rakuten.co.jp/rb/14943334/?scid=af_pc_etc&sc2id=af_103_1_10000645",
		SoldOutText: `<span class="status">ご注文できない商品`,
	},
	{
		Name:        "ヤマダウェブコム - Nintendo Switch スプラトゥーン2セット",
		URL:         "http://www.yamada-denkiweb.com/1178028018?q=%E3%82%B9%E3%83%97%E3%83%A9%E3%83%88%E3%82%A5%E3%83%BC%E3%83%B32",
		SoldOutText: `<button type="submit" class="btn btn-disabled btn-block" disabled="disabled">売り切れました</button>`,
	},
	{
		Name:        "ノジマオンライン: 【Switch】 ニンテンドースイッチ本体 スプラトゥーン2セット  HAC-S-KACEA",
		URL:         "https://online.nojima.co.jp/Nintendo-HAC-S-KACEA-%E3%80%90Switch%E3%80%91-%E3%83%8B%E3%83%B3%E3%83%86%E3%83%B3%E3%83%89%E3%83%BC%E3%82%B9%E3%82%A4%E3%83%83%E3%83%81%E6%9C%AC%E4%BD%93-%E3%82%B9%E3%83%97%E3%83%A9%E3%83%88%E3%82%A5%E3%83%BC%E3%83%B32%E3%82%BB%E3%83%83%E3%83%88/4902370537338/1/cd/",
		SoldOutText: `<span>完売御礼</span>`,
	},
	{
		Name:        "ヨドバシ.com - 任天堂 Nintendo Nintendo Switch スプラトゥーン2セット [Nintendo Switch本体]",
		URL:         "http://www.yodobashi.com/product/100000001003570628/",
		SoldOutText: `<div class="salesInfo"><p>予定数の販売を終了しました</p></div>`,
	},
	// {
	// 	Name:        "ローチケHMV: Nintendo Switch スプラトゥーン2セット",
	// 	URL:         "http://www.hmv.co.jp/artist_Game-Hard_000000000119544/item_Nintendo-Switch-%E3%82%B9%E3%83%97%E3%83%A9%E3%83%88%E3%82%A5%E3%83%BC%E3%83%B32%E3%82%BB%E3%83%83%E3%83%88_7916445",
	// 	SoldOutText: `<p class="date">限定盤の為完売しております。申し訳ございませんがご注文いただけません。</p>`,
	// },
	{
		Name:        "トイザらス: Nintendo Switch スプラトゥーン2 セット",
		URL:         "https://www.toysrus.co.jp/s/dsg-580782400",
		SoldOutText: `<span id="isStock_a">在庫あり</span>`,
		// AvailableText: `<span id="isStock_a">在庫あり</span>`,
	},
}
