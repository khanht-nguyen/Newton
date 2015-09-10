function index() {
  writeMessage("")
}

function test() {
  writeMessage("in test()")
}

setRoute('/', index);
setRoute('/test', test);
startServer(':8080');
