document.addEventListener("DOMContentLoaded", function () {

  /*
  - - - - - - - - - - - - - - - - - - - - -
  - - - - - Here be local counters - - - - 
  - - - - - - - - - - - - - - - - - - - - -
  */

  let localCtr = 0;
  let localVal = document.getElementById('local-counter-value');
  localVal.textContent = localCtr;

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

  /*
  - - - - - - - - - - - - - - - - - - - - -
  - - - - - Here be network counters - - - 
  - - - - - - - - - - - - - - - - - - - - -
  */

  netVal = document.getElementById('net-counter-value');
  netConn = document.getElementById('net-connectivity');

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
    netConn.textContent = "Consecutive successful messages: " + successCounter;
    netConn.style = 'color: green'
    netVal.style = 'color: black'
    failureCounter = 0;
  }

  let log = document.getElementById('net-log')

  function outputLog(msg) {
    const node = document.createElement("p");
    x = new Date();
    node.textContent = x.toLocaleTimeString() + ' - ' + msg;
    log.insertBefore(node, log.firstChild);
    log.scrollTo(0,0);
  }

  function messageServer(com) {
    var xhr = new XMLHttpRequest({mozSystem: true});
    xhr.open("POST", "http://localhost:3000/counter");
    xhr.send(JSON.stringify({
      ID: "1", 
      Command: com
    }));
    xhr.onload = function () {
      if (xhr.status >= 200 && xhr.status < 300) {
        updateValue(JSON.parse(xhr.responseText)['Value'])
      } else {
        registerFailedConnection(com + " failed")
      }
    };
  }

  document.getElementById('net-inc').addEventListener("click", () => {
    messageServer('inc')
    outputLog('Incremented!')
  })


  document.getElementById('net-dec').addEventListener("click", () => {
    messageServer('dec')
    outputLog('Decremented!')
  })

  document.getElementById('net-flip').addEventListener("click", () => {
    messageServer('flip')
    outputLog('Flipped!')
  })

  document.getElementById('net-reset').addEventListener("click", () => {
    messageServer('reset')
    outputLog('Reset!')
  })

  window.setInterval(() => {
    var xhr = new XMLHttpRequest({mozSystem: true});
    xhr.open("POST", "http://localhost:3000/counter");
    // xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify({
      ID: "1", 
      Command: "get"
    }));
    xhr.onload = function () {
      if (xhr.status >= 200 && xhr.status < 300) {
        // console.log(JSON.parse(xhr.responseText))
        updateValue(JSON.parse(xhr.responseText)['Value'])
      } else {
        registerFailedConnection("ping failed")
      }
    };
  }, 500)
})