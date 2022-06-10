window.addEventListener(
  "click",
  () => {
    hideOverlayAndConnectToWS();
    start();
  },
  { once: true }
);

let socket;

function hideOverlayAndConnectToWS() {
  document.getElementById("overlay").style.visibility = "hidden";

  socket = new WebSocket("ws://localhost:3000/broadcaster");

  let closeMessage = false;
  socket.onmessage = (e) => {
    if (e.data == "Iemand is de piano al aan het gebruiken!") closeMessage = true;
  };

  socket.onclose = (e) => {
    document.getElementById("overlay").style.visibility = "visible";
    document.getElementById("overlay-text").innerHTML = `Verbinding verloren${
      closeMessage ? " omdat iemand de piano al aan het gebruiken is!" : "."
    }<br>Klik om de pagina te herladen.`;

    window.addEventListener("click", () => {
      window.location.reload();
    });
  };
}

function start() {
  const getElementByNote = (note) => note && document.querySelector(`[note="${note}"]`);

  const keys = {
    A: { element: getElementByNote("C"), note: "C", octaveOffset: 0, color: "#820623" },
    W: { element: getElementByNote("C#"), note: "C#", octaveOffset: 0, color: "#aa082e" },
    S: { element: getElementByNote("D"), note: "D", octaveOffset: 0, color: "#f70a0c" },
    E: { element: getElementByNote("D#"), note: "D#", octaveOffset: 0, color: "#ff5900" },
    D: { element: getElementByNote("E"), note: "E", octaveOffset: 0, color: "#f65e0c" },
    F: { element: getElementByNote("F"), note: "F", octaveOffset: 0, color: "#ffc73c" },
    T: { element: getElementByNote("F#"), note: "F#", octaveOffset: 0, color: "#ddf82f" },
    G: { element: getElementByNote("G"), note: "G", octaveOffset: 0, color: "#9cdf50" },
    Y: { element: getElementByNote("G#"), note: "G#", octaveOffset: 0, color: "#87c146" },
    H: { element: getElementByNote("A"), note: "A", octaveOffset: 1, color: "#609c1c" },
    U: { element: getElementByNote("A#"), note: "A#", octaveOffset: 1, color: "#11beb3" },
    J: { element: getElementByNote("B"), note: "B", octaveOffset: 1, color: "#56f0e6" },
    K: { element: getElementByNote("C2"), note: "C", octaveOffset: 1, color: "#2f80ed" },
    O: { element: getElementByNote("C#2"), note: "C#", octaveOffset: 1, color: "#1057b7" },
    L: { element: getElementByNote("D2"), note: "D", octaveOffset: 1, color: "#8000e3" },
    P: { element: getElementByNote("D#2"), note: "D#", octaveOffset: 1, color: "#ff5ef4" },
    semicolon: { element: getElementByNote("E2"), note: "E", octaveOffset: 1, color: "#e337d7" },
  };

  const getHz = (note = "A", octave = 4) => {
    const A4 = 440;
    let N = 0;
    switch (note) {
      default:
      case "A":
        N = 0;
        break;
      case "A#":
      case "Bb":
        N = 1;
        break;
      case "B":
        N = 2;
        break;
      case "C":
        N = 3;
        break;
      case "C#":
      case "Db":
        N = 4;
        break;
      case "D":
        N = 5;
        break;
      case "D#":
      case "Eb":
        N = 6;
        break;
      case "E":
        N = 7;
        break;
      case "F":
        N = 8;
        break;
      case "F#":
      case "Gb":
        N = 9;
        break;
      case "G":
        N = 10;
        break;
      case "G#":
      case "Ab":
        N = 11;
        break;
    }
    N += 12 * (octave - 4);
    return A4 * Math.pow(2, N / 12);
  };

  const pressedNotes = new Map();
  let clickedKey = "";

  const playKey = (key) => {
    if (!keys[key]) {
      return;
    }

    const freq = getHz(keys[key].note, (keys[key].octaveOffset || 0) + 3);

    socket.send(JSON.stringify({ Key: key, Frequency: freq, Color: keys[key].color }));

    keys[key].element.classList.add("pressed");
    pressedNotes.set(key, "exists");
  };

  const stopKey = (key) => {
    if (!keys[key]) {
      return;
    }

    keys[key].element.classList.remove("pressed");

    pressedNotes.delete(key);
  };

  document.addEventListener("keydown", (e) => {
    const eventKey = e.key.toUpperCase();
    const key = eventKey === ";" ? "semicolon" : eventKey;

    if (!key || pressedNotes.get(key)) {
      return;
    }
    playKey(key);
  });

  document.addEventListener("keyup", (e) => {
    const eventKey = e.key.toUpperCase();
    const key = eventKey === ";" ? "semicolon" : eventKey;

    if (!key) {
      return;
    }
    stopKey(key);
  });

  for (const [key, { element }] of Object.entries(keys)) {
    element.addEventListener("mousedown", () => {
      playKey(key);
      clickedKey = key;
    });
  }

  document.addEventListener("mouseup", () => {
    stopKey(clickedKey);
  });
}
