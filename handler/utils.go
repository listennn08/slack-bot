package handler

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

func convertBytesToArray(b []byte) [][]string {
  reader := csv.NewReader(bytes.NewBuffer(b))
  reader.FieldsPerRecord = -1
  reader.TrimLeadingSpace = true

  var ret [][]string
  for {
    record, err := reader.Read()
    if err == io.EOF {
      break
    }

    if err != nil {
      log.Println(fmt.Errorf("Can not read csv data from reader: %w", err))
      return nil
    }
    var raw []string

    for _, cell := range record {
      if cell != "" {
        raw = append(raw, cell)
      }
    }

    ret = append(ret, raw)
  }

  return ret
}

func getSheet() []byte {
  sheetURL := os.Getenv("SHEET_URL")
  res, err := http.Get(sheetURL)
  if err != nil {
    log.Println(fmt.Errorf("Can not get csv data from sheet: %w", err))
    return nil
  }

  defer res.Body.Close()

  text, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.Println(fmt.Errorf("Can not read csv data: %w", err))
    return nil
  }

  return text
}

func getRandomEat(user *slack.User) string {
  var result = convertBytesToArray(getSheet())
  y := len(result)
  randomY := rand.Intn(y)

  x := len(result[randomY])
  randomX := rand.Intn(x)

  return "Hi " + user.Profile.RealName + "\n今天午餐吃\n*" + result[randomY][randomX] + "*"
}
