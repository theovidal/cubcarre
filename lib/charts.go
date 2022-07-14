package lib

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/imgg"

	"github.com/theovidal/cubcarre/db"
)

func MakePlot(stats []db.Time) chart.ScatterChart {
	datasets := make(map[string][]chart.EPoint)
	for _, stat := range stats {
		if stat.DNF {
			continue
		}
		if _, found := datasets[stat.Cube]; !found {
			datasets[stat.Cube] = []chart.EPoint{}
		}
		if stat.PTwo {
			stat.Value += 2000
		}
		datasets[stat.Cube] = append(datasets[stat.Cube], chart.EPoint{
			X: float64(stat.Timestamp),
			Y: float64(stat.Value),
		})
	}

	plot := chart.ScatterChart{}
	i := 0
	for cube, stat := range datasets {
		plot.AddData(cube, stat, chart.PlotStyleLinesPoints, chart.AutoStyle(i, false))
		i++
	}
	plot.XRange.TicSetting.Format = func(timestamp float64) string {
		return time.Unix(int64(math.Floor(timestamp)), 0).Format("06-01-02 15h04")
	}
	plot.YRange.TicSetting.Format = func(timestamp float64) string {
		return GenerateTimeDisplay(uint64(math.Floor(timestamp)))
	}
	return plot
}

func WriteImage(plot chart.ScatterChart) telegram.FileReader {
	img := image.NewRGBA(image.Rect(0, 0, 1280, 720))
	igr := imgg.AddTo(img, 0, 0, 1280, 720, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}, nil, nil)
	plot.Plot(igr)

	writer := bytes.NewBuffer(nil)
	err := png.Encode(writer, img)
	if err != nil {
		log.Fatalf("error writing a png image: %v", err)
	}

	return telegram.FileReader{
		Name:   "stats.png",
		Reader: writer,
	}
}
