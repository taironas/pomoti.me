library clock;

import 'dart:html';
import 'dart:async';

import 'package:angular2/angular2.dart';

enum ClockState{
  pomodoro,
  rest
}

@Component(selector: 'clock')

@View(template: '''
<p>Clock: <i>{{ counter }}</i></p>
<button (click)="start()">Start</button>
<button (click)="stop()">Stop</button>
<button (click)="reset()">Reset</button>
<p *ng-if="isPomotime()">Currently in pomodoro</p>
<p *ng-if="!isPomotime()">Currently in rest</p>
<p>Pomodoro duration: <i>{{ durationPomodoro }}</i>
<input #durationpomodoro>
<button (click)="setDurationPomodoro(durationpomodoro.value)">change</button>
</p>
<p>Rest duration: <i>{{ durationRest }}</i>
<input #durationrest>
<button (click)="setDurationRest(durationrest.value)">change</button>
</p>
''', directives: const [NgIf])

class Clock{

  int startPomodoroAt, startRestAt;
  String durationRest, durationPomodoro;
  String counter;
  Stopwatch watch = new Stopwatch();
  Timer timer;

  ClockState currentState;

  Clock(){
    startPomodoroAt = (25*60);
    startRestAt = (5*60);
    durationRest = prettyPrintTime(startRestAt);
    durationPomodoro = prettyPrintTime(startPomodoroAt);
    counter= prettyPrintTime(25*60);
    currentState = ClockState.pomodoro;
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
    var s  = startPomodoroAt - watch.elapsedMilliseconds~/1000;
    
    if( s < 0){ 
      stop();
      return; 
    }
    counter = prettyPrintTime(s);
  }

  void setDurationPomodoro(string duration){
    var value = int.parse(duration, onError: (source) => null);
    if (value != null){
      startPomodoroAt = value*60;
      durationPomodoro = prettyPrintTime(startPomodoroAt);
      counter = durationPomodoro;
    }
  }

  void setDurationRest(string duration){
    var value = int.parse(duration, onError: (source) => null);
    if (value != null){
      durationRest = prettyPrintTime(value*60);
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

  bool isPomotime(){
    return currentState == ClockState.pomodoro;
  }
}