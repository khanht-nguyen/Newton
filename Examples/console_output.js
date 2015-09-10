function t() {
  return [200, 'Fully functional server'];
}

setRoute('/test', t);
startServer(':8080');
