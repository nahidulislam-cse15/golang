Date :Date : 14 March 2022
# Interface
Interfaces in Go provide a way to specify the behavior of an object .It  allowes polymorphism.A value can be more than one type.<br>
"Interfaces are types that just declare behavior. This behavior is never implemented by the
 interface type directly, but instead by user-defined types via methods. When a
 user-defined type implements the set of methods declared by an interface type, values of
 the user-defined type can be assigned to values of the interface type. This assignment
 stores the value of the user-defined type into the interface value.

 If a method call is made against an interface value, the equivalent method for the
 stored user-defined value is executed. Since any user-defined type can implement any
 interface, method calls against an interface value are polymorphic in nature. The
 user-defined type in this relationship is often called a concrete type, since interface values
 have no concrete behavior without the implementation of the stored user-defined value."
  - Bill Kennedy
 
  Interface types express generalizations or abstractions about the behaviors of other types.
By generalizing, interfaces let us write functions that are more flexible and adaptable
because they are not tied to the details of one particular implementation.
<br>


In most object-oriented languages, methods are associated with class, but in Go, methods associate with a struct type.
In Go, explicit declaration of interface implementation is not required. You just need to implement the methods defined in the interface into your struct type where you want to implement an interface type.

Basic Syntax of interface 
```
type interface_name interface{
    // methods
} 

```
This is called EMBEDDING an interface. 
Conceptually, a value of an interface type, or INTERFACE VALUE, has two components,
- a CONCRETE TYPE and a
- VALUE OF THAT TYPE.

These are called the interface's
- DYNAMIC TYPE and
- DYNAMIC VALUE.

# Polymorphism
  "Polymorphism is the ability to write code that can take on different behavior through the
 implementation of types. Once a type implements an interface, an entire world of
 functionality can be opened up to values of that type."
 - Bill Kennedy
 <br>

poly means many <br>
we see the override of the interface method speak individually 