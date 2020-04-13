package pages

import (
	"html/template"
	"net/http"

	"github.com/quan-to/chevron/models"
)

const serveIndexHTML = `
<!DOCTYPE HTML>
<html>
<head>
  <meta charset="UTF-8">
  <title>Chevron Web Interface</title>
  <style>body { padding: 0; margin: 0; }</style>
  <link rel="stylesheet" href="/assets/css/bootstrap.min.css" >
  <script src="/assets/js/elm.js"></script>

</head>

<body>

<pre id="elm"></pre>


<script>
var app = Elm.Main.init({
    node: document.getElementById('elm')
  });
  </script>
</body>
</html>`

var indexTemplate = template.Must(template.New("serveIndex").Parse(serveIndexHTML))

// Serveindex serves the index HTML Page Template
func ServeIndex(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", models.MimeHTML)
	w.WriteHeader(200)
	_ = indexTemplate.Execute(w, nil)
}
