package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// enableGzip adds Gzip compression middleware to the Gin engine.
func enableGzip(engine *gin.Engine) {
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
}

func main() {
	// Configure logging format
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Initialize Gin engine
	engine := gin.New()

	// Uncomment the following line to enable Gzip compression
	// enableGzip(engine)

	// Define routes
	engine.POST("/query", handleDownload)
	engine.GET("/", showHomepage)

	// Start server
	if err := engine.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// handleDownload handles the file download request.
func handleDownload(ctx *gin.Context) {
	// Get form data
	reqData, exists := ctx.GetPostForm("json")
	if !exists {
		ctx.String(400, "Missing required JSON form data")
		return
	}

	// Placeholder for query logic (if any)
	log.Printf("Received JSON data: %s", reqData)

	// Prepare file for download
	now := time.Now()
	fileName := now.Format("20060102_150405.csv")
	ctx.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	ctx.Writer.Header().Set("Transfer-Encoding", "chunked")
	ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	ctx.Writer.WriteHeader(200)

	// Simulate streaming data generation
	for i := 0; i < 100; i++ {
		line := fmt.Sprintf("\"data_%d\"\t\"%s\"\n", i+1, time.Now().Format("2006-01-02 15:04:05"))
		_, err := ctx.Writer.WriteString(line)
		if err != nil {
			log.Printf("Error writing data: %v", err)
			break
		}
		ctx.Writer.Flush()
		time.Sleep(500 * time.Millisecond)
	}
}

// showHomepage serves the homepage with a file download simulation.
func showHomepage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.String(200, `
<html>
<body>
    <h1>File Download Example</h1>
    <p>Click the link below to download a file:</p>
    <a href="javascript:download()">Download File</a>
    <script>
        function download() {
            var handle = window.open("about:blank", "my_download_window");
            document.forms[0].target = "my_download_window";
            document.forms[0].json.value = "test_data";
            document.forms[0].submit();
        }
    </script>
    <form action="/query" method="POST" enctype="multipart/form-data">
        <input type="hidden" name="json" value="" />
    </form>
</body>
</html>
`)
}
