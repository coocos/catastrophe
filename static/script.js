const socket = new WebSocket("ws://localhost:8080/websocket");
socket.onmessage = message => {
  console.log(JSON.parse(message.data));
};
