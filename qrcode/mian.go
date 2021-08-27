package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
	"image/png"
	"log"
	"os"
)

func main(){
	path,_ := os.Getwd()
	filename := path + "/qrcode/" + "qrcode.png"
	log.Println(filename)
	generateQrcode(filename,"Hello,this is my ID.")
	decodeQrcode(filename)
}

func generateQrcode(filename,message string){

	qrCode, _ := qr.Encode(message, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create(filename)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}

func decodeQrcode(filename string){
	fi, err := os.Open(filename)
	if err != nil{
		log.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil{
		log.Println(err.Error())
		return
	}
	log.Println(qrmatrix.Content)
}

