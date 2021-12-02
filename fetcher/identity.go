package fetcher

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

const IdentityApiCount = 2

func (f *fetcher) FetchIdentity(address string) (IdentityEntryList, error) {

	var identityArr IdentityEntryList
	ch := make(chan IdentityEntry)

	// Part 1 - Demo data source
	// Context API
	go f.processContext(address, ch)
	// Superrare API
	go f.processSuperrare(address, ch)
	// Part 2 - Add other data source here
	go f.processFoundationNonSocial(address, ch)
	// TODO

	// Final Part - Merge entry
	for i := 0; i < IdentityApiCount; i++ {
		entry := <-ch
		if entry.Err != nil {
			zap.L().With(zap.Error(entry.Err)).Error("identity api error: " + entry.Msg)
			continue
		}
		if entry.OpenSea != nil {
			identityArr.OpenSea = append(identityArr.OpenSea, *entry.OpenSea)
		}
		if entry.Twitter != nil {
			entry.Twitter.Handle = convertTwitterHandle(entry.Twitter.Handle)
			identityArr.Twitter = append(identityArr.Twitter, *entry.Twitter)
		}
		if entry.Superrare != nil {
			identityArr.Superrare = append(identityArr.Superrare, *entry.Superrare)
		}
		if entry.Rarible != nil {
			identityArr.Rarible = append(identityArr.Rarible, *entry.Rarible)
		}
		if entry.Context != nil {
			identityArr.Context = append(identityArr.Context, *entry.Context)
		}
		if entry.Zora != nil {
			identityArr.Zora = append(identityArr.Zora, *entry.Zora)
		}
		if entry.Foundation != nil {
			identityArr.Foundation = append(identityArr.Foundation, *entry.Foundation)
		}
		if entry.FoundationNonSocial != nil {
			identityArr.FoundationNonSocial = append(identityArr.FoundationNonSocial, *entry.FoundationNonSocial)
		}
		if entry.Showtime != nil {
			identityArr.Showtime = append(identityArr.Showtime, *entry.Showtime)
		}
		if entry.Ens != nil {
			identityArr.Ens = entry.Ens.Ens
		}
	}

	return identityArr, nil
}

func (f *fetcher) processContext(address string, ch chan<- IdentityEntry) {
	var result IdentityEntry

	body, err := sendRequest(f.httpClient, RequestArgs{
		url:    fmt.Sprintf(ContextUrl, address),
		method: "GET",
	})
	if err != nil {
		result.Err = err
		result.Msg = "[processContext] fetch identity failed"
		ch <- result
		return
	}
	contextProfile := ContextAppResp{}
	err = json.Unmarshal(body, &contextProfile)
	if err != nil {
		result.Err = err
		result.Msg = "[processContext] identity response json unmarshal failed"
		ch <- result
		return
	}

	if value, ok := contextProfile.Ens[address]; ok {
		result.Ens = &UserEnsIdentity{
			Ens:        value,
			DataSource: CONTEXT,
		}
	}

	for _, profileList := range contextProfile.Profiles {
		for _, entry := range profileList {
			switch entry.Contract {
			case SuperrareContractAddress:
				result.Superrare = &UserSuperrareIdentity{
					Homepage:   entry.Url,
					Username:   entry.Username,
					DataSource: CONTEXT,
				}
			case OpenSeaContractAddress:
				result.OpenSea = &UserOpenSeaIdentity{
					Homepage:   entry.Url,
					Username:   entry.Username,
					DataSource: CONTEXT,
				}
			case RaribleContractAddress:
				result.Rarible = &UserRaribleIdentity{
					Homepage:   entry.Url,
					Username:   entry.Username,
					DataSource: CONTEXT,
				}
			case FoundationContractAddress:
				result.Foundation = &UserFoundationIdentity{
					Website:    entry.Website,
					Username:   entry.Username,
					DataSource: CONTEXT,
				}
			case ZoraContractAddress:
				result.Zora = &UserZoraIdentity{
					Website:    entry.Website,
					Username:   entry.Username,
					DataSource: CONTEXT,
				}
			case ContextContractAddress:
				result.Context = &UserContextIdentity{
					Username:      entry.Username,
					Website:       entry.Website,
					FollowerCount: contextProfile.FollowerCount,
					DataSource:    CONTEXT,
				}
			default:
			}
		}
	}

	ch <- result
	return
}

func (f *fetcher) processSuperrare(address string, ch chan<- IdentityEntry) {
	var result IdentityEntry

	body, err := sendRequest(f.httpClient, RequestArgs{
		url:    fmt.Sprintf(SuperrareUrl, address),
		method: "GET",
	})
	if err != nil {
		result.Err = err
		result.Msg = "[processSuperrare] fetch identity failed"
		ch <- result
		return
	}

	sprProfile := SuperrareProfile{}
	err = json.Unmarshal(body, &sprProfile)
	if err != nil {
		result.Err = err
		result.Msg = "[processSuperrare] identity response json unmarshal failednti"
		ch <- result
		return
	}

	newSprRecord := UserSuperrareIdentity{
		Username:       sprProfile.Result.Username,
		Location:       sprProfile.Result.Location,
		Bio:            sprProfile.Result.Bio,
		InstagramLink:  sprProfile.Result.InstagramLink,
		TwitterLink:    sprProfile.Result.TwitterLink,
		SteemitLink:    sprProfile.Result.SteemitLink,
		Website:        sprProfile.Result.Website,
		SpotifyLink:    sprProfile.Result.SpotifyLink,
		SoundCloudLink: sprProfile.Result.SoundCloudLink,
		DataSource:     SUPERRARE,
	}

	if newSprRecord.Username != "" || newSprRecord.Location != "" || newSprRecord.Bio != "" || newSprRecord.InstagramLink != "" ||
		newSprRecord.TwitterLink != "" || newSprRecord.SteemitLink != "" || newSprRecord.Website != "" ||
		newSprRecord.SpotifyLink != "" || newSprRecord.SoundCloudLink != "" {
		result.Superrare = &newSprRecord
	}

	ch <- result
}

// processFoundationNonSocial will query the Foundation GraphQL API
// The Foundation API does not provide endpoints for social media at this moment
// processFoundationNonSocial will get NFT, ETH Financial and Creator data for an address instead
func (f *fetcher) processFoundationNonSocial(address string, ch chan<- IdentityEntry) {
	var result IdentityEntry

	// GraphQL query that gets data from an account that matches the address
	gqlQuery := map[string]string{
		"query": fmt.Sprintf(`{
			accounts(where: {id: "%s"}) {
					isAdmin,
					netRevenueInETH,
					nfts {
						tokenIPFSPath,
						name,
						description,
						image,
						dateMinted,
						lastSalePriceInETH
					}
					creator {
						netSalesInETH,
						netSalesPendingInETH,
						netRevenueInETH,
						netRevenuePendingInETH
					}
					withdrawals {
						amountInETH,
						date
					}
				}
			}
		`, address),
	}

	jsonQuery, err := json.Marshal(gqlQuery)
	if err != nil {
		result.Err = err
		result.Msg = "[processFoundationNonSocial] Marshalling GraphQL query to JSON failed"
		ch <- result
		return
	}

	// sending a POST request which contains the GraphQL query in the body
	body, err := sendRequest(f.httpClient, RequestArgs{
		url:    FoundationUrl,
		method: "POST",
		body:   jsonQuery,
	})
	if err != nil {
		result.Err = err
		result.Msg = "[processFoundationNonSocial] fetch identity failed"
		ch <- result
		return
	}
	fndProfile := FoundationProfileNonSocial{}
	err = json.Unmarshal(body, &fndProfile)
	if err != nil {
		result.Err = err
		result.Msg = "[processFoundationNonSocial] identity response json unmarshal failed"
		ch <- result
		return
	}
	// using Accounts[0] here since the JSON response is an array of accounts but we are only using one address currently
	var newFndRecord UserFoundationIdentityNonSocial
	if len(fndProfile.Data.Accounts) > 0 {
		newFndRecord = UserFoundationIdentityNonSocial{
			IsAdmin:         fndProfile.Data.Accounts[0].IsAdmin,
			NetRevenueInETH: fndProfile.Data.Accounts[0].NetRevenueInETH,
			Nfts:            fndProfile.Data.Accounts[0].Nfts,
			Creator:         fndProfile.Data.Accounts[0].Creator,
			Withdrawals:     fndProfile.Data.Accounts[0].Withdrawals,
			DataSource:      FOUNDATION,
		}
	}
	if len(newFndRecord.Nfts) != 0 || len(newFndRecord.Withdrawals) != 0 || newFndRecord.NetRevenueInETH != "0" || newFndRecord.IsAdmin != false {
		result.FoundationNonSocial = &newFndRecord
	}
	ch <- result
}
