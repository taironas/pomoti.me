library clock;

import 'dart:html';
import 'package:angular2/angular2.dart';

@Component(selector: 'clock')

@View(template: '''<p>Clock: <i>{{ message }}</i></p>''')

class Clock{

  String message;

  Clock(){
    message = "clock here";
  }
}