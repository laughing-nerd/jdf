package utils

func GetHTML() string {
	return `
  <html>

<head>
  <title>jdf</title>
  <style>
    body {
      margin: 0;
      padding: 0;
      overflow: hidden;
      color: white;
      background: #1f1f29;
      overscroll-behavior: none;
      font-family: "Arial", sans-serif;
    }

    button {
      background-color: #F8F8F8;
      border: 1px solid #222222;
      border-radius: 5px;
      box-sizing: border-box;
      color: #222222;
      cursor: pointer;
      display: inline-block;
      margin: 0;
      outline: none;
      padding: 9px 10px;
      position: relative;
      text-align: center;
      width: fit-content;
      font-size: 11px;
      text-transform: uppercase;
      font-weight: 600;
    }

    button:active {
      background-color: #F7F7F7;
      border-color: #000000;
      transform: scale(.96);
    }

   .event {
      padding: 10px;
      margin-top: 2px;
      margin-bottom: 2px;
      margin-left: 10px;
      margin-right: 1%%;
      border-radius: 5px;
      background: #3f3f4a;
      width: 100%%;
      max-height: 400px;  /* Large enough value for expansion */
      overflow: hidden;
      transition: max-height 0.3s ease-out, padding 0.3s ease-out;
    }

    .event-wrapper {
      width: 100%%;
      height: auto;
      max-height: none;
      display: flex;
      align-items: start;
      gap: 5px;
      margin: 7px 0;
    }

    .event.contract {
      max-height: 18px;
      overflow: hidden;
    }

    .event.expand {
      max-height: 400px;
      overflow: auto;
    }

    #events {
      height: 90vh;
      overflow-y: auto;
      padding: 10px;
    }

    #jump-status {
      padding: 10px 30px;
      font-size: 15px;
      font-weight: 600;
      background-color: #3f3f4a;
    }

    .marker {
      border-radius: 12px;
      border: 1px solid #ccc;
      width: 15px;
      height: 15px;
      margin-top: 12px;
      cursor: pointer;
    }

    .marker-fill {
      background: #aa0000;
    }

    .toggle-expand-contract{
      background: #3f3f4a;
      color: white;
      font-size: 15px;
      width: 30px;
      padding: 10px;
    }

    .navigate {
      position: absolute;
      bottom: 10px;
      right: 10px;
      display: flex;
      justify-content: center;
      align-items: center;
      gap: 10px;
      flex-direction: column;
    }

    .navigate button {
      padding: 10px 15px;
      font-size: 15px;
      width: fit-content;
    }

    .actions {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 0 10px;
    }

    .actions > :first-child {
      display: flex;
      align-items: center;
      width: 100%%;
      gap: 10px;
    }

    .actions > :last-child {
      display: flex;
      align-items: center;
      justify-content: end;
      width: 100%%;
    }
  </style>
</head>

<body>
  <div class="actions">
    <div>
      <button onclick="alert('The developer is too lazy to add this feature 😪')">Export data</button>
      <button onclick="expandAll()">Expand all</button>
      <button onclick="contractAll()">Contract all</button>
      <button onclick="removeJumps()">Remove jumps</button>
    </div>

    <div>
      <button onclick="jump('backward')">&larr;</button>
      <p id="jump-status"></p>
      <button onclick="jump('forward')">&rarr;</button>
    </div>
  </div>

  <div class="navigate">
    <button title="Go Top" onclick="navigate('top')">&uarr;</button>
    <button title="Go Bottom" onclick="navigate('bottom')">&darr;</button>
  </div>
  
  <div id="events"></div>

  <script>
    const jumpArr = []
    let pointer = -1
    let index = -1

    jumpStatus = document.querySelector("#jump-status")
    updateJumpStatus()

    function navigate(direction) {
      const events = document.querySelector("#events")
      if (direction === "top") {
        events.scrollTo({
          top: 0,
          behavior: 'smooth'
        })
      } else {
        events.scrollTo({
          top: events.scrollHeight,
          behavior: 'smooth'
        });
      }
    }
    
    const eventSource = new EventSource('/events');
    
    eventSource.onmessage = function (event) {
      index++
      const eventsDiv = document.querySelector("#events");

      const eventWrapper = document.createElement("div");
      const newEvent = document.createElement("div");
      const marker = document.createElement("div");
      const toggleButton = document.createElement("button");

      eventWrapper.classList.add("event-wrapper");
      newEvent.classList.add("event", "expand");
      marker.classList.add("marker");
      toggleButton.classList.add("toggle-expand-contract");
      toggleButton.innerText = "-";

      newEvent.setAttribute("data-index", index);

      marker.addEventListener("click", function () {
        const eventIndex = newEvent.getAttribute("data-index");
        if (marker.classList.contains("marker-fill")) {
          marker.classList.remove("marker-fill");
          jumpArr.splice(jumpArr.indexOf(eventIndex), 1)
        } else {
          marker.classList.add("marker-fill");
          jumpArr.push(eventIndex)
        }
        console.log(jumpArr)
        updateJumpStatus()
      });

      toggleButton.addEventListener("click", function () {
        newEvent.classList.toggle("expand");
        newEvent.classList.toggle("contract");
        toggleButton.innerText = newEvent.classList.contains("expand") ? "-" : "+";
      });

      if (event.data) {
        newEvent.innerHTML = event.data;
        eventWrapper.append(marker, toggleButton, newEvent);
        eventsDiv.appendChild(eventWrapper);
      } else {
        console.error("No data received in event.");
      }
    }

    function jump(direction) {
      if (jumpArr.length === 0) return;
      if (direction === "forward") {
        pointer = Math.min(pointer + 1, jumpArr.length - 1)
      } else {
        pointer = Math.max(pointer - 1, 0)
      }
      updateJumpStatus()

      const eventDiv = document.querySelector("[data-index='" + jumpArr[pointer] + "']");
      if (eventDiv) {
        const eventsContainer = document.querySelector("#events");

        eventsContainer.scrollTo({
          top: eventDiv.offsetTop - eventsContainer.offsetTop,
          behavior: "smooth",
        });
      }
    }

    function updateJumpStatus() {
      if (jumpArr.length === 0) {
        pointer = -1
      }
      if (pointer >= jumpArr.length) {
        pointer = jumpArr.length - 1
      }
      jumpStatus.innerHTML = (pointer+1) + " / " + jumpArr.length
    }

    eventSource.onerror = function (err) {
      console.error("Error with SSE connection", err);
    };
    
    function removeJumps() {
      jumpArr.length = 0
      const markers = document.querySelectorAll('.marker-fill');
      markers.forEach(marker => marker.classList.remove('marker-fill'))
      updateJumpStatus()
    }
    

    function expandAll() {
      const eventWrappers = document.querySelectorAll('.event-wrapper');

      eventWrappers.forEach(wrapper => {
        const event = wrapper.querySelector('.event');
        const toggleButton = wrapper.querySelector('.toggle-expand-contract');
        
        event.classList.remove('contract');
        event.classList.add('expand');
        
        toggleButton.innerText = "-";
      });
    }

    function contractAll() {
      const eventWrappers = document.querySelectorAll('.event-wrapper');

      eventWrappers.forEach(wrapper => {
        const event = wrapper.querySelector('.event');
        const toggleButton = wrapper.querySelector('.toggle-expand-contract');
        
        event.classList.add('contract');
        event.classList.remove('expand');
        
        toggleButton.innerText = "+";
      });
    }

  </script>

</body>
</html>
  `
}
