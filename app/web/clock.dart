library clock;

import 'dart:html';
import 'dart:async';

import 'package:angular2/angular2.dart';

@Component(selector: 'clock')

@View(template: '''
<p>Clock: <i>{{ counter }}</i></p>
<button (click)="start()" id="startButton">Start</button>
<button (click)="stop()" id="startButton">Stop</button>
<button (click)="reset()" id="startButton">Reset</button>
''')

class Clock{
  @observable String counter='25:00';
  Stopwatch watch = new Stopwatch();
  Timer timer;
  
  void start() {
    watch.start();
    var oneSecond = new Duration(seconds:1);
    timer = new Timer.periodic(oneSecond, updateTimeRemaining);
  }

  void stop() {
    watch.stop();
    timer.cancel();
  }
  
  void reset() {
    watch.reset();
    counter = '00:00';
  }

  void updateTime(Timer _) {
    var s = watch.elapsedMilliseconds~/1000;
    var m = 0;
    
    // The operator ~/ divides and returns an integer.
    if (s >= 60) { m = s ~/ 60; s = s % 60; }
    
    String minute = (m <= 9) ? '0$m' : '$m';
    String second = (s <= 9) ? '0$s' : '$s';
    counter = '$minute:$second';
  }

  void updateTimeRemaining(Timer _) {
    var s  = (25*60) - watch.elapsedMilliseconds~/1000;
    var m = 0;
    
    if( s < 0){ 
      stop();
      return; 
    }
    // The operator ~/ divides and returns an integer.
    if (s >= 60) { m = s ~/ 60; s = s % 60; }
    
    String minute = (m <= 9) ? '0$m' : '$m';
    String second = (s <= 9) ? '0$s' : '$s';
    counter = '$minute:$second';
  }

    

}