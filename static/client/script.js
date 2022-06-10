/**
 * Verbind met de dirigent server.
 */
function ConnectToWS() {
  let socket = new WebSocket(`ws://localhost:3000/middleman`);

  listenToWS(socket);
}

/**
 * Verbreek de WebSocket connectie.
 * @param {Websocket} socket
 */
function DisconnectFromWS(socket) {
  socket.close();
}

/**
 * Handler voor de verschillende Websocket.on events.
 * @param {Websocket} socket
 */
function listenToWS(socket) {
  const switchButton = document.getElementById("connect-button");

  /**
   * Functie wordt uitgevoerd zodra er succesvol een WebSocket verbinding geopend is.
   * @param {Response} e
   */
  socket.onopen = () => {
    switchButton.innerHTML = "Verbonden!<br><br>Verbreek verbinding";
    switchButton.style.backgroundColor = "green";
    switchButton.onclick = () => {
      DisconnectFromWS(socket);
    };
  };

  /**
   * Functie wordt uitgevoerd wanneer de WebSocket verbinding gesloten wordt.
   */
  socket.onclose = () => {
    switchButton.onclick = () => {
      ConnectToWS();
    };
    switchButton.innerHTML = "Maak opnieuw verbinding.";
    switchButton.style.backgroundColor = "gray";
  };

  /**
   * Functie wordt uitgevoerd wanneer er een message via de WebSocket verbinding binnenkomt.
   * @param {Response} e
   */
  socket.onmessage = (e) => {
    const { Key: _key, Color: _color, Frequency: _freq } = JSON.parse(e.data);

    document.getElementsByTagName("body")[0].style.backgroundColor = _color;
    document.documentElement.style.setProperty("--color", _color);

    playNote(_key, _freq);
  };

  const NotesToPlay = new Map();

  /**
   * Speel een toon af met behulp van de Audio API
   * @param {string} note
   */
  function playNote(key, frequency) {
    const audioContext = new (window.AudioContext || window.webkitAudioContext)();

    const osc = audioContext.createOscillator();
    const noteGainNode = audioContext.createGain();
    noteGainNode.connect(audioContext.destination);

    noteGainNode.gain.value = 0.00001;
    const setAttack = () =>
      noteGainNode.gain.exponentialRampToValueAtTime(0.5, audioContext.currentTime + 0.01);
    const setDecay = () =>
      noteGainNode.gain.exponentialRampToValueAtTime(0.001, audioContext.currentTime + 1);
    const setRelease = () =>
      noteGainNode.gain.exponentialRampToValueAtTime(0.00001, audioContext.currentTime + 2);

    setAttack();
    setDecay();
    setRelease();

    osc.connect(noteGainNode);
    osc.type = "triangle";

    if (Number.isFinite(frequency)) {
      osc.frequency.value = frequency;
    }

    NotesToPlay.set(key, osc);
    NotesToPlay.get(key).start();
  }
}
