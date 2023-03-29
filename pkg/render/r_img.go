package render

// Generate image from library
// https://github.com/danielgatis/imgcat
import (
	"log"

	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/creack/pty"
	"github.com/disintegration/imaging"
	"github.com/gabriel-vasile/mimetype"
	"github.com/mat/besticon/ico"
	"github.com/mattn/go-isatty"
	"github.com/zbeaver/cafe/pkg/vui"
)

var resizeFactorX = 1
var resizeFactorY = 1
var resizeOffsetY = 1

const DEFAULT_TERM_COLS = 80
const DEFAULT_TERM_ROWS = 24
const FPS = 15

const ANSI_CURSOR_UP = "\x1B[%dA"
const ANSI_CURSOR_HIDE = "\x1B[?25l"
const ANSI_CURSOR_SHOW = "\x1B[?25h"
const ANSI_BG_TRANSPARENT_COLOR = "\x1b[0;39;49m"
const ANSI_BG_RGB_COLOR = "\x1b[48;2;%d;%d;%dm"
const ANSI_FG_TRANSPARENT_COLOR = "\x1b[0m "
const ANSI_FG_RGB_COLOR = "\x1b[38;2;%d;%d;%dm▄"
const ANSI_RESET = "\x1b[0m"

type Img struct{}

func (r *Img) Style(base styling, n vui.INode) styling {
	elm, ok := n.(vui.Elementary)
	if !ok {
		return base
	}
	return TransformFrom(base)(elm.Style())
}

func (r *Img) Render(n vui.INode, s styling, child string) string {
	if elm, ok := n.(vui.ImgElementary); ok {
		return print(escape(scale(decode(read(elm.Src())), elm.Width(), elm.Height())))
	}
	return "Counld not load image"
}

func read(input string) []byte {
	var err error
	var buf []byte

	u, err := url.Parse(input)
	if err != nil {
		return []byte("cannot load image")
	}

	if u.Scheme != "" && u.Host != "" {
		res, err := http.Get(input)

		if err != nil {
			return []byte("cannot load image")
		}

		if buf, err = io.ReadAll(res.Body); err != nil {
			return []byte("cannot load image")
		}
	} else {
		if buf, err = os.ReadFile(input); err != nil {
			return []byte("cannot load image")
		}
	}

	return buf
}

func decode(buf []byte) []image.Image {
	mime, err := mimetype.DetectReader(bytes.NewReader(buf))
	if err != nil {
		log.Panicf("failed to detect the mime type: %v", err)
	}

	allowed := []string{"image/gif", "image/png", "image/jpeg", "image/bmp", "image/x-icon"}
	if !mimetype.EqualsAny(mime.String(), allowed...) {
		log.Fatal("invalid MIME type")
	}

	frames := make([]image.Image, 0)

	if mime.Is("image/gif") {
		gifImage, err := gif.DecodeAll(bytes.NewReader(buf))

		if err != nil {
			log.Panicf("failed to decode the gif: %v", err)
		}

		var lowestX int
		var lowestY int
		var highestX int
		var highestY int

		for _, img := range gifImage.Image {
			if img.Rect.Min.X < lowestX {
				lowestX = img.Rect.Min.X
			}
			if img.Rect.Min.Y < lowestY {
				lowestY = img.Rect.Min.Y
			}
			if img.Rect.Max.X > highestX {
				highestX = img.Rect.Max.X
			}
			if img.Rect.Max.Y > highestY {
				highestY = img.Rect.Max.Y
			}
		}

		imgWidth := highestX - lowestX
		imgHeight := highestY - lowestY

		overPaintImage := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
		draw.Draw(overPaintImage, overPaintImage.Bounds(), gifImage.Image[0], image.Point{}, draw.Src)

		for _, srcImg := range gifImage.Image {
			draw.Draw(overPaintImage, overPaintImage.Bounds(), srcImg, image.Point{}, draw.Over)
			frame := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
			draw.Draw(frame, frame.Bounds(), overPaintImage, image.Point{}, draw.Over)
			frames = append(frames, frame)
		}

		return frames
	} else {
		var frame image.Image
		var err error

		if mime.Is("image/x-icon") {
			frame, err = ico.Decode(bytes.NewReader(buf))
		} else {
			frame, _, err = image.Decode(bytes.NewReader(buf))
		}

		if err != nil {
			log.Panicf("failed to decode the image: %v", err)
		}

		imb := frame.Bounds()
		if imb.Max.X < 2 || imb.Max.Y < 2 {
			log.Fatal("the input image is to small")
		}

		frames = append(frames, frame)
	}

	return frames
}

func scale(frames []image.Image, w, h int) []image.Image {
	type data struct {
		i  int
		im image.Image
	}

	var err error

	cols := DEFAULT_TERM_COLS
	rows := DEFAULT_TERM_ROWS

	if isatty.IsTerminal(os.Stdout.Fd()) {
		if rows, cols, err = pty.Getsize(os.Stdout); err != nil {
			log.Panicf("failed to get the terminal size: %v", err)
		}
	}

	if w == 0 {
		w = cols
		h = rows
	}

	l := len(frames)
	r := make([]image.Image, l)
	c := make(chan *data, l)

	for i, f := range frames {
		go func(i int, f image.Image) {
			c <- &data{i, imaging.Fit(f, w, h, imaging.Lanczos)}
		}(i, f)
	}

	for range r {
		d := <-c
		r[d.i] = d.im
	}

	return r
}

func escape(frames []image.Image) [][]string {
	type data struct {
		i   int
		str string
	}

	escaped := make([][]string, 0)

	for _, f := range frames {
		imb := f.Bounds()
		maxY := imb.Max.Y - imb.Max.Y%2
		maxX := imb.Max.X

		c := make(chan *data, maxY/2)
		lines := make([]string, maxY/2)

		for y := 0; y < maxY; y += 2 {
			go func(y int) {
				var sb strings.Builder

				for x := 0; x < maxX; x++ {
					r, g, b, a := f.At(x, y).RGBA()
					if a>>8 < 128 {
						sb.WriteString(ANSI_BG_TRANSPARENT_COLOR)
					} else {
						sb.WriteString(fmt.Sprintf(ANSI_BG_RGB_COLOR, r>>8, g>>8, b>>8))
					}

					r, g, b, a = f.At(x, y+1).RGBA()
					if a>>8 < 128 {
						sb.WriteString(ANSI_FG_TRANSPARENT_COLOR)
					} else {
						sb.WriteString(fmt.Sprintf(ANSI_FG_RGB_COLOR, r>>8, g>>8, b>>8))
					}
				}

				sb.WriteString(ANSI_RESET)
				sb.WriteString("\n")

				c <- &data{y / 2, sb.String()}
			}(y)
		}

		for range lines {
			line := <-c
			lines[line.i] = line.str
		}

		escaped = append(escaped, lines)
	}

	return escaped
}

func print(frames [][]string) string {
	str := strings.Builder{}

	frameCount := len(frames)

	if frameCount == 1 {
		str.WriteString(strings.Join(frames[0], ""))
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		tick := time.Tick(time.Second / time.Duration(FPS))
		h := len(frames[0]) + 2 // two extra lines for the exit msg
		playing := true

		go func() {
			<-c
			playing = false
		}()

		for i := 0; playing; i++ {
			if i != 0 {
				str.WriteString(fmt.Sprintf(ANSI_CURSOR_UP, h))
			}

			str.WriteString(strings.Join(frames[i%frameCount], ""))

			<-tick
		}
	}
	return str.String()
}
