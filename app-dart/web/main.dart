import 'package:angular2/angular2.dart';

// These imports will go away soon:
import 'package:angular2/src/reflection/reflection.dart' show reflector;
import 'package:angular2/src/reflection/reflection_capabilities.dart' show ReflectionCapabilities;

@Component(
  selector: 'my-app'
)
@View(
  template: '<h1>Hello {{ name }} in dart</h1>'
)
class AppComponent {
  String name = 'Bob';
}

main() {
  // Temporarily needed.
  reflector.reflectionCapabilities = new ReflectionCapabilities();
  
  bootstrap(AppComponent);
}