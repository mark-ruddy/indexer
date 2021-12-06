package fetcher

const (
	RARIBLE    = "Rarible"
	CONTEXT    = "Context"
	CONVO      = "Convo"
	TWITTER    = "Twtter"
	OPENSEA    = "Opensea"
	ZORA       = "Zora"
	FOUNDATION = "Foundation"
	SHOWTIME   = "Showtime"
	SYBIL      = "Sybil"
	SUPERRARE  = "Superrare"
	INFURA     = "Infura"
)

const (
	SuperrareContractAddress  = "0x41a322b28d0ff354040e2cbc676f0320d8c8850d"
	OpenSeaContractAddress    = "0x495f947276749ce646f68ac8c248420045cb7b5e"
	RaribleContractAddress    = "0xd07dc4262bcdbf85190c01c996b4c06a461d2430"
	FoundationContractAddress = "0x3b3ee1931dc30c1957379fac9aba94d1c48a5405"
	ZoraContractAddress       = "0xabefbc9fd2f806065b4f3c237d4b59d9a97bcac7"
	ContextContractAddress    = "ctx"
)

const (
	ContextUrl   = "https://context.app/api/profile/%s"
	SuperrareUrl = "https://superrare.com/api/v2/user?address=%s"

	// FoundationUrl Usage/Docs: https://thegraph.com/hosted-service/subgraph/f8n/fnd
	FoundationUrl = "https://api.thegraph.com/subgraphs/name/f8n/fnd"

	// OpenSeaUrl Usage/Docs: https://docs.opensea.io/reference/api-overview
	OpenSeaUrl = "https://api.opensea.io/api/v1"

	// ZoraUrl Usage/Docs: https://thegraph.com/hosted-service/subgraph/ourzora/zora-v1
	ZoraUrl = "https://api.thegraph.com/subgraphs/name/ourzora/zora-v1"

	// RaribleUrl Usage/Docs: https://api.rarible.org/v0.1/doc
	RaribleUrl = "https://api.rarible.org/v0.1"

	RaribleFollowingUrl = "https://api-mainnet.rarible.com/marketplace/api/v4/followings?owner=%s"
	RaribleFollowerUrl  = "https://api-mainnet.rarible.com/marketplace/api/v4/followers?user=%s"
)

type ConnectionEntryList struct {
	Conn []ConnectionEntry
	Err  error
	msg  string
}
type ConnectionEntry struct {
	From     string
	To       string
	Platform string
}

type IdentityEntryList struct {
	OpenSea             []UserOpenSeaIdentity
	Twitter             []UserTwitterIdentity
	Superrare           []UserSuperrareIdentity
	Rarible             []UserRaribleIdentity
	Context             []UserContextIdentity
	Zora                []UserZoraIdentity
	Foundation          []UserFoundationIdentity
	FoundationNonSocial []UserFoundationIdentityNonSocial
	Showtime            []UserShowtimeIdentity
	Ens                 string
}

type IdentityEntry struct {
	OpenSea             *UserOpenSeaIdentity
	Twitter             *UserTwitterIdentity
	Superrare           *UserSuperrareIdentity
	Rarible             *UserRaribleIdentity
	Context             *UserContextIdentity
	Zora                *UserZoraIdentity
	Ens                 *UserEnsIdentity
	Foundation          *UserFoundationIdentity
	FoundationNonSocial *UserFoundationIdentityNonSocial
	Showtime            *UserShowtimeIdentity
	Err                 error
	Msg                 string
}

type UserTwitterIdentity struct {
	Handle     string
	DataSource string
}

type UserRaribleIdentity struct {
	Username string
	Homepage string

	Owned struct {
		Total int `json:"total"`
		Items []struct {
			RaribleItem
		} `json:"items"`
	}
	Created struct {
		Total int `json:"total"`
		Items []struct {
			RaribleItem
		} `json:"items"`
	}
	Activities []struct {
		ID              string `json:"id"`
		Type            string `json:"@type"`
		From            string `json:"from"`
		Owner           string `json:"owner"`
		Contract        string `json:"contract"`
		TokenID         string `json:"tokenID"`
		Value           string `json:"value"`
		TransactionHash string `json:"transactionHash"`
		Date            string `json:"date"`
	} `json:"activities"`
	DataSource string
}

type UserOpenSeaIdentity struct {
	Username        string
	Homepage        string
	ProfileImageUrl string
	Assets          []struct {
		ID               int    `json:"id"`
		TokenID          string `json:"token_id"`
		NumSales         int    `json:"num_sales"`
		ImageUrl         string `json:"image_url"`
		ImagePreviewUrl  string `json:"image_preview_url"`
		ImageOriginalUrl string `json:"image_original_url"`
		AnimationUrl     string `json:"animation_url"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		Permalink        string `json:"permalink"`
		Creator          struct {
			User struct {
				Username string `json:"username"`
			} `json:"user"`
			ProfileImageUrl string `json:"profile_img_url"`
			Address         string `json:"address"`
		} `json:"creator"`
	} `json:"assets"`
	DataSource string
}

type UserEnsIdentity struct {
	Ens        string
	DataSource string
}

type UserContextIdentity struct {
	FollowerCount int
	Username      string
	Website       string
	DataSource    string
}

type UserSuperrareIdentity struct {
	Username       string
	Homepage       string
	Location       string
	Bio            string
	InstagramLink  string
	TwitterLink    string
	SteemitLink    string
	Website        string
	SpotifyLink    string
	SoundCloudLink string
	DataSource     string
}

type UserFoundationIdentity struct {
	Username   string
	Bio        string
	Tiktok     string
	Twitch     string
	Discord    string
	Twitter    string
	Website    string
	Youtube    string
	Facebook   string
	Snapchat   string
	Instagram  string
	DataSource string
}

type UserFoundationIdentityNonSocial struct {
	IsAdmin         bool
	NetRevenueInETH string
	Nfts            []struct {
		TokenIPFSPath      string `json:"tokenIPFSPath"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		Image              string `json:"image"`
		LastSalePriceInETH string `json:"lastSalePriceInETH"`
		DateMinted         string `json:"dateMinted"`
	} `json:"nfts"`
	Creator struct {
		NetSalesInETH          string `json:"netSalesInETH"`
		NetSalesPendingInETH   string `json:"netSalesPendingInETH"`
		NetRevenueInETH        string `json:"netRevenueInETH"`
		NetRevenueInPendingETH string `json:"netRevenueInPendingETH"`
	} `json:"creator"`
	Withdrawals []struct {
		AmountInETH string `json:"amountInETH"`
		Date        string `json:"date"`
	}
	DataSource string
}

type UserZoraIdentity struct {
	Username string
	Website  string

	// Collection is the media(NFTs) owned by the user, while Creations is media created by the user
	Collection []struct {
		ZoraMedia
	} `json:"collection"`
	Creations []struct {
		ZoraMedia
	} `json:"creations"`
	CurrentBids []struct {
		ID                 string `json:"id"`
		Currency           string `json:"currency"`
		Amount             string `json:"amount"`
		CreatedAtTimestamp string `json:"createdAtTimestamp"`
	} `json:"currentBids"`
	DataSource string
}

type UserShowtimeIdentity struct {
	Name             string
	Username         string
	Bio              string
	TwitterHandle    string
	LinkTreeHandle   string
	CryptoArtHandle  string
	FoundationHandle string
	HicetnuncHandle  string
	OpenseaHandle    string
	RaribleHandle    string
	DataSource       string
}

type RaribleConnectionResp struct {
	Following struct {
		From string `json:"owner"`
		To   string `json:"user"`
	} `json:"following"`
}

type ContextAppResp struct {
	FollowerCount int               `json:"followerCount"`
	Ens           map[string]string `json:"ens"`
	Profiles      map[string]([]struct {
		Address  string `json:"address,omitempty"`
		Contract string `json:"contract,omitempty"`
		Url      string `json:"url,omitempty"`
		Website  string `json:"website,omitempty"`
		Username string `json:"username,omitempty"`
	}) `json:"profiles"`
}

type ContextConnection struct {
	Relationships []struct {
		Actor string `json:"actor"`
	} `json:"relationships"`
	Profiles map[string]([]struct {
		Address string `json:"address"`
	}) `json:"profiles"`
}

type SuperrareProfile struct {
	Result struct {
		Username       string `json:"username"`
		Location       string `json:"location"`
		Bio            string `json:"bio"`
		InstagramLink  string `json:"instagramLink"`
		TwitterLink    string `json:"twitterLink"`
		SteemitLink    string `json:"steemitLink"`
		Website        string `json:"website"`
		SpotifyLink    string `json:"spotifyLink"`
		SoundCloudLink string `json:"soundcloudLink"`
	} `json:"result"`
}

type FoundationProfileNonSocial struct {
	Data struct {
		Accounts []struct {
			IsAdmin         bool   `json:"isAdmin"`
			NetRevenueInETH string `json:"netRevenueInETH"`

			Nfts []struct {
				TokenIPFSPath      string `json:"tokenIPFSPath"`
				Name               string `json:"name"`
				Description        string `json:"description"`
				Image              string `json:"image"`
				LastSalePriceInETH string `json:"lastSalePriceInETH"`
				DateMinted         string `json:"dateMinted"`
			} `json:"nfts"`
			Creator struct {
				NetSalesInETH          string `json:"netSalesInETH"`
				NetSalesPendingInETH   string `json:"netSalesPendingInETH"`
				NetRevenueInETH        string `json:"netRevenueInETH"`
				NetRevenueInPendingETH string `json:"netRevenueInPendingETH"`
			} `json:"creator"`
			Withdrawals []struct {
				AmountInETH string `json:"amountInETH"`
				Date        string `json:"date"`
			} `json:"withdrawals"`
		} `json:"accounts"`
	} `json:"data"`
}

type OpenSeaProfileAccount struct {
	Data struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		ProfileImageUrl string `json:"profile_img_url"`
	} `json:"data"`
}

type OpenSeaProfileNft struct {
	Assets []struct {
		ID               int    `json:"id"`
		TokenID          string `json:"token_id"`
		NumSales         int    `json:"num_sales"`
		ImageUrl         string `json:"image_url"`
		ImagePreviewUrl  string `json:"image_preview_url"`
		ImageOriginalUrl string `json:"image_original_url"`
		AnimationUrl     string `json:"animation_url"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		Permalink        string `json:"permalink"`

		// The owner of the NFT may or may not be the creator of the NFT as well
		Creator struct {
			User struct {
				Username string `json:"username"`
			} `json:"user"`
			ProfileImageUrl string `json:"profile_img_url"`
			Address         string `json:"address"`
		} `json:"creator"`
	} `json:"assets"`
}

// ZoraMedia is used for JSON unmarshalling of NFTs from the Zora GraphQL API
type ZoraMedia struct {
	ID                 string `json:"id"`
	TransactionHash    string `json:"transactionHash"`
	ContentHash        string `json:"contentHash"`
	MetadataHash       string `json:"metadataHash"`
	ContentURI         string `json:"contentURI"`
	MetadataURI        string `json:"metadataURI"`
	CreatedAtTimestamp string `json:"createdAtTimestamp"`

	CurrentAsk struct {
		Amount             string `json:"amount"`
		CreatedAtTimestamp string `json:"createdAtTimestamp"`
		Currency           struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"currency"`
	} `json:"currentAsk"`
}

type ZoraProfile struct {
	Data struct {
		Users []struct {
			Collection []struct {
				ZoraMedia
			} `json:"collection"`
			Creations []struct {
				ZoraMedia
			} `json:"creations"`
			CurrentBids []struct {
				ID                 string `json:"id"`
				Currency           string `json:"currency"`
				Amount             string `json:"amount"`
				CreatedAtTimestamp string `json:"createdAtTimestamp"`
			} `json:"currentBids"`
		} `json:"users"`
	} `json:"data"`
}

// RaribleItem is used for JSON unmarshalling of NFTs from the Rarible HTTPS API
type RaribleItem struct {
	ID         string `json:"id"`
	Blockchain string `json:"blockchain"`
	Contract   string `json:"contract"`
	TokenID    string `json:"tokenId"`
	LazySupply string `json:"lazySupply"`
	MintedAt   string `json:"mintedAt"`
	Supply     string `json:"supply"`
	TotalStock string `json:"totalStock"`
}

type RaribleItemProfile struct {
	Total int `json:"total"`
	Items []struct {
		RaribleItem
	} `json:"items"`
}

type RaribleUserActivityProfile struct {
	// Activities is an array of activity items defined in below struct
	// Some fields may be used by TRANSFER and others by BUY,SELL activity types
	Activities []struct {
		ID              string `json:"id"`
		Type            string `json:"@type"`
		From            string `json:"from"`
		Owner           string `json:"owner"`
		Contract        string `json:"contract"`
		TokenID         string `json:"tokenID"`
		Value           string `json:"value"`
		TransactionHash string `json:"transactionHash"`
		Date            string `json:"date"`
	} `json:"activities"`
}

type FoundationIdentity struct {
	Data struct {
		User struct {
			Username string `json:"username"`
			Bio      string `json:"bio"`
			Links    struct {
				Tiktok struct {
					Handle string `json:"handle"`
				} `json:"tiktok"`
				Twitch struct {
					Handle string `json:"handle"`
				} `json:"twitch"`
				Discord struct {
					Handle string `json:"handle"`
				} `json:"discord"`
				Twitter struct {
					Handle string `json:"handle"`
				} `json:"twitter"`
				Website struct {
					Handle string `json:"handle"`
				} `json:"website"`
				Youtube struct {
					Handle string `json:"handle"`
				} `json:"youtube"`
				Facebook struct {
					Handle string `json:"handle"`
				} `json:"facebook"`
				Snapchat struct {
					Handle string `json:"handle"`
				} `json:"snapchat"`
				Instagram struct {
					Handle string `json:"handle"`
				} `json:"instagram"`
			} `json:"links"`
			TwitSocialVerifs []struct {
				Username string `json:"username"`
			} `json:"twitSocialVerifs"`
			InstaSocialVerifs []struct {
				Username string `json:"username"`
			} `json:"instaSocialVerifs"`
		} `json:"user"`
	} `json:"data"`
}
