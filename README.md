# pomoti.me
a simple web app to keep track of the things you do.

## idea:

Pomoti.me is a simple web app to keep track of the things you do.

It is based on the [pomodoro technique](), a time management method that uses a timer to break down an activity into 2 intervals, a **25 mins** and a **5 mins** break.

Pomoti.me keeps track of the accomplished **pomodoros** as you go so that you can measure how much you have done over a day, week, month, year.

It might be interesting to show and/or compare your activity to someone else's in a team, or with your family. This is an interesting concept as you can see how other people work or finish an activity, how efficient they are, etc.

## stack

This app is also built as a way for me to learn new technologies. I want the frontend to be done in [Dart](http://dartlang.org) [AngularJS 2.0](https://angular.io/) and the backend will be in [Go](http://golang.org).

## set up:

get the source and build the backend.

    cd $GOPATH/src
    go get github.com/taironas/route
    cd $GOPATH/src/github.com/taironas/pomoti.me
    go get ./backend
    export PORT=8080

###set up and run the frontend in dart.

* download [dart](https://www.dartlang.org/downloads/)

* build the frontend

~~~
    >cd app-dart
    >pub get
    >pub build
~~~

* run in javascript (prod mode, open in any navigator)

~~~
    >pwd
    github.com/taironas/pomiti.me
    >go get ./backend
    >backend -dart -prod
~~~~


* run in dart (dev mode, open in chromium)


~~~
    >pwd
    github.com/taironas/pomoti.me
    >go get ./backend
    >backend -dart
    2015/05/25 18:50:57 main.go:43: Listening on 8080
~~~

Open `localhost:8080` in chromium.

Note: The good thing with the **dev mode** is that you don't need to build anything, to see the changes, only update the browser.
