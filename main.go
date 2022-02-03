package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GrandOichii/argsparser"
	"github.com/GrandOichii/mtgsdk"
)

const (
	keyCARDNAMEQUERY = "card name query"
	keySETNAME       = "set name"
	keyOUTPATH       = "out path"
	keyDECKPATH      = "deck path"
)

var correlationMap = map[string]string{
	"-set":       keySETNAME,
	"--set":      keySETNAME,
	"-setname":   keySETNAME,
	"--setname":  keySETNAME,
	"-out":       keyOUTPATH,
	"--out":      keyOUTPATH,
	"-outpath":   keyOUTPATH,
	"--outpath":  keyOUTPATH,
	"-deck":      keyDECKPATH,
	"--deck":     keyDECKPATH,
	"-deckpath":  keyDECKPATH,
	"--deckpath": keyDECKPATH,
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Specify the cards name")
		return
	}
	parsed, err := argsparser.Parse(os.Args, correlationMap)
	checkError(err)
	parsed[keyCARDNAMEQUERY] = os.Args[1]
	err = mtgsdk.DownloadCardImages(
		map[string]string{
			mtgsdk.CardNameKey: parsed[keyCARDNAMEQUERY],
			mtgsdk.SetNameKey:  parsed[keySETNAME],
		},
		parsed[keyDECKPATH],
		parsed[keyOUTPATH],
		mtgsdk.ImageQualityNormal,
	)
	checkError(err)
}
