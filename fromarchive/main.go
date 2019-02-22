package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/pkg/errors"
)

type EntityMedia struct {
	DisplayURL        string   `json:"display_url"`
	ExpandedURL       string   `json:"expanded_url"`
	ID                string   `json:"id"`
	IDStr             string   `json:"id_str"`
	Indices           []string `json:"indices"`
	MediaURL          string   `json:"media_url"`
	MediaURLHTTPS     string   `json:"media_url_https"`
	SourceStatusID    string   `json:"source_status_id"`
	SourceStatusIDStr string   `json:"source_status_id_str"`
	Type              string   `json:"type"`
	URL               string   `json:"url"`
}

type Entities struct {
	Media []EntityMedia `json:"media"`
	URLs  []struct {
		DisplayURL  string   `json:"display_url"`
		ExpandedURL string   `json:"expanded_url"`
		Indices     []string `json:"indices"`
		URL         string   `json:"url"`
	} `json:"urls"`
	UserMentions []struct {
		ID         string   `json:"id"`
		IDStr      string   `json:"id_str"`
		Indices    []string `json:"indices"`
		Name       string   `json:"name"`
		ScreenName string   `json:"screen_name"`
	} `json:"user_mentions"`
}

type ArchiveTweet struct {
	CreatedAt            string   `json:"created_at"`
	Entities             Entities `json:"entities"`
	FullText             string   `json:"full_text"`
	ID                   string   `json:"id"`
	IDStr                string   `json:"id_str"`
	InReplyToScreenName  string   `json:"in_reply_to_screen_name"`
	InReplyToStatusID    string   `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string   `json:"in_reply_to_status_id_str"`
	InReplyToUserID      string   `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string   `json:"in_reply_to_user_id_str"`
	Lang                 string   `json:"lang"`
	PossiblySensitive    bool     `json:"possibly_sensitive"`
	Source               string   `json:"source"`
}

type ArchiveAccount struct {
	Account struct {
		AccountDisplayName string `json:"accountDisplayName"`
		AccountID          string `json:"accountId"`
		CreatedAt          string `json:"createdAt"`
		CreatedVia         string `json:"createdVia"`
		Email              string `json:"email"`
		PhoneNumber        string `json:"phoneNumber"`
		Username           string `json:"username"`
	} `json:"account"`
}

type couple struct {
	predicate string
	object    string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	tweetFile, err := os.Open("tweet.js")
	if err != nil {
		return err
	}
	tweetR := bufio.NewReader(tweetFile)
	_, err = tweetR.ReadString('=')
	if err != nil {
		return err
	}

	accountFile, err := os.Open("account.js")
	if err != nil {
		return err
	}
	accountR := bufio.NewReader(accountFile)
	_, err = accountR.ReadString('=')
	if err != nil {
		return err
	}

	tl, err := retrieveTweets(tweetR, accountR)
	if err != nil {
		return err
	}

	// return json.NewEncoder(os.Stdout).Encode(tl)
	return gen(os.Stdout, tl)
}

func retrieveTweets(tweetR, accountR io.Reader) ([]anaconda.Tweet, error) {
	var tweets []ArchiveTweet
	if err := json.NewDecoder(tweetR).Decode(&tweets); err != nil {
		return nil, err
	}

	var accounts []ArchiveAccount
	if err := json.NewDecoder(accountR).Decode(&accounts); err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, errors.New("no account in account.js")
	}
	account := accounts[0]

	var rtIDStrs []string
	var rtIDs []int64
	for _, t := range tweets {
		if strings.HasPrefix(t.FullText, "RT") {
			rtIDStrs = append(rtIDStrs, t.IDStr)

			id, err := strconv.ParseInt(t.IDStr, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "could not parse the ID as int64: %s", t.IDStr)
			}

			rtIDs = append(rtIDs, id)
		}
	}

	twitter := anaconda.NewTwitterApiWithCredentials("2388044550-JMHmB3k6E2cLobfDk8w1zUqMrGQa4F7Xf4qIxDi", "07GSEPHnPK6VzaOLjgqvK6wXg8P0poWnXz9BGtFIL40Z2", "N0WZZn6D9MgyeM8Lf4pbrAYsC", "H2eFU961e4zLBGQ6T77XVpTeJ8nSklP78cncLPmaLEoVjipwtT")

	params := url.Values{}
	params.Add("include_entities", "true")

	var retweets []anaconda.Tweet
	for cursor := 0; cursor <= len(rtIDs); cursor += 100 {
		rts, err := twitter.GetTweetsLookupByIds(rtIDs[cursor:min(cursor+100, len(rtIDs))], params)
		if err != nil {
			return nil, err
		}

		fmt.Fprintf(os.Stderr, "retrieved %d retweets\n", len(retweets))

		retweets = append(retweets, rts...)

		rtmap := make(map[string]struct{})
		for _, rt := range rts {
			rtmap[rt.IdStr] = struct{}{}
		}
		for _, id := range rtIDStrs[cursor:min(cursor+100, len(rtIDStrs))] {
			if _, ok := rtmap[id]; !ok {
				fmt.Fprintf(os.Stderr, "warn: could not fetch the retweet which the ID is %s\n", id)
			}
		}
	}

	retweetsByID := make(map[string]anaconda.Tweet)
	for _, t := range retweets {
		retweetsByID[t.IdStr] = t
	}

	tl := make([]anaconda.Tweet, 0, len(tweets))
	for _, t := range tweets {
		if rt, ok := retweetsByID[t.IDStr]; ok {
			tl = append(tl, rt)
			continue
		}

		conv, err := convTweet(t, account)
		if err != nil {
			return nil, err
		}

		tl = append(tl, conv)
	}

	return tl, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func convTweet(t ArchiveTweet, account ArchiveAccount) (anaconda.Tweet, error) {
	var media []anaconda.EntityMedia
	for _, m := range t.Entities.Media {
		id, err := convID(m.ID)
		if err != nil {
			return anaconda.Tweet{}, err
		}

		indices, err := convIndices(m.Indices)
		if err != nil {
			return anaconda.Tweet{}, err
		}

		sourceStatusID, err := convID(m.SourceStatusID)
		if err != nil {
			return anaconda.Tweet{}, err
		}

		media = append(media, anaconda.EntityMedia{
			Display_url:          m.DisplayURL,
			Expanded_url:         m.ExpandedURL,
			Id:                   id,
			Id_str:               m.IDStr,
			Indices:              indices,
			Media_url:            m.MediaURL,
			Media_url_https:      m.MediaURLHTTPS,
			Source_status_id:     sourceStatusID,
			Source_status_id_str: m.SourceStatusIDStr,
			Type:                 m.Type,
			Url:                  m.URL,
		})
	}

	var urls []struct {
		Indices      []int  `json:"indices"`
		Url          string `json:"url"`
		Display_url  string `json:"display_url"`
		Expanded_url string `json:"expanded_url"`
	}
	for _, url := range t.Entities.URLs {
		indices, err := convIndices(url.Indices)
		if err != nil {
			return anaconda.Tweet{}, err
		}

		urls = append(urls, struct {
			Indices      []int  `json:"indices"`
			Url          string `json:"url"`
			Display_url  string `json:"display_url"`
			Expanded_url string `json:"expanded_url"`
		}{
			Indices:      indices,
			Url:          url.URL,
			Display_url:  url.DisplayURL,
			Expanded_url: url.ExpandedURL,
		})
	}

	var userMentions []struct {
		Name        string `json:"name"`
		Indices     []int  `json:"indices"`
		Screen_name string `json:"screen_name"`
		Id          int64  `json:"id"`
		Id_str      string `json:"id_str"`
	}
	for _, m := range t.Entities.UserMentions {
		id, err := convID(m.ID)
		if err != nil {
			return anaconda.Tweet{}, err
		}

		indices, err := convIndices(m.Indices)
		if err != nil {
			return anaconda.Tweet{}, err
		}

		userMentions = append(userMentions, struct {
			Name        string `json:"name"`
			Indices     []int  `json:"indices"`
			Screen_name string `json:"screen_name"`
			Id          int64  `json:"id"`
			Id_str      string `json:"id_str"`
		}{
			Name:        m.Name,
			Indices:     indices,
			Screen_name: m.ScreenName,
			Id:          id,
			Id_str:      m.IDStr,
		})
	}

	id, err := convID(t.ID)
	if err != nil {
		return anaconda.Tweet{}, err
	}
	inReplyToStatusID, err := convID(t.InReplyToStatusID)
	if err != nil {
		return anaconda.Tweet{}, err
	}
	inReplyToUserID, err := convID(t.InReplyToUserID)
	if err != nil {
		return anaconda.Tweet{}, err
	}
	userID, err := convID(account.Account.AccountID)
	if err != nil {
		return anaconda.Tweet{}, err
	}

	return anaconda.Tweet{
		CreatedAt: t.CreatedAt,
		Entities: anaconda.Entities{
			Media:         media,
			Urls:          urls,
			User_mentions: userMentions,
		},
		FullText:            t.FullText,
		Id:                  id,
		IdStr:               t.IDStr,
		InReplyToScreenName: t.InReplyToScreenName,
		InReplyToStatusID:   inReplyToStatusID,
		InReplyToUserID:     inReplyToUserID,
		Lang:                t.Lang,
		PossiblySensitive:   t.PossiblySensitive,
		Source:              t.Source,
		User: anaconda.User{
			CreatedAt:  account.Account.CreatedAt,
			Email:      account.Account.Email,
			Id:         userID,
			IdStr:      account.Account.AccountID,
			Name:       account.Account.AccountDisplayName,
			ScreenName: account.Account.Username,
		},
	}, nil
}

func convID(orig string) (int64, error) {
	if orig == "" {
		return 0, nil
	}
	return strconv.ParseInt(orig, 10, 64)
}

func convIndices(orig []string) ([]int, error) {
	var art []int
	for _, v := range orig {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		art = append(art, n)
	}

	return art, nil
}

func gen(w io.Writer, tl []anaconda.Tweet) error {
	fmt.Fprintln(w, "@prefix dc: <http://purl.org/dc/terms/>.")
	fmt.Fprintln(w, "@prefix sioc: <http://rdfs.org/sioc/ns#>.")
	fmt.Fprintln(w)

	triples := make(map[string][]couple)
	for _, t := range tl {
		if err := describe(triples, &t, nil); err != nil {
			return errors.Wrapf(err, "could not describe the Tweet %#+v", t)
		}
	}

	subjects := make([]string, 0, len(triples))
	for subject := range triples {
		subjects = append(subjects, subject)
	}
	sort.Strings(subjects)
	for _, subject := range subjects {
		fmt.Fprintf(w, "%s a sioc:Post;\n", subject)
		couples := triples[subject]
		for i, couple := range couples {
			fmt.Fprintf(w, "\t%s %s", couple.predicate, couple.object)
			if i < len(couples)-1 { // if not last
				fmt.Fprintln(w, ";")
			}
		}
		fmt.Fprintln(w, ".")
		fmt.Fprintln(w)
	}

	return nil
}

type option struct {
	Retweet     string
	RetweetedBy string
}

func describe(triples map[string][]couple, t *anaconda.Tweet, opt *option) error {
	var predicate, object string
	subject := fmt.Sprintf("<https://twitter.com/%s/status/%s>", t.User.ScreenName, t.IdStr)

	predicate = "dc:created"
	createdAt, err := t.CreatedAtTime()
	if err != nil {
		return err
	}
	object = quote(createdAt.Format(time.RFC3339))
	triples[subject] = append(triples[subject], couple{predicate, object})

	predicate = "sioc:has_creator"
	object = fmt.Sprintf("<https://twitter.com/%s>", t.User.ScreenName)
	triples[subject] = append(triples[subject], couple{predicate, object})

	predicate = "sioc:content"
	object = quote(t.FullText)
	triples[subject] = append(triples[subject], couple{predicate, object})

	for _, media := range t.Entities.Media {
		predicate = "sioc:attachment"
		object = fmt.Sprintf("<%s>", media.Media_url_https)
		triples[subject] = append(triples[subject], couple{predicate, object})
	}

	if opt != nil && opt.RetweetedBy != "" && opt.Retweet != "" {
		predicate = "sioc:sibling"
		object = opt.Retweet
		triples[subject] = append(triples[subject], couple{predicate, object})

		predicate = "sioc:shared_by"
		object = opt.RetweetedBy
		triples[subject] = append(triples[subject], couple{predicate, object})
	}

	if t.RetweetedStatus != nil {
		predicate = "sioc:sibling"
		object = fmt.Sprintf("<https://twitter.com/%s/status/%s>", t.RetweetedStatus.User.ScreenName, t.RetweetedStatus.IdStr)
		triples[subject] = append(triples[subject], couple{predicate, object})

		describe(triples, t.RetweetedStatus, &option{
			Retweet:     subject,
			RetweetedBy: fmt.Sprintf("<https://twitter.com/%s>", t.User.ScreenName),
		})
	}

	if t.QuotedStatus != nil {
		predicate = "sioc:links_to"
		object = fmt.Sprintf("<https://twitter.com/%s/status/%s>", t.QuotedStatus.User.ScreenName, t.QuotedStatus.IdStr)
		triples[subject] = append(triples[subject], couple{predicate, object})

		describe(triples, t.QuotedStatus, nil)
	}

	return nil
}

func quote(s string) string {
	return fmt.Sprintf("%q", s)
}
