var f = 0;

function index() {
  setHeader('Content-Type', 'text/html');
  f++;
  writeMessage("<!doctype html> <html lang='en'> <head> <meta charset='utf-8'> <title>The HTML5 Herald</title> <meta name='description' content='Test HTML'> <meta name='author' content='test'><!--[if lt IE 9]> <script src='http://html5shiv.googlecode.com/svn/trunk/html5.js'></script> <![endif]--> </head> <body>"+ f +"</body> </html>");
}

function test() {
  writeMessage("in test()")
}

setRoute('/', index);
setRoute('/test', test);
startServer(':8080');
