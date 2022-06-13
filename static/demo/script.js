let socket = new WebSocket("ws://dirigent.hanskazan.space/demo");

// Connection established
socket.onopen = (e) => {
  document.getElementById("displayState").innerHTML = "Verbonden!";
};

// Message received
socket.onmessage = (e) => {
  document.getElementsByTagName("body")[0].style.backgroundColor = e.data;
};

// Connection closed
socket.onclose = (e) => {
  if (e.wasClean) {
    document.getElementById(
      "displayState"
    ).innerHTML = `Verbinding verloren...<br>Reden: ${e.reason}`;
  } else {
    document.getElementById("displayState").innerHTML = `Verbinding onverwachts verloren...`;
  }
  document.getElementsByTagName("body")[0].style.backgroundColor = "white";
};

// Error
socket.onerror = (e) => {
  document.getElementById("displayState").innerHTML = `Error! ${error.message}`;
};
