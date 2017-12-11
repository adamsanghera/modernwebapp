document.addEventListener("DOMContentLoaded", function () {

  let localCtr = 0;
  let localVal = document.getElementById('local-counter-value');
  localVal.textContent = localCtr;
  netVal = document.getElementById('net-counter-value');
  netConn = document.getElementById('net-connectivity');

  document.getElementById('local-inc').addEventListener("click", () => {
    localCtr = localCtr + 1;
    localVal.textContent = localCtr;
  })

  document.getElementById('local-dec').addEventListener("click", () => {
    localCtr = localCtr - 1;
    localVal.textContent = localCtr;
  })

  document.getElementById('local-flip').addEventListener("click", () => {
    localCtr = localCtr * -1;
    localVal.textContent = localCtr;
  })

  document.getElementById('local-reset').addEventListener("click", () => {
    localCtr = 0;
    localVal.textContent = localCtr;
  })

  let failureCounter = 0;
  let successCounter = 0;

  function registerFailedConnection(failMessage) {
    console.log(failMessage)
    failureCounter += 1;
    netVal.textContent = "Unable to contact the server.";
    netVal.style = 'color: red'
    netConn.style = 'color: red'
    netConn.textContent = "Consecutive failed pings: " + failureCounter;
    successCounter = 0;
  }

  function updateValue(response) {
    console.log("received a response: " + response)
    successCounter += 1;
    netVal.textContent = response;
    netConn.textContent = "Consecutive successful pings: " + successCounter;
    netConn.style = 'color: green'
    netVal.style = 'color: black'
    failureCounter = 0;
  }

  function sendNetSignal(signal) {
    $.ajax( {
      url:"http://localhost:3000/"+signal,
      type: "post",
      success: function(response) {
        console.log("received a response: " + response)
      },
      error: (xhr) => {
        registerFailedConnection(xhr)
      }
    });
  }

  let log = document.getElementById('net-log')

  function outputLog(msg) {
    const node = document.createElement("p");
    x = new Date();
    node.textContent = x.toLocaleTimeString() + ' - ' + msg;
    log.insertBefore(node, log.firstChild);
    log.scrollTo(0,0);
  }

  document.getElementById('net-inc').addEventListener("click", () => {
    sendNetSignal("incCounter")
    outputLog('Incremented!')
  })

  document.getElementById('net-dec').addEventListener("click", () => {
    sendNetSignal("decCounter")
    outputLog('Decremented!')
  })

  document.getElementById('net-flip').addEventListener("click", () => {
    sendNetSignal("flipCounter")
    outputLog('Flipped!')
  })

  document.getElementById('net-reset').addEventListener("click", () => {
    sendNetSignal("resetCounter")
    outputLog('Reset!')
  })

  window.setInterval((e) => {
    data = JSON.stringify({
      ID: "1", 
      Command: "get"
    });
    var xhr = new XMLHttpRequest({mozSystem: true});
    xhr.open("POST", "http://localhost:3000/counter");
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.send(data);
    // xhr.onreadystatechange = function () {
    //   if (xhr.status >= 200 && xhr.status < 300) {
    //     console.log(xhr.responseText);
    //   } else {
    //     console.warn(xhr);
    //   }
    // };
  }, 500)
})