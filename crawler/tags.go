package crawler

// TAGS 検索するタグの配列
// https://qiita.com/tags?page=1 を見ながら qiita のタグ一覧の上から数式が含まれそうなのを適当にピックアップしている
// 投稿数1000以上 (35行目の Latex より上) は割と手広く選んだが それ以下は学問チックな言葉をチョイス
var TAGS = [...]string{
	"機械学習",
	"DeepLearning",
	"TensorFlow",
	"OpenCV",
	"AtCoder",
	"MachineLearning",
	"競技プログラミング",
	"pandas",
	"数学",
	"アルゴリズム",
	"Haskell",
	"Keras",
	"Elasticsearch",
	"自然言語処理",
	"Jupyter",
	"numpy",
	"データ分析",
	"matplotlib",
	"AI",
	"Blockchain",
	"正規表現",
	"画像処理",
	"Anaconda",
	"Ethereum",
	"深層学習",
	"統計学",
	"PyTorch",
	"Bluemix",
	"LaTex",
	"データサイエンス",
	"論文読み",
	"強化学習",
	"画像認識",
	"量子コンピュータ",
	"人工知能",
	"数値計算",
	"学習",
	"ニューラルネットワーク",
	"データ構造",
	"暗号",
	"統計",
	"機械学習入門",
	"線形代数",
	"物理",
	"化学",
	"科学技術計算",
	"シミュレーション",
	"代数学",
	"最小二乗法",
	"math",
	"数値積分",
	"ゲーム理論",
	"公式証明",
	"電磁気学",
}
