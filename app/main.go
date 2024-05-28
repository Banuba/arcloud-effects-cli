package main

import (
	"archive/zip"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/umputun/go-flags"
)

var opts struct {
	Source      string `short:"s" long:"source" description:"Path to effects' zip file" required:"true"`
	Destination string `short:"d" long:"destination" description:"Destination folder" default:"effects"`
	ApiUrl      string `short:"u" long:"api-url" description:"Arcloud URL" required:"true"`
}

type EffectArchive struct {
	File      *os.File
	ZipWriter *zip.Writer
}

type JsonEffect struct {
	URL     string `json:"URL"`
	Preview string `json:"Preview"`
	ETag    string `json:"ETag"`
}

func NewJsonEffect(name, etag, apiUrl string) JsonEffect {
	u, err := url.JoinPath(apiUrl, name)
	if err != nil {
		log.Fatalf("Error creating URL with params name: %s, etag: %s, apiUrl: %s", name, etag, apiUrl)
	}
	return JsonEffect{
		URL:     fmt.Sprintf("%s.zip", u),
		Preview: fmt.Sprintf("%s.png", u),
		ETag:    fmt.Sprintf(`"%s"`, etag),
	}
}

func splitPath(path string) (root, filePath string) {
	split := strings.SplitN(path, "/", 2)
	return split[0], split[1]
}

func isValidFile(file *zip.File) bool {
	return !(file.FileInfo().IsDir() ||
		strings.HasPrefix(file.Name, "__MACOSX") ||
		strings.HasSuffix(file.Name, ".DS_Store"))
}

func createEffects(zipReader *zip.ReadCloser, dest string) map[string]EffectArchive {
	effectsMap := make(map[string]EffectArchive)
	defer zipReader.Close()

	for _, file := range zipReader.File {

		if !isValidFile(file) {
			continue
		}

		root, filePath := splitPath(file.Name)
		root = strings.ReplaceAll(root, " ", "_")
		if root == "" {
			log.Fatalf("Wrong archive structure. File %s in root diractory.", file.Name)
		}
		if _, ok := effectsMap[root]; !ok {
			root = strings.ReplaceAll(root, " ", "_")
			archive, err := os.Create(path.Join(dest, root+".zip"))
			if err != nil {
				log.Fatalf("Error creating archive %s.zip: %v", root, err)
			}
			defer archive.Close()

			zipWriter := zip.NewWriter(archive)
			defer zipWriter.Close()

			effectsMap[root] = EffectArchive{
				File:      archive,
				ZipWriter: zipWriter,
			}
		}

		func() {
			fReader, err := file.Open()
			if err != nil {
				log.Fatalf("Error opening file %s: %v", filePath, err)
			}
			defer fReader.Close()

			fileData, err := io.ReadAll(fReader)
			if err != nil {
				log.Fatalf("Error reading file %s: %v", filePath, err)
			}

			fWriter, err := effectsMap[root].ZipWriter.Create(path.Join(root, filePath))
			if err != nil {
				log.Fatalf("Error creating file %s in archive: %v", filePath, err)
			}

			_, err = fWriter.Write(fileData)
			if err != nil {
				log.Fatalf("Error writing file %s: %v", filePath, err)
			}

			if filePath == "preview.png" {
				previewFilePath := fmt.Sprintf(path.Join(dest, "%s.png"), root)
				err = os.WriteFile(previewFilePath, fileData, fs.ModePerm)
				if err != nil {
					log.Fatalf("Error creating preview image %s: %v", previewFilePath, err)
				}
			}
		}()
	}
	return effectsMap
}

func run(zipFilePath, apiUrl, dest string) {
	zipReader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	effectsMap := createEffects(zipReader, dest)

	var keys []string
	for k := range effectsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var j []JsonEffect
	for _, k := range keys {
		data, _ := os.ReadFile(effectsMap[k].File.Name())
		etag := fmt.Sprintf("%x", md5.Sum(data))
		j = append(j, NewJsonEffect(k, etag, apiUrl))
	}

	file, err := os.Create(path.Join(dest, "api_response"))
	if err != nil {
		log.Fatalf("Error creating manifest file: %s", err)
	}
	defer file.Close()

	e := struct {
		Effects []JsonEffect `json:"effects"`
	}{j}

	data, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		log.Fatalf("Error marshaling: %s", err)
	}

	_, err = file.Write(data)
	if err != nil {
		log.Fatalf("Error writing manifest file: %s", err)
	}
}

func main() {
	p := flags.NewParser(&opts, flags.PrintErrors|flags.PassDoubleDash|flags.HelpFlag)
	p.SubcommandsOptional = true
	if _, err := p.Parse(); err != nil {
		os.Exit(2)
	}

	if _, err := os.Stat(opts.Destination); os.IsNotExist(err) {
		err := os.Mkdir(opts.Destination, fs.ModePerm)
		if err != nil {
			log.Fatalf("Error creating folder %s", opts.Destination)
		}
	}

	run(opts.Source, opts.ApiUrl, opts.Destination)
}
