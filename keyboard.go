package ElmaViber

type InputFieldState string

const (
	Regular   InputFieldState = "regular"
	Minimized InputFieldState = "minimized"
	Hidden    InputFieldState = "hidden"
)

type Keyboard struct {
	Type                string             `json:"Type"`
	BgColor             string             `json:"BgColor,omitempty"`
	DefaultHeight       bool               `json:"DefaultHeight,omitempty"`
	CustomDefaultHeight int                `json:"CustomDefaultHeight,omitempty"`
	HeightScale         int                `json:"HeightScale,omitempty"`
	ButtonsGroupColumns int                `json:"ButtonsGroupColumns,omitempty"`
	ButtonsGroupRows    int                `json:"ButtonsGroupRows,omitempty"`
	InputFieldState     InputFieldState    `json:"InputFieldState,omitempty"`
	FavoritesMetadata   *FavoritesMetadata `json:"FavoritesMetadata,omitempty"`
	Buttons             []Button           `json:"Buttons"`
}

type TypeMetadata string

const (
	GifMetadata   TypeMetadata = "gif"
	LinkMetadata  TypeMetadata = "link"
	VideoMetadata TypeMetadata = "video"
)

type FavoritesMetadata struct {
	Type            TypeMetadata `json:"type"`
	URL             string       `json:"url"`
	Title           string       `json:"title,omitempty"`
	Thumbnail       string       `json:"thumbnail,omitempty"`
	Domain          string       `json:"domain,omitempty"`
	Width           int          `json:"width,omitempty"`
	Height          int          `json:"height,omitempty"`
	AlternativeUrl  string       `json:"alternativeUrl,omitempty"`
	AlternativeText string       `json:"alternativeText,omitempty"`
}

type BgMediaType string

const (
	BgPicture BgMediaType = "picture"
	BgGif     BgMediaType = "gif"
)

type ScaleType string

const (
	Crop ScaleType = "crop"
	Fill ScaleType = "fill"
	Fit  ScaleType = "fit"
)

type ActionType string

const (
	Reply          ActionType = "reply"
	OpenUrl        ActionType = "open-url"
	LocationPicker ActionType = "location-picker"
	SharePhone     ActionType = "share-phone"
	None           ActionType = "none"
)

type TextVAlign string

const (
	Top    TextVAlign = "top"
	Middle TextVAlign = "middle"
	Bottom TextVAlign = "bottom"
)

type TextHAlign string

const (
	Left   TextHAlign = "left"
	Center TextHAlign = "center"
	Right  TextHAlign = "right"
)

type TextSize string

const (
	SmallSize   TextSize = "small"
	RegularSize TextSize = "regular"
	LargeSize   TextSize = "large"
)

type OpenURLType string

const (
	Internal OpenURLType = "internal"
	External OpenURLType = "external"
)

type OpenURLMediaType string

const (
	NotMediaOpenURL OpenURLMediaType = "not-media"
	VideoOpenURL    OpenURLMediaType = "video"
	GifOpenURL      OpenURLMediaType = "gif"
	PictureOpenURL  OpenURLMediaType = "picture"
)

type ActionButton string

const (
	Forward        ActionButton = "forward"
	Send           ActionButton = "send"
	OpenExternally ActionButton = "open-externally"
	SendToBot      ActionButton = "send-to-bot"
)

type TitleType string

const (
	Domain  TitleType = "domain"
	Default TitleType = "default"
)

type Mode string

const (
	FullscreenMode          Mode = "fullscreen"
	FullscreenPortraitMode  Mode = "fullscreen-portrait"
	FullscreenLandscapeMode Mode = "fullscreen-landscape"
	PartialSizeMode         Mode = "partial-size"
)

type FooterType string

const (
	DefaultFooter FooterType = "default"
	HiddenFooter  FooterType = "hidden"
)

type Frame struct {
	BorderWidth  int    `json:"BorderWidth,omitempty"`
	BorderColor  string `json:"BorderColor,omitempty"`
	CornerRadius int    `json:"CornerRadius,omitempty"`
}

type Map struct {
	Latitude  string `json:"Latitude,omitempty"`
	Longitude string `json:"Longitude,omitempty"`
}

type InternalBrowser struct {
	ActionButton        ActionButton `json:"ActionButton,omitempty"`
	ActionPredefinedURL string       `json:"ActionPredefinedURL,omitempty"`
	TitleType           TitleType    `json:"TitleType,omitempty"`
	CustomTitle         string       `json:"CustomTitle,omitempty"`
	Mode                Mode         `json:"Mode,omitempty"`
	FooterType          FooterType   `json:"FooterType,omitempty"`
	ActionReplyData     string       `json:"ActionReplyData,omitempty"`
}

type MediaPlayer struct {
	Title        string `json:"Title"`
	Subtitle     string `json:"Subtitle"`
	ThumbnailURL string `json:"ThumbnailURL"`
	Loop         bool   `json:"Loop"`
}

type Button struct {
	Columns             int              `json:"Columns,omitempty"`
	Rows                int              `json:"Rows,omitempty"`
	BgColor             string           `json:"BgColor,omitempty"`
	Silent              bool             `json:"Silent,omitempty"`
	BgMediaType         BgMediaType      `json:"BgMediaType,omitempty"`
	BgMedia             string           `json:"BgMedia,omitempty"`
	BgMediaScaleType    ScaleType        `json:"BgMediaScaleType,omitempty"`
	ImageScaleType      ScaleType        `json:"ImageScaleType,omitempty"`
	BgLoop              bool             `json:"BgLoop,omitempty"`
	ActionType          ActionType       `json:"ActionType,omitempty"`
	ActionBody          string           `json:"ActionBody"`
	Image               string           `json:"Image,omitempty"`
	Text                string           `json:"Text,omitempty"`
	TextVAlign          TextVAlign       `json:"TextVAlign,omitempty"`
	TextHAlign          TextHAlign       `json:"TextHAlign,omitempty"`
	TextOpacity         int              `json:"TextOpacity,omitempty"`
	TextSize            TextSize         `json:"TextSize,omitempty"`
	OpenURLType         OpenURLType      `json:"OpenURLType,omitempty"`
	OpenURLMediaType    OpenURLMediaType `json:"OpenURLMediaType,omitempty"`
	TextBgGradientColor string           `json:"TextBgGradientColor,omitempty"`
	TextShouldFit       bool             `json:"TextShouldFit,omitempty"`
	InternalBrowser     *InternalBrowser `json:"InternalBrowser,omitempty"`
	Map                 *Map             `json:"Map,omitempty"`
	Frame               *Frame           `json:"Frame,omitempty"`
	MediaPlayer         *MediaPlayer     `json:"MediaPlayer,omitempty"`
}

//https://developers.viber.com/docs/tools/keyboards/
func GetKeyboard() Keyboard {
	return Keyboard{
		Type: "keyboard",
	}
}

func (k *Keyboard) AddButton(b Button) {
	k.Buttons = append(k.Buttons, b)
}
