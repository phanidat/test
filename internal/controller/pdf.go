package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"

	"git.kasikornline.com/pdf-decrypt/internal/models"
)

func GetBook(c *gin.Context) {

	response := models.Book{
		Name:   "Randy",
		Title:  "Book 1",
		Author: "Randy",
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)

	json.NewEncoder(c.Writer).Encode(response)
}

func DecryptPdfFile(c *gin.Context) {

	var req models.PDFDecryptRequest

	if err := c.ShouldBind(&req); err != nil {
		boom.BadRequest(c.Writer, err.Error())
		return
	}

	//read file
	file, err := req.PdfFile.Open()
	if err != nil {
		boom.BadRequest(c.Writer, err.Error())
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	var decryptedBuf bytes.Buffer
	w := &decryptedBuf

	_, err = io.Copy(&buf, file)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error())
		return
	}

	rs := bytes.NewReader(buf.Bytes())

	conf := model.NewAESConfiguration("", req.PdfPassword, 256)
	err = api.Decrypt(rs, w, conf)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error())
		return
	}

	var waterBuf bytes.Buffer
	x := &waterBuf
	rs = bytes.NewReader(decryptedBuf.Bytes())
	mm := model.DefaultWatermarkConfig()

	mm.TextString = "Hahaha"
	mm.Scale = 1
	mm.Opacity = 0.3
	mm.OnTop = true

	err = api.AddWatermarks(rs, x, []string{}, mm, nil)
	if err != nil {

	}
	c.Writer.Header().Set("Content-Type", "application/pdf")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(waterBuf.Bytes())
}
