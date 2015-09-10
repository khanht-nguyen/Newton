function t() {
  return 'Fully functional server';
}

setRoute('/test', t);
startServer(':8080');
