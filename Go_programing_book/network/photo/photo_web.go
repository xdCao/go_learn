package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "./uploads"
	ListDir    = 0x0001
)

var templates map[string]*template.Template = make(map[string]*template.Template)

func init() {
	files, _ := os.ReadDir("./views")
	for _, tmpl := range files {
		t := template.Must(template.ParseFiles("./views/" + tmpl.Name()))
		templates[tmpl.Name()] = t
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderHtml(w, "upload.html", nil)
		check(err)
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			check(err)
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}

}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	path := UPLOAD_DIR + "/" + imageId
	if !isExist(path) {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("content-type", "image")
	http.ServeFile(w, r, path)
}

func renderHtml(w http.ResponseWriter, temp string, locals map[string]interface{}) error {
	templates[temp].Execute(w, locals)
	return nil
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := os.ReadDir(UPLOAD_DIR)

	locals := make(map[string]interface{})
	images := []string{}

	for _, v := range files {
		imgid := v.Name()
		images = append(images, imgid)
	}
	locals["images"] = images
	err := renderHtml(w, "list.html", locals)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError) // 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
			}
		}()
		fn(w, r)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExist(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/list", safeHandler(listHandler))
	http.ListenAndServe(":9999", nil)
}
