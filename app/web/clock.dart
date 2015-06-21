library clock;

import 'dart:html';
import 'dart:async';

import 'package:angular2/angular2.dart';

@Component(selector: 'clock')

@View(template: '''
<p>Clock: <i>{{ counter }}</i></p>
<button (click)="start()">Start</button>
<button (click)="stop()">Stop</button>
<button (click)="reset()">Reset</button>
<p>Pomodoro duration: <i>{{ durationPomodoro }}</i>
<input #durationpomodoro>
<button (click)="setDurationPomodoro(durationpomodoro.value)">change</button>
</p>
<p>Break duration: <i>{{ durationBreak }}</i>
<input #durationbreak>
<button (click)="setDurationBreak(durationbreak.value)">change</button>
</p>
''')

class Clock{

  int startAt;
  String durationBreak, durationPomodoro;
  String counter;
  Stopwatch watch = new Stopwatch();
  Timer timer;

  Clock(){
    startAt = (25*60);
    durationBreak = prettyPrintTime(5*60);
    durationPomodoro = prettyPrintTime(25*60);
    counter= prettyPrintTime(25*60);
  }
  
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
    var s  = startAt - watch.elapsedMilliseconds~/1000;
    
    if( s < 0){ 
      stop();
      return; 
    }
    counter = prettyPrintTime(s);
  }

  void setDurationPomodoro(string duration){
    var value = int.parse(duration, onError: (source) => null);
    if (value != null){
      startAt = value*60;
      durationPomodoro = prettyPrintTime(startAt);
      counter = durationPomodoro;
    }
  }

  void setDurationBreak(string duration){
    var value = int.parse(duration, onError: (source) => null);
    if (value != null){
      durationBreak = prettyPrintTime(value*60);
    }
  }

  String prettyPrintTime(int s){
    var m = 0;
    
    // The operator ~/ divides and returns an integer.
    if (s >= 60) { m = s ~/ 60; s = s % 60; }
    
    String minute = (m <= 9) ? '0$m' : '$m';
    String second = (s <= 9) ? '0$s' : '$s';
    return '$minute:$second';
  }
}