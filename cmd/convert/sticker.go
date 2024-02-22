package cmd

import (
	x "mywabot/system"
)

func init() {
	x.NewCmd(&x.ICmd{
		Name: "(s|stiker)",
		Cmd:  []string{"sticker"},
		Tags: "convert",
		IsMedia: true,
		Prefix: true,
		Exec: func(sock *x.Nc, m *x.IMsg) {
			m.React("⏱️")

			// quoted sticker
			if m.IsQuotedSticker {
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.StickerMessage)
        
        s := x.StickerApi(&x.Sticker{
          File: byte,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
            }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)
				m.React("✅")
			}

      
			// quoted image
			if m.IsQuotedImage {
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.ImageMessage)
        
        s := x.StickerApi(&x.Sticker{
          File: byte,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
            }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)
				m.React("✅")
			}

			// quoted video
			if m.IsQuotedVideo {
				byte, _ := sock.WA.Download(m.Quoted.QuotedMessage.VideoMessage)

        s := x.StickerApi(&x.Sticker{
          File: byte,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
            }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)
				m.React("✅")
			}

			// from video
			if m.IsVideo {
				byte, _ := sock.WA.Download(m.Media)
        
        s := x.StickerApi(&x.Sticker{
          File: byte,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
            }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)
				m.React("✅")
			}

			// from image
			if m.IsImage {
				byte, _ := sock.WA.Download(m.Media)
    
        s := x.StickerApi(&x.Sticker{
          File: byte,
          Tipe: func() x.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return x.IMAGE
            } else {
              return x.VIDEO
            }
          }(),
        }, &x.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Query == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        sock.SendSticker(m.From, s.Build(), *m)
				m.React("✅")
			}

		},
	})
}
