package main

type PricingType int;
type MobileType int;
type StillRating int;

const (
	MobileType_UNKNOWN MobileType = iota
	MobileType_PC_ONLY
	MobileType_MOBILE_ONLY
	MobileType_BOTH
)

const (
	PricingType_UNKNOWN PricingType = iota
	PricingType_FREE 
	PricingType_ONE_TIME
	PricingType_SUBSCRIPTION
	PricingType_FREEMIUM
	PricingType_FREE_WITH_ADS
	PricingType_FREE_WITH_IN_APP_PURCHASES
	PricingType_EXTERNAL_SUBSCRIPTION
)

const (
	stillRating_UNKNOWN StillRating = iota
	stillRating_ONE
	stillRating_TWO
	stillRating_THREE
	stillRating_FOUR
	stillRating_FIVE
	stillRating_FIVE_PLUS
)

type APP struct {
	id string `yaml:"-"`
	name string `yaml:"name"`
	primary_src string `yaml:"primary_src"`
	src_pkg_name string `yaml:"src_pkg_name"`
	icon_url string `yaml:"icon_url"`
	author string `yaml:"author"`
	summary string `yaml:"summary"`
	description string `yaml:"description"`
	categories []string `yaml:"categories"`
	keywords []string `yaml:"keywords"` // Optional
	mimetypes []string `yaml:"mimetypes"` // Optional
	license string `yaml:"license"` // Optional, assumes Proprietary if not set
	pricing PricingType `yaml:"pricing"` // Optional, assumes FREE if not set
	mobile MobileType `yaml:"mobile"` // Optional, assumes UNKNOWN if not set
	still_rating StillRating `yaml:"still_rating"` // Optional, assumes UNKNOWN if not set
	still_rating_notes string `yaml:"still_rating_notes"` // Optional
	homepage string `yaml:"homepage"` // Optional
	donate_url string `yaml:"donate_url"` // Optional
	screenshot_urls []string `yaml:"ss_urls"` // Optional
	demo_urls string `yaml:"demo_urls"` // Optional
	addons []string `yaml:"addons"` // Optional
}

func main() {

}