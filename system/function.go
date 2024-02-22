package system

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
  "mime/multipart"
  "io"
  "path/filepath"

  "github.com/PuerkitoBio/goquery"
	"github.com/nfnt/resize"
  
)

func Instagram(URL string) ([]map[string]string, error) {
  client := &http.Client{}
  resp, err := client.Get("https://indown.io/")
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  bodyBytes, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  body := string(bodyBytes)
  doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
  if err != nil {
    return nil, err
  }

  referer := doc.Find("input[name=referer]").AttrOr("value", "")
  locale := doc.Find("input[name=locale]").AttrOr("value", "")
  token := doc.Find("input[name=_token]").AttrOr("value", "")


  params := url.Values{}
  params.Set("link", URL)
  params.Set("referer", referer)
  params.Set("locale", locale)
  params.Set("_token", token)

  req, err := http.NewRequest("POST", "https://indown.io/download", strings.NewReader(params.Encode()))
  if err != nil {
    return nil, err
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Content-Length", fmt.Sprint(len(params.Encode())))
  req.Header.Add("Cookie", strings.Join(resp.Header["Set-Cookie"], "; "))

  resp, err = client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  bodyBytes, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  body = string(bodyBytes)
  doc, err = goquery.NewDocumentFromReader(strings.NewReader(body))
  if err != nil {
    return nil, err
  }

  result := []map[string]string{}
  doc.Find("#result video").Each(func(i int, s *goquery.Selection) {
    thumbnail, _ := s.Attr("poster")
    videoURL, _ := s.Find("source").Attr("src")

    result = append(result, map[string]string{
      "type":      "video",
      "thumbnail": thumbnail,
      "url":       videoURL,
    })
  })

  doc.Find("#result img").Each(func(i int, s *goquery.Selection) {
    imageURL, _ := s.Attr("src")

    result = append(result, map[string]string{
      "type": "image",
      "url":  imageURL,
    })
  })

  return result, nil
}

func Remini(buffer []byte, method string) ([]byte, error) {
  methods := []string{"enhance", "recolor", "dehaze"}
  if !contains(methods, method) {
    method = methods[0]
  }

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  err := writer.WriteField("model_version", "1")
  if err != nil {
    return nil, err
  }

  part, err := writer.CreateFormFile("image", "enhance_image_body.jpg")
  if err != nil {
    return nil, err
  }
  _, err = part.Write(buffer)
  if err != nil {
    return nil, err
  }
  err = writer.Close()
  if err != nil {
    return nil, err
  }

  req, err := http.NewRequest("POST", "https://inferenceengine.vyro.ai/"+method, body)
  if err != nil {
    return nil, err
  }
  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Windows; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.114 Safari/537.36")
  req.Header.Set("Connection", "Keep-Alive")
  req.Header.Set("Accept-Encoding", "gzip")
  req.Header.Set("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != http.StatusOK {
    return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
  }

  return ioutil.ReadAll(resp.Body)
}

func contains(slice []string, item string) bool {
  for _, s := range slice {
    if s == item {
      return true
    }
  }
  return false
}

func ShortUrl(query string) (string, error) {
  client := &http.Client{}
  data := url.Values{}
  data.Set("url", query)

  req, err := http.NewRequest("POST", "https://tinyurl.com/api-create.php", strings.NewReader(data.Encode())) 
  if err != nil {
    return "", err
  }

  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  return string(body), nil
}

var c = time.Now()
func GetUptime() time.Time{
  return c
}

func Capcutdl(url string) ([]byte, error) {
  resp, err := http.Get(url)
if err != nil {
  return nil, err
}
  defer resp.Body.Close()


  re := regexp.MustCompile(`\d+`)
  token := re.FindString(strings.Split(resp.Request.URL.String(), "?")[0])

  if token == "" {
    return nil, nil
  }

  downloadURL := fmt.Sprintf("https://ssscap.net/api/download/%s", token)
  req, err := http.NewRequest(http.MethodGet, downloadURL, nil)
if err != nil {
  return nil, err
}

  req.Header.Set("Cookie", "sign=2cbe441f7f5f4bdb8e99907172f65a42; device-time=1685437999515")

  client := &http.Client{}
  downResp, err := client.Do(req)
if err != nil {
  return nil, err
}
  defer downResp.Body.Close()

  data, err := ioutil.ReadAll(downResp.Body)
if err != nil {
  return nil, err
}

 return data, nil
}

type Result struct {
  Desc   string `csv:"desc"`
  Thumb  string `csv:"thumb"`
  HD     string `csv:"video_hd"`
  SD     string `csv:"video_sd"`
  URL    string `csv:"url"`
  Locale string `csv:"locale"`
}

func Fb(URL string) (*Result, error) {
  data := URL + "?locale=en"
  resp, err := http.PostForm("https://getmyfb.com/process", url.Values{"id": {data}})
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  doc, err := goquery.NewDocumentFromReader(resp.Body)
  if err != nil {
    return nil, err
  }

  thumb := doc.Find(".results-item img").AttrOr("src", "")
  desc := doc.Find(".results-item > .results-item-text").Text()
  hd := doc.Find(".results-download > .results-list > .results-list-item a").Eq(0).AttrOr("href", "")
  sd := doc.Find(".results-download > .results-list > .results-list-item a").Eq(1).AttrOr("href", "")

  result := &Result{
    Desc:   desc,
    Thumb:  thumb,
    HD:     hd,
    SD:     sd,
    URL:    URL,
    Locale: "en",
  }

  return result, nil
}

func Sleep(duration time.Duration) {
  time.Sleep(duration)
}

type ConversionOptions struct {
	Width  int
	Height int
	Fps    int
}

/* delay */
func Delay(duration time.Duration) {
	time.Sleep(duration)
}

/* json marshal */
func MyLog(msg interface{}) string {
	jsonRes, err := json.MarshalIndent(msg, "", " ")
	if err != nil {
		// Handle error jika terjadi kesalahan saat marshalling
		return fmt.Sprintf("Error: %v", err)
	}
	return string(jsonRes)
}

/* format number */
func FormatNumber(jid string) string {
	jid = strings.ReplaceAll(jid, "+", "")
	jid = strings.ReplaceAll(jid, "-", "")
	jid = strings.ReplaceAll(jid, " ", "")
	return jid
}

/* add https */
func Url(query string) string {
	if !strings.HasPrefix(query, "http://") && !strings.HasPrefix(query, "https://") {
		query = "https://" + query
	}
	return query
}

/* get result api */
func GetResult(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}


func ToAudio(buffer []byte, ext string) ([]byte, error) {
  tmpFile := fmt.Sprintf("%d.%s", time.Now().UnixNano(), ext)
  outFile := fmt.Sprintf("%d.%s.mp3", time.Now().UnixNano(), ext)

  err := ioutil.WriteFile(tmpFile, buffer, 0644)
  if err != nil {
    return nil, err
  }
  defer os.Remove(tmpFile)

  cmd := exec.Command("ffmpeg",
    "-y",
    "-i", tmpFile,
    "-vn",
    "-ac", "2",
    "-b:a", "128k",
    "-ar", "44100",
    "-f", "mp3",
    outFile,
  )
  var stderr bytes.Buffer
  cmd.Stderr = &stderr
  err = cmd.Run()
  if err != nil {
    return nil, fmt.Errorf("error running ffmpeg: %s\n%s", err, stderr.String())
  }
  defer os.Remove(outFile)

  audioData, err := ioutil.ReadFile(outFile)
  if err != nil {
    return nil, err
  }

  return audioData, nil
}


/* image to webp */
func ImgToWebp(RawPath string, ConvertedPath string) error {
	exc := exec.Command("ffmpeg",
		"-i", RawPath,
		"-vf", "scale='min(320,iw)':min'(320,ih)':force_original_aspect_ratio=decrease,fps=15, pad=320:320:-1:-1:color=white@0.0, split [a][b]; [a] palettegen=reserve_transparent=on:transparency_color=ffffff [p]; [b][p] paletteuse",
		"-framerate", "15",
		ConvertedPath,
	)

	err := exc.Run()
	if err != nil {
		return err
	}

	createExif := fmt.Sprintf("webpmux -set exif %s %s -o %s", "tmp/exif/mywabot.exif", ConvertedPath, ConvertedPath)
	cmd := exec.Command("bash", "-c", createExif)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to set webp metadata", err)
	}
	return nil
}

/* video to webp */
func VideoToWebp(RawPath string, ConvertedPath string) error {
	exc := exec.Command("ffmpeg",
		"-i", RawPath,
		"-vf", "scale='min(320,iw)':min'(320,ih)':force_original_aspect_ratio=decrease,fps=15, pad=320:320:-1:-1:color=white@0.0, split [a][b]; [a] palettegen=reserve_transparent=on:transparency_color=ffffff [p]; [b][p] paletteuse",
		"-loop", "0",
		"-ss", "00:00:00",
		"-t", "00:00:05",
		"-preset", "default",
		"-an", "-vsync",
		"0",
		ConvertedPath,
	)

	err := exc.Run()
	if err != nil {
		return err
	}

	createExif := fmt.Sprintf("webpmux -set exif %s %s -o %s", "tmp/exif/mywabot.exif", ConvertedPath, ConvertedPath)
	cmd := exec.Command("bash", "-c", createExif)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to set webp metadata", err)
	}

	return nil
}

/* format size file */
var sizes = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}

func FormatSize(s float64, base float64) string {
	unitsLimit := len(sizes)
	i := 0
	for s >= base && i < unitsLimit {
		s = s / base
		i++
	}

	f := "%.0f %s"
	if i > 1 {
		f = "%.2f %s"
	}

	return fmt.Sprintf(f, s, sizes[i])
}

/* custom id message */
func GenerateID() string {
	id := make([]byte, 14)
	_,
		err := rand.Read(id)
	if err != nil {
		panic(err)
	}
	return "WSC" + strings.ToUpper(hex.EncodeToString(id))
}

/* create exif */
func CreateExif(fileName string, packname string, author string) *string {

	jsonData := map[string]interface{}{
		"sticker-pack-id":        "amirull.dev",
		"sticker-pack-name":      packname,
		"sticker-pack-publisher": author,
		"android-app-store-link": "https://play.google.com/store/apps/details?id=",
		"ios-app-store-link":     "https://apps.apple.com/app/id625334537",
		"emojis": []string{
			"👋"},
	}

	jsonBytes,
		err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	littleEndian := []byte{
		0x49,
		0x49,
		0x2a,
		0x00,
		0x08,
		0x00,
		0x00,
		0x00,
		0x01,
		0x00,
		0x41,
		0x57,
		0x07,
		0x00,
	}

	bytes := []byte{
		0x00,
		0x00,
		0x16,
		0x00,
		0x00,
		0x00}

	len := len(jsonBytes)
	var last string

	if len > 256 {
		len = len - 256
		bytes = append([]byte{
			0x01}, bytes...)
	} else {
		bytes = append([]byte{
			0x00}, bytes...)
	}

	if len < 16 {
		last = fmt.Sprintf("0%x", len)
	} else {
		last = fmt.Sprintf("%x", len)
	}

	buf2,
		err := hex.DecodeString(last)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	buf3 := bytes
	buf4 := jsonBytes

	buffer := append(littleEndian, buf2...)
	buffer = append(buffer, buf3...)
	buffer = append(buffer, buf4...)

	err = os.WriteFile("tmp/exif/"+fileName, buffer, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return &fileName
}

/* get file */
func GetFile(buffer []byte) (string, []byte, error) {
	contentType := http.DetectContentType(buffer)
	exts, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", nil, err
	}

	if len(exts) == 0 {
		return "", nil, errors.New("unknown file type")
	}

	ext := exts[0]
	return ext, buffer, nil
}

/* get random name */
func GetRandomName(exts []string, ext string) string {
	return "random_filename" + ext
}

/* rgb to uint */
func RgbaToUint32(r, g, b, a int) uint32 {
	return uint32((a << 24) | (r << 16) | (g << 8) | b)
}

/* hex to uint */
func HextoUint32(hexString string) (uint32, error) {
	// Dekode string heksadesimal menjadi []byte
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return 0, err
	}

	// Konversi []byte ke uint32
	var result uint32
	if len(bytes) >= 4 {
		result = uint32(bytes[0])<<24 | uint32(bytes[1])<<16 | uint32(bytes[2])<<8 | uint32(bytes[3])
	} else {
		return 0, fmt.Errorf("string heksadesimal tidak memiliki cukup byte untuk dikonversi menjadi uint32")
	}

	return result, nil
}

/* resize image */
func Resize(imgByte []byte, width int, height int, path string) error {
	var width_ori = width
	var height_ori = height

	var percentage = 10 // di isi biar gk kosong
	var max_pixel = 150

	if width_ori > height_ori {
		percentage = max_pixel / (width_ori / 100)
	} else {
		percentage = max_pixel / (height_ori / 100)
	}

	width = (width_ori * percentage) / 100
	height = (height_ori * percentage) / 100

	img, err := jpeg.Decode(bytes.NewReader(imgByte))
	if err != nil {
		return err
	}

	newImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	err = jpeg.Encode(out, newImg, nil)
	if err != nil {
		return err
	}

	return nil
}

/* get thumbnail vid */
func GetThumbnail(videoBytes []byte, maxPixel int, videoFile string, thumb string) (string, error) {
	err := saveBytesToFile(videoBytes, videoFile)
	if err != nil {
		return "", err
	}
	//defer os.Remove(videoFile)

	thumbnailPath := "tmp/" + thumb
	cmd := exec.Command("ffmpeg", "-i", videoFile, "-ss", "00:00:01", "-vframes", "1", "-f", "image2", thumbnailPath)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing ffmpeg: %v", err)
	}

	img, err := loadImageFromFile(thumbnailPath)
	if err != nil {
		return "", fmt.Errorf("error loading thumbnail image: %v", err)
	}

	widthOri := img.Bounds().Dx()
	heightOri := img.Bounds().Dy()

	var percentage int
	if widthOri > heightOri {
		percentage = maxPixel * 100 / widthOri
	} else {
		percentage = maxPixel * 100 / heightOri
	}

	width := widthOri * percentage / 100
	height := heightOri * percentage / 100

	// Resize thumbnail
	resizedThumbnail := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	// Simpan thumbnail yang sudah diresize
	resizedThumbnailPath := "tmp/0" + thumb

	err = saveImageToFile(resizedThumbnailPath, resizedThumbnail)
	if err != nil {
		return "", fmt.Errorf("error saving resized thumbnail: %v", err)
	}
	os.Remove(thumbnailPath)
	return resizedThumbnailPath, nil
}

func saveBytesToFile(data []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func loadImageFromFile(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

func saveImageToFile(filename string, img image.Image) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	return err
}

/* url to byte */
func Getbyte(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

/* res to byte */
func ToByte(resultData interface{}) ([]byte, error) {
	switch data := resultData.(type) {
	case string:
		if u, err := url.Parse(data); err == nil && u.Scheme != "" && u.Host != "" {
			return Getbyte(data)
		}
	case []byte:
		return data, nil
	default:
		return nil, fmt.Errorf("tipe data tidak didukung")
	}

	return nil, fmt.Errorf("tidak dapat memproses result data")
}

/* tiktok url */
func IsTiktok(url string) bool {
	tikTokURLPattern := `^(https?://)?(www\.)?(tiktok\.com|vt\.tiktok\.com|vm\.tiktok\.com|tiktok\.com/@[a-zA-Z0-9_-]+/video/[0-9]+)(\S*)?$`
	re := regexp.MustCompile(tikTokURLPattern)

	return re.MatchString(url)
}

func Upload(mediaPath string) (string, error) {
  if _, err := os.Stat(mediaPath); os.IsNotExist(err) {
      return "", fmt.Errorf("File not found")
  }

  media, err := os.Open(mediaPath)
  if err != nil {
      return "", err
  }
  defer media.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  part, err := writer.CreateFormFile("files[]", filepath.Base(mediaPath))
  if err != nil {
      return "", err
  }

  _, err = io.Copy(part, media)
  if err != nil {
      return "", err
  }

  err = writer.Close()
  if err != nil {
      return "", err
  }

  req, err := http.NewRequest("POST", "https://pomf.lain.la/upload.php", body)
  if err != nil {
      return "", err
  }
  req.Header.Set("Content-Type", writer.FormDataContentType())
  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      return "", err
  }
  defer resp.Body.Close()

  responseData, err := ioutil.ReadAll(resp.Body)
  if err != nil {
      return "", err
  }

  // Parse the response body
  type ResponseData struct {
      Files []struct {
          URL string `json:"url"`
      } `json:"files"`
  }
  var data ResponseData
  err = json.Unmarshal(responseData, &data)
  if err != nil {
      return "", err
  }

  if len(data.Files) == 0 {
      return "", fmt.Errorf("Failed to retrieve file URL")
  }

  return data.Files[0].URL, nil
}