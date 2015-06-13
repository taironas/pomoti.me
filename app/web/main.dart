import 'dart:html';
import 'package:angular2/angular2.dart';
import 'dart:convert';
import 'clock.dart';
import 'hello.dart';

// These imports will go away soon:
import 'package:angular2/src/reflection/reflection.dart' show reflector;
import 'package:angular2/src/reflection/reflection_capabilities.dart' show ReflectionCapabilities;

main() {
  // Temporarily needed.
  reflector.reflectionCapabilities = new ReflectionCapabilities();
  
  bootstrap(Hello);
  bootstrap(Clock);
}