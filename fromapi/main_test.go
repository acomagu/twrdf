package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/ChimeraCoder/anaconda"
	"github.com/knakk/rdf"
	"github.com/matryer/is"
)

var tls = `[
  {
    "contributors": null,
    "coordinates": null,
    "created_at": "Wed Jan 30 08:35:35 +0000 2019",
    "display_text_range": [
      0,
      35
    ],
    "entities": {
      "urls": [
        {
          "indices": [
            36,
            59
          ],
          "url": "https://t.co/Ix0NRDaZgR",
          "display_url": "twitter.com/about_hiroppy/…",
          "expanded_url": "https://twitter.com/about_hiroppy/status/1090502272193187840"
        }
      ],
      "hashtags": [],
      "url": {
        "urls": null
      },
      "user_mentions": [],
      "media": null
    },
    "extended_entities": {
      "urls": null,
      "hashtags": null,
      "url": {
        "urls": null
      },
      "user_mentions": null,
      "media": null
    },
    "extended_tweet": {
      "full_text": "",
      "display_text_range": null,
      "entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": null
        },
        "user_mentions": null,
        "media": null
      },
      "extended_entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": null
        },
        "user_mentions": null,
        "media": null
      }
    },
    "favorite_count": 0,
    "favorited": false,
    "filter_level": "",
    "full_text": "型が強い言語のライブラリとか読んでるとたまにこういうコード出てくるよね https://t.co/Ix0NRDaZgR",
    "has_extended_profile": false,
    "id": 1090528940655968300,
    "id_str": "1090528940655968256",
    "in_reply_to_screen_name": "",
    "in_reply_to_status_id": 0,
    "in_reply_to_status_id_str": "",
    "in_reply_to_user_id": 0,
    "in_reply_to_user_id_str": "",
    "is_translation_enabled": false,
    "lang": "ja",
    "place": {
      "attributes": null,
      "bounding_box": {
        "coordinates": null,
        "type": ""
      },
      "contained_within": null,
      "country": "",
      "country_code": "",
      "full_name": "",
      "geometry": {
        "coordinates": null,
        "type": ""
      },
      "id": "",
      "name": "",
      "place_type": "",
      "polylines": null,
      "url": ""
    },
    "quoted_status_id": 1090502272193187800,
    "quoted_status_id_str": "1090502272193187840",
    "quoted_status": {
      "contributors": null,
      "coordinates": null,
      "created_at": "Wed Jan 30 06:49:36 +0000 2019",
      "display_text_range": [
        0,
        0
      ],
      "entities": {
        "urls": [],
        "hashtags": [],
        "url": {
          "urls": null
        },
        "user_mentions": [],
        "media": [
          {
            "id": 1090502266459574300,
            "id_str": "1090502266459574272",
            "media_url": "http://pbs.twimg.com/media/DyI-CJsU0AASEq_.jpg",
            "media_url_https": "https://pbs.twimg.com/media/DyI-CJsU0AASEq_.jpg",
            "url": "https://t.co/OpfLFQyEba",
            "display_url": "pic.twitter.com/OpfLFQyEba",
            "expanded_url": "https://twitter.com/about_hiroppy/status/1090502272193187840/photo/1",
            "sizes": {
              "medium": {
                "w": 1006,
                "h": 500,
                "resize": "fit"
              },
              "thumb": {
                "w": 150,
                "h": 150,
                "resize": "crop"
              },
              "small": {
                "w": 680,
                "h": 338,
                "resize": "fit"
              },
              "large": {
                "w": 1006,
                "h": 500,
                "resize": "fit"
              }
            },
            "source_status_id": 0,
            "source_status_id_str": "",
            "type": "photo",
            "indices": [
              0,
              23
            ],
            "video_info": {
              "aspect_ratio": null,
              "duration_millis": 0,
              "variants": null
            },
            "ext_alt_text": ""
          }
        ]
      },
      "extended_entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": null
        },
        "user_mentions": null,
        "media": [
          {
            "id": 1090502266459574300,
            "id_str": "1090502266459574272",
            "media_url": "http://pbs.twimg.com/media/DyI-CJsU0AASEq_.jpg",
            "media_url_https": "https://pbs.twimg.com/media/DyI-CJsU0AASEq_.jpg",
            "url": "https://t.co/OpfLFQyEba",
            "display_url": "pic.twitter.com/OpfLFQyEba",
            "expanded_url": "https://twitter.com/about_hiroppy/status/1090502272193187840/photo/1",
            "sizes": {
              "medium": {
                "w": 1006,
                "h": 500,
                "resize": "fit"
              },
              "thumb": {
                "w": 150,
                "h": 150,
                "resize": "crop"
              },
              "small": {
                "w": 680,
                "h": 338,
                "resize": "fit"
              },
              "large": {
                "w": 1006,
                "h": 500,
                "resize": "fit"
              }
            },
            "source_status_id": 0,
            "source_status_id_str": "",
            "type": "photo",
            "indices": [
              0,
              23
            ],
            "video_info": {
              "aspect_ratio": null,
              "duration_millis": 0,
              "variants": null
            },
            "ext_alt_text": ""
          }
        ]
      },
      "extended_tweet": {
        "full_text": "",
        "display_text_range": null,
        "entities": {
          "urls": null,
          "hashtags": null,
          "url": {
            "urls": null
          },
          "user_mentions": null,
          "media": null
        },
        "extended_entities": {
          "urls": null,
          "hashtags": null,
          "url": {
            "urls": null
          },
          "user_mentions": null,
          "media": null
        }
      },
      "favorite_count": 28,
      "favorited": false,
      "filter_level": "",
      "full_text": "https://t.co/OpfLFQyEba",
      "has_extended_profile": false,
      "id": 1090502272193187800,
      "id_str": "1090502272193187840",
      "in_reply_to_screen_name": "",
      "in_reply_to_status_id": 0,
      "in_reply_to_status_id_str": "",
      "in_reply_to_user_id": 0,
      "in_reply_to_user_id_str": "",
      "is_translation_enabled": false,
      "lang": "und",
      "place": {
        "attributes": null,
        "bounding_box": {
          "coordinates": null,
          "type": ""
        },
        "contained_within": null,
        "country": "",
        "country_code": "",
        "full_name": "",
        "geometry": {
          "coordinates": null,
          "type": ""
        },
        "id": "",
        "name": "",
        "place_type": "",
        "polylines": null,
        "url": ""
      },
      "quoted_status_id": 0,
      "quoted_status_id_str": "",
      "quoted_status": null,
      "possibly_sensitive": false,
      "possibly_sensitive_appealable": false,
      "retweet_count": 6,
      "retweeted": false,
      "retweeted_status": null,
      "source": "a href=\"https://about.twitter.com/products/tweetdeck\" rel=\"nofollow\"TweetDeck/a",
      "scopes": null,
      "text": "https://t.co/OpfLFQyEba",
      "user": {
        "contributors_enabled": false,
        "created_at": "Sat Feb 13 13:09:07 +0000 2010",
        "default_profile": false,
        "default_profile_image": false,
        "description": "@nodejs core team member. Working on @nodejs, @webpack @stylelint, @babeljs and Node.js Japan User Group. A developer of Dwango and technical advisor of Mercari",
        "email": "",
        "entities": {
          "urls": null,
          "hashtags": null,
          "url": {
            "urls": [
              {
                "indices": [
                  0,
                  23
                ],
                "url": "https://t.co/hIDKhhzpb5",
                "display_url": "hiroppy.me",
                "expanded_url": "https://hiroppy.me"
              }
            ]
          },
          "user_mentions": null,
          "media": null
        },
        "favourites_count": 4691,
        "follow_request_sent": false,
        "followers_count": 1497,
        "following": true,
        "friends_count": 323,
        "geo_enabled": false,
        "has_extended_profile": true,
        "id": 113915322,
        "id_str": "113915322",
        "is_translator": false,
        "is_translation_enabled": false,
        "lang": "en",
        "listed_count": 95,
        "location": "Minato-ku/Tokyo/Japan",
        "name": "hiroppy😶",
        "notifications": false,
        "profile_background_color": "000000",
        "profile_background_image_url": "http://abs.twimg.com/images/themes/theme4/bg.gif",
        "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme4/bg.gif",
        "profile_background_tile": false,
        "profile_banner_url": "https://pbs.twimg.com/profile_banners/113915322/1508204313",
        "profile_image_url": "http://pbs.twimg.com/profile_images/1721255496/101010_normal.GIF",
        "profile_image_url_https": "https://pbs.twimg.com/profile_images/1721255496/101010_normal.GIF",
        "profile_link_color": "3498DB",
        "profile_sidebar_border_color": "FFFFFF",
        "profile_sidebar_fill_color": "95E8EC",
        "profile_text_color": "3C3940",
        "profile_use_background_image": true,
        "protected": false,
        "screen_name": "about_hiroppy",
        "show_all_inline_media": false,
        "status": null,
        "statuses_count": 98251,
        "time_zone": "",
        "url": "https://t.co/hIDKhhzpb5",
        "utc_offset": 0,
        "verified": false,
        "withheld_in_countries": null,
        "withheld_scope": ""
      },
      "withheld_copyright": false,
      "withheld_in_countries": null,
      "withheld_scope": ""
    },
    "possibly_sensitive": false,
    "possibly_sensitive_appealable": false,
    "retweet_count": 0,
    "retweeted": false,
    "retweeted_status": null,
    "source": "a href=\"https://mobile.twitter.com\" rel=\"nofollow\"Twitter Web App/a",
    "scopes": null,
    "text": "型が強い言語のライブラ��",
    "user": {
      "contributors_enabled": false,
      "created_at": "Thu Mar 13 23:14:08 +0000 2014",
      "default_profile": false,
      "default_profile_image": false,
      "description": "よねです。 はてな/インスタ/Swarm/Facebook/GitHub: acomagu 福島東高-会津大学\n(自動ブロックツール使ってるので知り合いでブロックされてる方いたら連絡ください)",
      "email": "",
      "entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": [
            {
              "indices": [
                0,
                23
              ],
              "url": "https://t.co/ovFH6c8VRW",
              "display_url": "acomagu.me",
              "expanded_url": "https://acomagu.me"
            }
          ]
        },
        "user_mentions": null,
        "media": null
      },
      "favourites_count": 5001,
      "follow_request_sent": false,
      "followers_count": 492,
      "following": false,
      "friends_count": 935,
      "geo_enabled": true,
      "has_extended_profile": false,
      "id": 2388044550,
      "id_str": "2388044550",
      "is_translator": false,
      "is_translation_enabled": false,
      "lang": "ja",
      "listed_count": 16,
      "location": "",
      "name": "acomagu",
      "notifications": false,
      "profile_background_color": "000000",
      "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
      "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
      "profile_background_tile": false,
      "profile_banner_url": "https://pbs.twimg.com/profile_banners/2388044550/1467219608",
      "profile_image_url": "http://pbs.twimg.com/profile_images/750397258214936576/XVInP_lz_normal.jpg",
      "profile_image_url_https": "https://pbs.twimg.com/profile_images/750397258214936576/XVInP_lz_normal.jpg",
      "profile_link_color": "FF691F",
      "profile_sidebar_border_color": "000000",
      "profile_sidebar_fill_color": "000000",
      "profile_text_color": "000000",
      "profile_use_background_image": false,
      "protected": false,
      "screen_name": "acomagu",
      "show_all_inline_media": false,
      "status": null,
      "statuses_count": 12513,
      "time_zone": "",
      "url": "https://t.co/ovFH6c8VRW",
      "utc_offset": 0,
      "verified": false,
      "withheld_in_countries": null,
      "withheld_scope": ""
    },
    "withheld_copyright": false,
    "withheld_in_countries": null,
    "withheld_scope": ""
  },
  {
    "contributors": null,
    "coordinates": null,
    "created_at": "Wed Jan 30 08:24:29 +0000 2019",
    "display_text_range": [
      0,
      27
    ],
    "entities": {
      "urls": [],
      "hashtags": [],
      "url": {
        "urls": null
      },
      "user_mentions": [],
      "media": null
    },
    "extended_entities": {
      "urls": null,
      "hashtags": null,
      "url": {
        "urls": null
      },
      "user_mentions": null,
      "media": null
    },
    "extended_tweet": {
      "full_text": "",
      "display_text_range": null,
      "entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": null
        },
        "user_mentions": null,
        "media": null
      },
      "extended_entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": null
        },
        "user_mentions": null,
        "media": null
      }
    },
    "favorite_count": 0,
    "favorited": false,
    "filter_level": "",
    "full_text": "スクラムマスターとふくやマスターって似過ぎでは...?",
    "has_extended_profile": false,
    "id": 1090526147488579600,
    "id_str": "1090526147488579584",
    "in_reply_to_screen_name": "",
    "in_reply_to_status_id": 0,
    "in_reply_to_status_id_str": "",
    "in_reply_to_user_id": 0,
    "in_reply_to_user_id_str": "",
    "is_translation_enabled": false,
    "lang": "ja",
    "place": {
      "attributes": null,
      "bounding_box": {
        "coordinates": null,
        "type": ""
      },
      "contained_within": null,
      "country": "",
      "country_code": "",
      "full_name": "",
      "geometry": {
        "coordinates": null,
        "type": ""
      },
      "id": "",
      "name": "",
      "place_type": "",
      "polylines": null,
      "url": ""
    },
    "quoted_status_id": 0,
    "quoted_status_id_str": "",
    "quoted_status": null,
    "possibly_sensitive": false,
    "possibly_sensitive_appealable": false,
    "retweet_count": 0,
    "retweeted": false,
    "retweeted_status": null,
    "source": "a href=\"https://mobile.twitter.com\" rel=\"nofollow\"Twitter Web App/a",
    "scopes": null,
    "text": "スクラムマスターと",
    "user": {
      "contributors_enabled": false,
      "created_at": "Thu Mar 13 23:14:08 +0000 2014",
      "default_profile": false,
      "default_profile_image": false,
      "description": "よねです。 はてな/インスタ/Swarm/Facebook/GitHub: acomagu 福島東高-会津大学\n(自動ブロックツール使ってるので知り合いでブロックされてる方いたら連絡ください)",
      "email": "",
      "entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": [
            {
              "indices": [
                0,
                23
              ],
              "url": "https://t.co/ovFH6c8VRW",
              "display_url": "acomagu.me",
              "expanded_url": "https://acomagu.me"
            }
          ]
        },
        "user_mentions": null,
        "media": null
      },
      "favourites_count": 5001,
      "follow_request_sent": false,
      "followers_count": 492,
      "following": false,
      "friends_count": 935,
      "geo_enabled": true,
      "has_extended_profile": false,
      "id": 2388044550,
      "id_str": "2388044550",
      "is_translator": false,
      "is_translation_enabled": false,
      "lang": "ja",
      "listed_count": 16,
      "location": "",
      "name": "acomagu",
      "notifications": false,
      "profile_background_color": "000000",
      "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
      "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
      "profile_background_tile": false,
      "profile_banner_url": "https://pbs.twimg.com/profile_banners/2388044550/1467219608",
      "profile_image_url": "http://pbs.twimg.com/profile_images/750397258214936576/XVInP_lz_normal.jpg",
      "profile_image_url_https": "https://pbs.twimg.com/profile_images/750397258214936576/XVInP_lz_normal.jpg",
      "profile_link_color": "FF691F",
      "profile_sidebar_border_color": "000000",
      "profile_sidebar_fill_color": "000000",
      "profile_text_color": "000000",
      "profile_use_background_image": false,
      "protected": false,
      "screen_name": "acomagu",
      "show_all_inline_media": false,
      "status": null,
      "statuses_count": 12513,
      "time_zone": "",
      "url": "https://t.co/ovFH6c8VRW",
      "utc_offset": 0,
      "verified": false,
      "withheld_in_countries": null,
      "withheld_scope": ""
    },
    "withheld_copyright": false,
    "withheld_in_countries": null,
    "withheld_scope": ""
  },
  {
    "contributors": null,
    "coordinates": null,
    "created_at": "Sat Jan 26 06:34:03 +0000 2019",
    "display_text_range": [
      0,
      104
    ],
    "entities": {
      "urls": [],
      "hashtags": [],
      "url": {
        "urls": null
      },
      "user_mentions": [
        {
          "name": "set0gut1",
          "indices": [
            3,
            12
          ],
          "screen_name": "set0gut1",
          "id": 991169559787851800,
          "id_str": "991169559787851776"
        }
      ],
      "media": null
    },
    "extended_entities": {
      "urls": null,
      "hashtags": null,
      "url": {
        "urls": null
      },
      "user_mentions": null,
      "media": null
    },
    "extended_tweet": {
      "full_text": "",
      "display_text_range": null,
      "entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": null
        },
        "user_mentions": null,
        "media": null
      },
      "extended_entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": null
        },
        "user_mentions": null,
        "media": null
      }
    },
    "favorite_count": 0,
    "favorited": false,
    "filter_level": "",
    "full_text": "RT @set0gut1: プログラミング教育で真に実施すべきなのは音ゲー。\n幼少期に音ゲーに習熟することで、高校生以降に音ゲーの練習に要する時間を短縮でき、プログラミングにより多くの時間を使うことができる。",
    "has_extended_profile": false,
    "id": 1089048805138288600,
    "id_str": "1089048805138288642",
    "in_reply_to_screen_name": "",
    "in_reply_to_status_id": 0,
    "in_reply_to_status_id_str": "",
    "in_reply_to_user_id": 0,
    "in_reply_to_user_id_str": "",
    "is_translation_enabled": false,
    "lang": "ja",
    "place": {
      "attributes": null,
      "bounding_box": {
        "coordinates": null,
        "type": ""
      },
      "contained_within": null,
      "country": "",
      "country_code": "",
      "full_name": "",
      "geometry": {
        "coordinates": null,
        "type": ""
      },
      "id": "",
      "name": "",
      "place_type": "",
      "polylines": null,
      "url": ""
    },
    "quoted_status_id": 0,
    "quoted_status_id_str": "",
    "quoted_status": null,
    "possibly_sensitive": false,
    "possibly_sensitive_appealable": false,
    "retweet_count": 50,
    "retweeted": true,
    "retweeted_status": {
      "contributors": null,
      "coordinates": null,
      "created_at": "Fri Jan 25 02:32:00 +0000 2019",
      "display_text_range": [
        0,
        90
      ],
      "entities": {
        "urls": [],
        "hashtags": [],
        "url": {
          "urls": null
        },
        "user_mentions": [],
        "media": null
      },
      "extended_entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": null
        },
        "user_mentions": null,
        "media": null
      },
      "extended_tweet": {
        "full_text": "",
        "display_text_range": null,
        "entities": {
          "urls": null,
          "hashtags": null,
          "url": {
            "urls": null
          },
          "user_mentions": null,
          "media": null
        },
        "extended_entities": {
          "urls": null,
          "hashtags": null,
          "url": {
            "urls": null
          },
          "user_mentions": null,
          "media": null
        }
      },
      "favorite_count": 152,
      "favorited": false,
      "filter_level": "",
      "full_text": "プログラミング教育で真に実施すべきなのは音ゲー。\n幼少期に音ゲーに習熟することで、高校生以降に音ゲーの練習に要する時間を短縮でき、プログラミングにより多くの時間を使うことができる。",
      "has_extended_profile": false,
      "id": 1088625506054881300,
      "id_str": "1088625506054881280",
      "in_reply_to_screen_name": "",
      "in_reply_to_status_id": 0,
      "in_reply_to_status_id_str": "",
      "in_reply_to_user_id": 0,
      "in_reply_to_user_id_str": "",
      "is_translation_enabled": false,
      "lang": "ja",
      "place": {
        "attributes": null,
        "bounding_box": {
          "coordinates": null,
          "type": ""
        },
        "contained_within": null,
        "country": "",
        "country_code": "",
        "full_name": "",
        "geometry": {
          "coordinates": null,
          "type": ""
        },
        "id": "",
        "name": "",
        "place_type": "",
        "polylines": null,
        "url": ""
      },
      "quoted_status_id": 0,
      "quoted_status_id_str": "",
      "quoted_status": null,
      "possibly_sensitive": false,
      "possibly_sensitive_appealable": false,
      "retweet_count": 50,
      "retweeted": true,
      "retweeted_status": null,
      "source": "a href=\"http://twitter.com\" rel=\"nofollow\"Twitter Web Client/a",
      "scopes": null,
      "text": "プログラミング教育で真に実施すべきなのは音ゲー。\n幼少期に音��",
      "user": {
        "contributors_enabled": false,
        "created_at": "Tue May 01 04:16:52 +0000 2018",
        "default_profile": false,
        "default_profile_image": false,
        "description": "エモさ駆動",
        "email": "",
        "entities": {
          "urls": null,
          "hashtags": null,
          "url": {
            "urls": [
              {
                "indices": [
                  0,
                  23
                ],
                "url": "https://t.co/rIqDnF8Cw9",
                "display_url": "set0gut1.com",
                "expanded_url": "https://set0gut1.com/"
              }
            ]
          },
          "user_mentions": null,
          "media": null
        },
        "favourites_count": 9822,
        "follow_request_sent": false,
        "followers_count": 3799,
        "following": false,
        "friends_count": 4042,
        "geo_enabled": true,
        "has_extended_profile": false,
        "id": 991169559787851800,
        "id_str": "991169559787851776",
        "is_translator": false,
        "is_translation_enabled": false,
        "lang": "ja",
        "listed_count": 24,
        "location": "Tokyo",
        "name": "set0gut1",
        "notifications": false,
        "profile_background_color": "000000",
        "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
        "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
        "profile_background_tile": false,
        "profile_banner_url": "https://pbs.twimg.com/profile_banners/991169559787851776/1530894597",
        "profile_image_url": "http://pbs.twimg.com/profile_images/1059170149389918208/g2qYxInB_normal.jpg",
        "profile_image_url_https": "https://pbs.twimg.com/profile_images/1059170149389918208/g2qYxInB_normal.jpg",
        "profile_link_color": "1B95E0",
        "profile_sidebar_border_color": "000000",
        "profile_sidebar_fill_color": "000000",
        "profile_text_color": "000000",
        "profile_use_background_image": false,
        "protected": false,
        "screen_name": "set0gut1",
        "show_all_inline_media": false,
        "status": null,
        "statuses_count": 589,
        "time_zone": "",
        "url": "https://t.co/rIqDnF8Cw9",
        "utc_offset": 0,
        "verified": false,
        "withheld_in_countries": null,
        "withheld_scope": ""
      },
      "withheld_copyright": false,
      "withheld_in_countries": null,
      "withheld_scope": ""
    },
    "source": "a href=\"https://mobile.twitter.com\" rel=\"nofollow\"Twitter Web App/a",
    "scopes": null,
    "text": "RT @set0gut1: プログラミング教育で真に実施すべきなのは音ゲー。\n幼少期に音��",
    "user": {
      "contributors_enabled": false,
      "created_at": "Thu Mar 13 23:14:08 +0000 2014",
      "default_profile": false,
      "default_profile_image": false,
      "description": "よねです。 はてな/インスタ/Swarm/Facebook/GitHub: acomagu 福島東高-会津大学\n(自動ブロックツール使ってるので知り合いでブロックされてる方いたら連絡ください)",
      "email": "",
      "entities": {
        "urls": null,
        "hashtags": null,
        "url": {
          "urls": [
            {
              "indices": [
                0,
                23
              ],
              "url": "https://t.co/ovFH6c8VRW",
              "display_url": "acomagu.me",
              "expanded_url": "https://acomagu.me"
            }
          ]
        },
        "user_mentions": null,
        "media": null
      },
      "favourites_count": 5001,
      "follow_request_sent": false,
      "followers_count": 492,
      "following": false,
      "friends_count": 935,
      "geo_enabled": true,
      "has_extended_profile": false,
      "id": 2388044550,
      "id_str": "2388044550",
      "is_translator": false,
      "is_translation_enabled": false,
      "lang": "ja",
      "listed_count": 16,
      "location": "",
      "name": "acomagu",
      "notifications": false,
      "profile_background_color": "000000",
      "profile_background_image_url": "http://abs.twimg.com/images/themes/theme1/bg.png",
      "profile_background_image_url_https": "https://abs.twimg.com/images/themes/theme1/bg.png",
      "profile_background_tile": false,
      "profile_banner_url": "https://pbs.twimg.com/profile_banners/2388044550/1467219608",
      "profile_image_url": "http://pbs.twimg.com/profile_images/750397258214936576/XVInP_lz_normal.jpg",
      "profile_image_url_https": "https://pbs.twimg.com/profile_images/750397258214936576/XVInP_lz_normal.jpg",
      "profile_link_color": "FF691F",
      "profile_sidebar_border_color": "000000",
      "profile_sidebar_fill_color": "000000",
      "profile_text_color": "000000",
      "profile_use_background_image": false,
      "protected": false,
      "screen_name": "acomagu",
      "show_all_inline_media": false,
      "status": null,
      "statuses_count": 12513,
      "time_zone": "",
      "url": "https://t.co/ovFH6c8VRW",
      "utc_offset": 0,
      "verified": false,
      "withheld_in_countries": null,
      "withheld_scope": ""
    },
    "withheld_copyright": false,
    "withheld_in_countries": null,
    "withheld_scope": ""
  }
]`

type tCouple struct {
	predicate, object string
	objectType        rdf.TermType
}

var expectedTriples = map[string][]tCouple{
	"https://twitter.com/acomagu/status/1090526147488579584": {
		{"rdf:type", "sioc:Post", rdf.TermIRI},
		{"dc:created", "2019-01-30T08:24:29Z", rdf.TermLiteral},
		{"sioc:content", "スクラムマスターとふくやマスターって似過ぎでは...?", rdf.TermLiteral},
		{"sioc:has_creator", "https://twitter.com/acomagu", rdf.TermIRI},
	},
	"https://twitter.com/set0gut1/status/1088625506054881280": {
		{"rdf:type", "sioc:Post", rdf.TermIRI},
		{"dc:created", "2019-01-25T02:32:00Z", rdf.TermLiteral},
		{"sioc:content", "プログラミング教育で真に実施すべきなのは音ゲー。\n幼少期に音ゲーに習熟することで、高校生以降に音ゲーの練習に要する時間を短縮でき、プログラミングにより多くの時間を使うことができる。", rdf.TermLiteral},
		{"sioc:has_creator", "https://twitter.com/set0gut1", rdf.TermIRI},
		{"sioc:shared_by", "https://twitter.com/acomagu", rdf.TermIRI},
		{"sioc:sibling", "https://twitter.com/acomagu/status/1089048805138288642", rdf.TermIRI},
	},
	"https://twitter.com/acomagu/status/1089048805138288642": {
		{"rdf:type", "sioc:Post", rdf.TermIRI},
		{"dc:created", "2019-01-26T06:34:03Z", rdf.TermLiteral},
		{"sioc:content", "RT @set0gut1: プログラミング教育で真に実施すべきなのは音ゲー。\n幼少期に音ゲーに習熟することで、高校生以降に音ゲーの練習に要する時間を短縮でき、プログラミングにより多くの時間を使うことができる。", rdf.TermLiteral},
		{"sioc:has_creator", "https://twitter.com/acomagu", rdf.TermIRI},
		{"sioc:sibling", "https://twitter.com/set0gut1/status/1088625506054881280", rdf.TermIRI},
	},
	"https://twitter.com/acomagu/status/1090528940655968256": {
		{"rdf:type", "sioc:Post", rdf.TermIRI},
		{"dc:created", "2019-01-30T08:35:35Z", rdf.TermLiteral},
		{"sioc:content", "型が強い言語のライブラリとか読んでるとたまにこういうコード出てくるよね https://t.co/Ix0NRDaZgR", rdf.TermLiteral},
		{"sioc:has_creator", "https://twitter.com/acomagu", rdf.TermIRI},
		{"sioc:links_to", "https://twitter.com/about_hiroppy/status/1090502272193187840", rdf.TermIRI},
	},
	"https://twitter.com/about_hiroppy/status/1090502272193187840": {
		{"rdf:type", "sioc:Post", rdf.TermIRI},
		{"dc:created", "2019-01-30T06:49:36Z", rdf.TermLiteral},
		{"sioc:has_creator", "https://twitter.com/about_hiroppy", rdf.TermIRI},
		{"sioc:content", "https://t.co/OpfLFQyEba", rdf.TermLiteral},
		{"sioc:attachment", "https://pbs.twimg.com/media/DyI-CJsU0AASEq_.jpg", rdf.TermIRI},
	},
}

func Test(t *testing.T) {
	is := is.New(t)

	var tl []anaconda.Tweet
	if err := json.Unmarshal([]byte(tls), &tl); err != nil {
		panic(err)
	}

	prefixReplacer := strings.NewReplacer(
		"rdf:", "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
		"dc:", "http://purl.org/dc/terms/",
		"sioc:", "http://rdfs.org/sioc/ns#",
	)

	buf := bytes.NewBuffer(nil)
	is.NoErr(gen(buf, tl))
	triples, err := rdf.NewTripleDecoder(buf, rdf.Turtle).DecodeAll()
	is.NoErr(err)
	for subject, couples := range expectedTriples {
		for _, couple := range couples {
			ti := -1
			for i, triple := range triples {
				if triple.Subj.Type() == rdf.TermIRI &&
					triple.Subj.String() == subject &&
					triple.Pred.String() == prefixReplacer.Replace(couple.predicate) &&
					triple.Obj.Type() == couple.objectType &&
					triple.Obj.String() == prefixReplacer.Replace(couple.object) {

					ti = i
					break
				}
			}

			if ti == -1 {
				t.Errorf("could not found triple: %s %s %s", subject, couple.predicate, couple.object)
			} else {
				triples = append(triples[:ti], triples[ti+1:]...)
			}
		}
	}

	for _, triple := range triples {
		t.Errorf("extra triple: %s", triple)
	}
}
