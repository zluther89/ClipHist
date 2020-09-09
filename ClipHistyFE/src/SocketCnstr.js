module.exports = function SocketInit(connString, cb) {
  var ws = new WebSocket(connString);
  ws.onopen = function () {
    console.log("connected");
    ws.send(JSON.stringify({ message: "hello server!" }));
  };
  ws.onmessage = function (event) {
    var { message } = JSON.parse(event.data);
    console.log("Received message", message);
    cb();
  };
  ws.onerror = function (event) {
    console.log(event);
  };
  return ws;
};
