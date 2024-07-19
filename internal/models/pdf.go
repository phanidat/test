package models

import "mime/multipart"

type PDFDecryptRequest struct {
	PdfPassword string                `form:"password"`
	PdfFile     *multipart.FileHeader `form:"pdf-file"`
}

type Book struct {
	Name   string
	Title  string
	Author string
}
