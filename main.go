package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// Global var
var cacheDirPath string

func getSvgDir() string {
	return cacheDirPath + "/svg-files/"
}
func getLogsDir() string {
	dirPath := cacheDirPath + "/logs-files/"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModePerm)
	}
	return dirPath
}

// runCommand and wait until finish
// returns logsFilePath
func runCommand(command string, args []string, logFileName string) string {
	var err error
	workDir := "./"
	logsFilePath := getLogsDir() + logFileName + ".log"

	// Set log files
	stdoutFile, err := os.Create(logsFilePath)
	if err != nil {
		panic(err)
	}
	defer stdoutFile.Close()
	stderrFile, err := os.Create(logsFilePath)
	if err != nil {
		panic(err)
	}
	defer stderrFile.Close()

	// Convert env data
	var envs = os.Environ()
	cmd := exec.Command(command, args...)
	cmd.Stdout = stdoutFile
	cmd.Stderr = stderrFile
	cmd.Dir = workDir
	cmd.Env = envs
	cmd.Wait()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Just ran subprocess %d.\n", a.Cyan(cmd.Process.Pid))
	cmd.Wait()

	return logsFilePath
}

// downloadCastFile downloads the cast file and
// returns castFilePath, castFileName
func downloadCastFile(castUrl string, outputDir string) (string, string) {
	// TODO parse URLs variations
	fileUrl := castUrl + ".cast"

	// Create hash for caching
	h := sha1.New()
	h.Write([]byte(castUrl))
	castUrlHash := fmt.Sprintf("%x", h.Sum(nil))

	// Set cast file vars
	castFileName := castUrlHash + ".cast"
	castFilePath := filepath.Join(outputDir, castFileName)

	// Check file already exist
	// and return if exist
	if _, err := os.Stat(castFilePath); os.IsNotExist(err) {
		// castFilePath does not exist
	} else {
		return castFilePath, castFileName
	}

	// Download Cast file
	resp, err := http.Get(fileUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create cast file to temp file
	castFile, err := os.Create(castFilePath)
	if err != nil {
		panic(err)
	}

	// Write cast file to temp file
	_, err = io.Copy(castFile, resp.Body)
	castFile.Close()

	return castFilePath, castFileName
}

func getFileContent(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func main() {
	// Signal handler
	if runtime.GOOS == "linux" {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-c
			fmt.Println("cleaning cache")
			os.RemoveAll(cacheDirPath)
			fmt.Println("exiting...")
			os.Exit(0)
		}()
	}

	StartHTTPServer()
}

func StartHTTPServer() {
	var err error
	var serverPort = "8080"

	// Create temp dir
	cacheDirPath, err = ioutil.TempDir("", "svg-term")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(cacheDirPath)

	// Info
	fmt.Println("Using cache dir :", cacheDirPath)

	// Router
	r := mux.NewRouter()
	r.Path("/").Handler(
		http.StripPrefix(
			"/",
			http.FileServer(http.Dir("./static")),
		),
	)
	r.PathPrefix("/svg/").Handler(
		http.StripPrefix(
			"/svg/",
			http.FileServer(http.Dir(cacheDirPath+"/svg-files")),
		),
	)
	r.PathPrefix("/convert-svg/{castId}").HandlerFunc(CreateSVG)

	// HTTP server
	fmt.Println("server started:")
	fmt.Println("http://0.0.0.0:" + serverPort)
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:" + serverPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func CreateSVG(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	castId := params["castId"]

	// Create temp dir
	if cacheDirPath == "" {
		panic("cacheDirPath doesn't exist!")
	}

	// Download cast file
	castFilePath, castFileName := downloadCastFile(
		"https://asciinema.org/a/"+castId,
		cacheDirPath,
	)

	// Set svgFilePath
	if _, err := os.Stat(cacheDirPath + "/svg-files"); os.IsNotExist(err) {
		os.Mkdir(cacheDirPath+"svg-files", os.ModePerm)
	}
	svgFilePath := cacheDirPath + "/svg-files/" + castId + ".svg"

	// Run command if svgFilePath don't exist
	logsFilePath := ""
	logContent := ""
	if _, err := os.Stat(svgFilePath); os.IsNotExist(err) {
		fmt.Printf("[%s] creating SVG img...\n", castId)
		args := []string{
			"--in=" + castFilePath,
			"--out=" + svgFilePath,
		}
		logsFilePath = runCommand("/usr/local/bin/svg-term", args, castFileName)
		logContent = string(getFileContent(logsFilePath))
		fmt.Printf("[%s] created     : %s\n", castId, svgFilePath)
	} else {
		fmt.Printf("[%s] using cache : %s\n", castId, svgFilePath)
	}

	// Return CSV or error msg
	if len(logContent) != 0 {
		fmt.Printf("[%s] ERROR...\n", castId)
		fmt.Fprintf(w, "ERROR!<br>")
		fmt.Fprintf(w, string(logContent))
	} else {
		w.Header().Set("Content-Type", "json")
		fmt.Fprintf(
			w,
			fmt.Sprintf("{\"svg_url\" = \"http://%s/svg/%s.svg\"}", r.Host, castId),
		)
	}

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}
