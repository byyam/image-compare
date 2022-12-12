const sharp = require('sharp');


sharp('input.jpg')
  .resize(300, 200)
  .toFile('output.jpg', function(err) {
    // output.jpg is a 300 pixels wide and 200 pixels high image
    // containing a scaled and cropped version of input.jpg
  });

const express = require('express')

const app = express()

var bodyParser = require('body-parser');
app.use(bodyParser.json({'limit': '1000kb'}));
  

app.post('/upload-multiple-files', (req, res) => {
  console.log("upload-multiple-files")
  res.end('Hello, World!')
})
  
app.post('/', (req, res) => {
  console.log("root")
  res.end('Hello, World!')
})

app.listen(3000)


