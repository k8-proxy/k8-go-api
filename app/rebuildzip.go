package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type req struct {
	Base64                 string `json:"Base64"`
	ContentManagementFlags struct {
		PdfContentManagement struct {
			Metadata int `json:"Metadata"`
		} `json:"PdfContentManagement"`
	} `json:"ContentManagementFlags"`
}

func rebuildzip(w http.ResponseWriter, r *http.Request) {

	//log about request
	log.Println("method:", r.Method)
	log.Printf("%v\n", r.URL)
	log.Printf("%v\n", r.RemoteAddr)
	log.Printf("%v\n", r.Host)
	log.Printf("%v\n", r.Header)
	log.Printf("%v\n", r.Header.Get("Content-Type"))

	//m max 5 MB file name we can change ut
	r.ParseMultipartForm(5 << 20)

	//myfileparam is the name of file in post request body

	log.Println(r.PostFormValue("contentManagementFlagJson"))

	cont := r.PostFormValue("contentManagementFlagJson")

	var mp map[string]json.RawMessage

	err := json.Unmarshal([]byte(cont), &mp)
	if err != nil {
		log.Println("error json:", err)
		http.Error(w, "malformed json", http.StatusBadRequest)

		return
	}

	log.Println("json request", mp)

	file, handler, err := r.FormFile("file")

	if err != nil {
		log.Println("file not found", err)
		http.Error(w, "file not found", http.StatusBadRequest)

		return
	}

	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		http.Error(w, "server intenal error", http.StatusInternalServerError)

		return
	}

	if handler.Header.Get("Content-Type") != "application/zip" || http.DetectContentType(buf) != "application/zip" {
		log.Println(w, "mediatype is", handler.Header.Get("Content-Type"))

		http.Error(w, "upload file should be zip file format", http.StatusUnsupportedMediaType)

		return

	}

	//this only to parse post form to extract data for log
	if errp := r.ParseForm(); errp != nil {
		log.Println(err)
	}

	for k, v := range r.Form {
		log.Printf("Form[%q] = %q\n", k, v)
	}

	log.Printf("%v\n", handler.Filename)
	log.Printf("%v\n", handler.Size)

	log.Printf("%v\n", handler.Header.Get("Content-Type"))
	log.Printf("%v\n", http.DetectContentType(buf))

	addgwheader(w, temp)
	s, e := w.Write(buf)
	if e != nil {
		log.Println(e)
		return
	}
	log.Println(s)

	// so  here we can use either open file or  ioutil.write file
	/*
		fmt.Fprintf(w, "%v\n", handler.Header)
		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	*/

}

func ioutilCopytodisk(f multipart.File, wn string, p os.FileMode) error {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return err
	}
	err = ioutil.WriteFile(wn, data, p)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}