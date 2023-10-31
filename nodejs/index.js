const fs = require("fs");
const http = require("http");
// const util = require("util");

const server = http.createServer(async (req, res) => {
  if (req.url === "/") {
    res.writeHead(200, {
      "Content-Type": "text/html;charset=utf-8",
      "Cache-Control": "no-cache, no-store, must-revalidate, no-transform",
    });

    const file = fs.readFileSync("./index.html");

    res.write(file);
  }

  if (req.url === "/provider") {
    res.writeHead(301, {
      "Content-Type": "text/html;charset=utf-8",
      Location: "/callback?client_id=1&token=1001&provider=nodesjs",
      "Cache-Control": "no-cache, no-store, must-revalidate, no-transform",
    });
  }

  if (req.url.includes("/callback")) {
    res.writeHead(200, {
      "Set-Cookie": `session=nedya-amril-prakasa-nodejs`,
      "Content-Type": "application/json;charset=utf-8",
      "Cache-Control": "no-cache, no-store, must-revalidate, no-transform",
      "Cache-Control": "no-cache, no-store, must-revalidate, no-transform",
    });

    const responseStatus = {
      message: "welcome Nakama in nodeJS",
      status: 'OK'
    };

    const jsonContent = JSON.stringify(responseStatus);

    res.write(jsonContent);
  }

  res.end();
});

server.listen(3000, () => {
  console.log("app starting at port 3000");
});
