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

  int startPomodoroAt, startRestAt, currentDuration;
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
    currentDuration = startPomodoroAt;
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
    counter = prettyPrintTime(currentDuration);
  }

  void updateTime(Timer _) {
    var s = watch.elapsedMilliseconds~/1000;
    counter = prettyPrintTime(s);
  }

  void updateState(){
    if (currentState == ClockState.pomodoro){
      currentState = ClockState.rest;
      return;
    }
    currentState = ClockState.pomodoro;
  }

  void updateTimeRemaining(Timer _) {
    var s  = currentDuration - watch.elapsedMilliseconds~/1000;
    
    if( s < 0){ 
      updateState();
      setCounter();
      reset();
      return; 
    }
    counter = prettyPrintTime(s);
  }

  void setDurationPomodoro(string duration){
    var value = int.parse(duration, onError: (source) => null);
    if (value != null){
      startPomodoroAt = value*60;
      durationPomodoro = prettyPrintTime(startPomodoroAt);
      if (currentState == ClockState.pomodoro){
        counter = durationPomodoro;
        currentDuration = startPomodoroAt;
      }
    }
  }

  void setDurationRest(string duration){
    var value = int.parse(duration, onError: (source) => null);
    if (value != null){
      startRestAt = value*60;
      durationRest = prettyPrintTime(startRestAt);
      if (currentState == ClockState.rest){
        counter = durationRest;
        currentDuration = startPomodoroAt;
      }      
    }
  }

  void setCounter(){
    switch(currentState){
      case ClockState.pomodoro:
        counter = prettyPrintTime(startPomodoroAt);
        currentDuration = startPomodoroAt;
        break;
      case ClockState.rest:
        counter = prettyPrintTime(startRestAt);
        currentDuration = startRestAt;
        break;
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