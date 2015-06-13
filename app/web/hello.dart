library hello;

import 'dart:html';
import 'package:angular2/angular2.dart';
import 'dart:convert';

@Component(
  selector: 'hello'
)
@View(
  template: '<p>From golang backend: /api/hello: <i>"{{message}}"<i></p>'
)


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