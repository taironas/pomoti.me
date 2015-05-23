# pomoti.me
a simple web app to keep track of the things you do.

## idea:

Pomoti.me is a simple web app to keep track of the things you do.

It is based on the [pomodoro technique](), a time management method that uses a timer to break down an activity into 2 intervals, a **25 mins** and a **5 mins** break.

Pomoti.me keeps track of the accomplished **pomodoros** as you go so that you can measure how much you have done over a day, week, month, year.

It might be interesting to show and/or compare your activity to someone else's in a team, or with your family. This is an interesting concept as you can see how other people work or finish an activity, how efficient they are, etc.

## stack

This app is also built as a way for me to learn new technologies. I want the frontend to use [AngularJS 2.0](https://angular.io/) and the backend will be in [Go](http://golang.org).

## set up:

get the source and build the backend.

    cd $GOPATH/src
    go get github.com/taironas/route
    cd $GOPATH/src/github.com/taironas/pomoti.me
    go get ./backend
    export PORT=8080


install [node](https://nodejs.org/download/)

set up the frontend.

    cd app/
    npm install -g tsd
    tsd query angular2 --action install

## how to make a change

you should have 2 terminals.

terminal 1:

    >  pwd
    go/src/github.com/taironas/pomoti.me
    > cd app/
    > tsc --watch -m commonjs -t es5 --emitDecoratorMetadata app.ts
    message TS6042: Compilation complete. Watching for file changes.

terminal 2:

    >  pwd
    go/src/github.com/taironas/pomoti.me
    > go get ./backend
    > backend
    2015/05/23 17:20:21 main.go:23: Listening on 8080

Open your browser and go to `localhost:8080`
