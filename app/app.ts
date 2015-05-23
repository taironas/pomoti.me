/// <reference path="typings/angular2/angular2.d.ts" />
import {Component, View, bootstrap} from 'angular2/angular2';

// Annotation section
@Component({
  selector: 'pomotime-app'
})
@View({
  template: '<h1>Hello {{ name }}</h1>'
})

// Component controller
class PomotimeComponent {
  name: string;
  
  constructor() {
    this.name = 'Pomoti.me';
  }
}
bootstrap(PomotimeComponent);
