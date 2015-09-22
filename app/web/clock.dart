library clock;

import 'dart:html';
import 'dart:async';
import 'package:intl/intl.dart'; // datetime formatter

import 'package:angular2/angular2.dart';

enum ClockState{
  pomodoro,
  rest
}

class Period{
  DateTime start;
  DateTime end;
  string label;

  Period(this.start, this.end, this.label);
}

@Component(selector: 'clock')

@View(template: '''
<p>Clock: <i>{{ counter }}</i></p>
<button (click)="start()">Start</button>
<button (click)="stop()">Stop</button>
<button (click)="reset()">Reset</button>
<button (click)="skip()">Skip</button>
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
<p>History:</p>
<ul>
    <li *ng-for="#period of sortPeriods()">
        {{formatTime(period.start)}} | {{formatTime(period.end)}} | {{period.label}}
    </li>
</ul>
''', directives: const [NgFor, NgIf])

class Clock{

  int startPomodoroAt, startRestAt, currentDuration;
  String durationRest, durationPomodoro;
  String counter;
  Stopwatch watch = new Stopwatch();
  Timer timer;
  List<Period> periods;

  ClockState currentState;
  DateFormat formatter;

  Clock(){
    startPomodoroAt = (25*60);
    startRestAt = (5*60);
    durationRest = prettyPrintTime(startRestAt);
    durationPomodoro = prettyPrintTime(startPomodoroAt);
    counter= prettyPrintTime(25*60);
    currentState = ClockState.pomodoro;
    currentDuration = startPomodoroAt;
    periods = new List();
    formatter = new DateFormat('MM-dd-yyyy HH:mm:ss');
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

  void skip(){
    updateState();
    setCounter();
    reset();
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
      addToHistory();
      reset();
      return; 
    }
    counter = prettyPrintTime(s);
  }

  void addToHistory(){
    switch(currentState){
      case ClockState.pomodoro:
        var now = new DateTime.now();
        var start = now.subtract(new Duration(seconds: startRestAt));
        var end = now;
        var formatter = new DateFormat('dd/MM/yyyy');

        periods.add(new Period(start, end,"rest"));
        var data = { 'type' : 'pomodoro', 'start': formatter.format(start), 'end': formatter.format(end) };
        HttpRequest.postFormData('/api/period/create', data).then((HttpRequest response) {
          print("Response status: ${response.status}");
          print("Response body: ${response.response}");
        });
        break;
      case ClockState.rest:
        var now = new DateTime.now();
        var start = now.subtract(new Duration(seconds: startPomodoroAt));
        var end = now;
        periods.add(new Period(start, end,"pomodoro"));
        var formatter = new DateFormat('dd/MM/yyyy');
        var data = { 'type' : 'pomodoro', 'start': formatter.format(start), 'end': formatter.format(end) };
        HttpRequest.postFormData('/api/period/create', data).then((HttpRequest response) {
          print("Response status: ${response.status}");
          print("Response body: ${response.response}");
        });
        break;
    }
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

  List<Period> sortPeriods(){
    periods.sort((x, y) => x.end.compareTo(y.end));
    return periods;
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

  string formatTime(t){
    return formatter.format(t);
  }
}
