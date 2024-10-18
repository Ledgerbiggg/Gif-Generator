package impls

import (
	"fmt"
	"gif_generator/src/config"
	"gif_generator/src/logs"
	"gif_generator/src/services/interfaces"
	"gif_generator/src/utils"
	"github.com/disintegration/imaging"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"os"
)

type GenGif struct {
	log    *logs.ConsoleLogger
	config *config.GConfig
}

func (g *GenGif) GenerateGif() error {
	imgArr := g.config.Image.ImageList
	outGif := &gif.GIF{}

	for _, img := range imgArr {
		inGif, err := utils.ImageDecode(fmt.Sprintf("./images/%s", img.File)) // 使用 img.File 获取文件名
		if err != nil {
			fmt.Println("Error opening image", err)
			continue // 跳过出错的图像
		}

		// 调整图片大小
		inGif = imaging.Resize(inGif, g.config.Image.Width, g.config.Image.Height, imaging.Lanczos)

		// 创建调色板图像
		bounds := inGif.Bounds()
		palettedImg := image.NewPaletted(bounds, palette.Plan9)
		draw.Draw(palettedImg, bounds, inGif, bounds.Min, draw.Src)

		// 添加帧和帧延迟
		outGif.Image = append(outGif.Image, palettedImg)
		outGif.Delay = append(outGif.Delay, img.Delay) // 使用每个图片的延迟
	}

	// 打开输出文件
	f, _ := os.OpenFile(g.config.Output.GifName, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	// 编码并保存 GIF
	err := gif.EncodeAll(f, outGif)
	if err != nil {
		g.log.Error(err.Error()) // 记录错误
		return err
	}

	return nil // GIF 生成成功
}

func NewGenGif(p Params) interfaces.IGenGif {
	g := &GenGif{}
	g.config = p.Config
	g.log = p.Log
	return g
}
