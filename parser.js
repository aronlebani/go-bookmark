const fs = require('fs');
const util = require('util');
const BMParser = require('bookmark-parser');

const htmlFilePath = './bookmarks_24_08_2020.html';
BMParser.readFromHTMLFile(htmlFilePath)
  .then(res => {
    const data = JSON.stringify(res, function(key, value) {
      return key == 'prevNode' ? value.id : value;
    });
    console.log(util.inspect(JSON.parse(data), { showHidden: false, depth: null }));
    fs.writeFileSync('bookmarks.json', data);
  });