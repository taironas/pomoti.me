import 'dart:html';
import 'package:angular2/angular2.dart';
import 'dart:convert';
import 'clock.dart';

// These imports will go away soon:
import 'package:angular2/src/reflection/reflection.dart' show reflector;
import 'package:angular2/src/reflection/reflection_capabilities.dart' show ReflectionCapabilities;

@Component(
  selector: 'my-app'
)
@View(
  template: '<p>From golang backend: /api/hello: <i>"{{hello.message}}"<i></p><p>{{clock.message}}</p>'
)

class AppComponent {
  Hello hello = new Hello();
  Clock clock = new Clock();
}


class Hello{

  String message;

  Hello(){
        HttpRequest.getString('/api/hello')
            .then((String content) {
              Map parsedMap = JSON.decode(content);
              message = parsedMap["Message"];
            })
            .catchError((Error error) {
              print(error.toString());
            });
  }
}

main() {
  // Temporarily needed.
  reflector.reflectionCapabilities = new ReflectionCapabilities();
  
  bootstrap(AppComponent);
}