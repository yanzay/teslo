window.addEventListener("load", function(evt) {
  var ws = new WebSocket("ws://localhost:8080/ws");
  ws.onopen = function(e) {
    console.log("OPEN");
  };
  ws.onclose = function(e) {
    console.log("CLOSE");
  };
  ws.onmessage = function(e) {
    var message = JSON.parse(e.data);
    var el = document.getElementById(message.id)
    el.outerHTML = message.content;
  };
  ws.onerror = function(e) {
    console.log("Error: ", e.data);
  };
  var app = document.getElementById("app")

  var clickHandler = function(e) {
    console.log("Click handler");
    console.log(e.target.id);
    var parentIds = $(e.target).parents().map(function(i, el) {return el.id;}).toArray().filter(function(id) {return id !== ""});
    console.log(parentIds);
    if (e.target.id) {
      ws.send(JSON.stringify({event: "click", id: e.target.id, parents: parentIds}));
    }
  };
  var submitHandler = function(e) {
    console.log("Submit handler");
    e.preventDefault();
    var form = $(e.target).serialize();
    console.log($(e.target));
    console.log(form);
    var resp = {
      event: "submit",
      id: e.target.id,
      data: form,
    };
    console.log(resp);
    ws.send(JSON.stringify(resp));
  };
  app.addEventListener("click", clickHandler);
  app.addEventListener("submit", submitHandler);
});

