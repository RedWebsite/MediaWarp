package static

// 通用头部

const basicPath = "/MediaWarp/static/"

const (
	WebModifyHeaderStart = "<head>\n" + `<!-- MediaWarp Web 页面修改功能 -->` + "\n"
	WebModifyHeaderEnd   = `<!-- MediaWarp Web 页面修改功能 end -->`
)

// 通用嵌入脚本
const (
	VideoTogether     = `<script src="https://2gether.video/release/extension.website.user.js"></script>`
	ExternalPlayerUrl = `<script src="` + basicPath + `embyExternalUrl/embyWebAddExternalUrl/embyLaunchPotplayer.js"></script>`
)

// Emby 相关嵌入脚本
const (
	EmbyCrx = `<link rel="stylesheet" id="theme-css" href="` + basicPath + `emby-crx/static/css/style.css" type="text/css" media="all" />
    <script src="` + basicPath + `emby-crx/static/js/common-utils.js"></script>
    <script src="` + basicPath + `emby-crx/static/js/jquery-3.6.0.min.js"></script>
    <script src="` + basicPath + `emby-crx/static/js/md5.min.js"></script>
    <script src="` + basicPath + `emby-crx/content/main.js"></script>`
	EmbyActorPlus  = `<script src="` + basicPath + `emby-web-mod/actorPlus/actorPlus.js"></script>`
	EmbyFanartShow = `<script src="` + basicPath + `emby-web-mod/fanart_show/fanart_show.js"></script>`
	EmbyDanmaku    = `<script src="` + basicPath + `dd-danmaku/ede.js" defer></script>`
)

// Jellyfin 相关嵌入脚本
const (
	JellyfinCrx = `<link rel="stylesheet" id="theme-css" href="` + basicPath + `jellyfin-crx/static/css/style.css" type="text/css" media="all" />
    <script src="` + basicPath + `jellyfin-crx/static/js/common-utils.js"></script>
    <script src="` + basicPath + `jellyfin-crx/static/js/jquery-3.6.0.min.js"></script>
    <script src="` + basicPath + `jellyfin-crx/static/js/md5.min.js"></script>
    <script src="` + basicPath + `jellyfin-crx/content/main.js"></script>`
	JellyfinActorPlus  = EmbyActorPlus
	JellyfinFanartShow = EmbyFanartShow
	JellyfinDanmaku    = `<script src="` + basicPath + `jellyfin-danmaku/ede.js" defer></script>`
)

// 飞牛影视相关嵌入脚本
const (
	FNTVPlayInfoHook = `<script src="` + basicPath + `fn-danmaku/fntv-play-info-hook.js"></script>`
	FNTVDanmaku      = `<script src="` + basicPath + `fn-danmaku/fn-danmaku.js" defer></script>`
)
