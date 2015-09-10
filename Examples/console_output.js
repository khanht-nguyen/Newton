function index() {
  return 'Fully functional server';
}

function test() {
  return "in test()";
}

setRoute('/', index);
setRoute('/test', test);
startServer(':8080');
