package commands

import (
	"log"
	"os"
	"path"
	"strings"

	t "github.com/eptaccio/willtg/types"
	"github.com/fogleman/gg"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Will command
func Will(m *tb.Message, b *tb.Bot) {

	imageConfig := &t.ImageConfig{
		Context: t.ImageContext{
			Height: 898,
			Width:  559,
		},
		FontSize:  30,
		ImagePath: "files/bc.jpg",
		Color:     "#FFFFFF",
	}

	im, err := gg.LoadImage(imageConfig.ImagePath)
	if err != nil {
		log.Fatal(err)
	}

	textOnImage := strings.Replace(m.Text, "/will", "", 1)

	dc := gg.NewContext(imageConfig.Context.Width, imageConfig.Context.Height)
	dc.LoadFontFace("files/font.ttf", imageConfig.FontSize)
	dc.DrawImage(im, 0, 0)
	dc.SetHexColor(imageConfig.Color)

	dc.DrawStringWrapped(
		textOnImage,
		float64(imageConfig.Context.Width)/2,
		float64(imageConfig.Context.Height)/2,
		0.5, 1, 500, 1.5, gg.AlignCenter)

	fileName := string(m.ID) + ".png"
	filePath := path.Join(os.TempDir(), fileName)
	dc.SavePNG(filePath)

	photo := &tb.Photo{File: tb.FromDisk(filePath)}
	b.SendAlbum(m.Chat, tb.Album{photo})
}