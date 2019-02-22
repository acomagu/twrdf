package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"sort"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/pkg/errors"
)

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
	tl, err := retrieveTweets("")
	if err != nil {
		return err
	}

	return gen(os.Stdout, tl)
}

func retrieveTweets(untilID string) ([]anaconda.Tweet, error) {
	access_token := os.Getenv("ACCESS_TOKEN")
	access_token_secret := os.Getenv("ACCESS_TOKEN_SECRET")
	consumer_key := os.Getenv("CONSUMER_KEY")
	consumer_secret := os.Getenv("CONSUMER_SECRET")
	twitter := anaconda.NewTwitterApiWithCredentials(access_token, access_token_secret, consumer_key, consumer_secret)

	params := url.Values{}

	params.Add("count", "3200")
	if untilID != "" {
		params.Add("max_id", untilID)
	}
	params.Add("exclude_replies", "false")
	params.Add("include_rts", "true")

	tl, err := twitter.GetUserTimeline(params)
	if err != nil {
		return nil, err
	}
	if len(tl) > 0 && tl[0].IdStr == untilID {
		tl = tl[1:]
	}

	fmt.Fprintf(os.Stderr, "retrieved %d tweets\n", len(tl))

	if len(tl) == 0 {
		return tl, nil
	}

	minID := tl[len(tl)-1].IdStr
	ntl, err := retrieveTweets(minID)
	if err != nil {
		return nil, err
	}

	return append(tl, ntl...), nil
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
