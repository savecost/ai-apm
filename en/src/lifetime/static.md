# &'static and T: 'static
`'static` is a reserved lifetime name, you might have encountered it several times:
```rust
// A reference with 'static lifetime:
let s: &'static str = "hello world";

// 'static as part of a trait bound:
fn generic<T>(x: T) where T: 'static {}
```

Though they are all `'static`, but subtly different.

## &'static
As a reference lifetime, `&'static` indicates the data pointed to by the reference lives as long as the running program. But it can still be coerced to a shorter lifetime.



1. 🌟🌟 There are several ways to make a variable with `'static` lifetime, two of them are stored in the read-only memory of the binary.

```rust,editable

/* Fill in the blank in two ways */
fn main() {
    __;
    need_static(v);

    println!("Success!")
}

fn need_static(r : &'static str) {
    assert_eq!(r, "hello");
}
```

2. 🌟🌟🌟🌟 Another way to make `'static` lifetime is using `Box::leak`
```rust,editable
#[derive(Debug)]
struct Config {
    a: String,
    b: String,
}
static mut config: Option<&mut Config> = None;

/* Make it work without changing the function signatures of `init`*/
fn init() -> Option<&'static mut Config> {
    Some(&mut Config {
        a: "A".to_string(),
        b: "B".to_string(),
    })
}


fn main() {
    unsafe {
        config = init();
        if config.is_none() {panic!("`None` isn't solution to this exercise")};

        println!("{:?}",config)
    }
}
```

3. 🌟 `&'static` only indicates that the data can live forever, not the reference. The latter one will be constrained by its scope.
```rust,editable
fn main() {
    {
        // Make a `string` literal and print it:
        let static_string = "I'm in read-only memory";
        println!("static_string: {}", static_string);

        // When `static_string` goes out of scope, the reference
        // can no longer be used, but the data remains in the binary.
    }

    println!("static_string reference remains alive: {}", static_string);
}
```

4. `&'static` can be coerced to a shorter lifetime.

**Example**
```rust,editable
// Make a constant with `'static` lifetime.
static NUM: i32 = 18;

// Returns a reference to `NUM` where its `'static`
// lifetime is coerced to that of the input argument.
fn coerce_static<'a>(_: &'a i32) -> &'a i32 {
    &NUM
}

fn main() {
    {
        // Make an integer to use for `coerce_static`:
        let lifetime_num = 9;

        // Coerce `NUM` to lifetime of `lifetime_num`:
        let coerced_static = coerce_static(&lifetime_num);

        println!("coerced_static: {}", coerced_static);
    }

    println!("NUM: {} stays accessible!", NUM);
}
```



##  T: 'static
As a trait bound, it means the type does not contain any non-static references. Eg. the receiver can hold on to the type for as long as they want and it will never become invalid until they drop it.


It's important to understand this means that any owned data always passes a `'static `lifetime bound, but a reference to that owned data generally does not.


5. 🌟🌟
```rust,editable
/* Make it work */
use std::fmt::Debug;

fn print_it<T: Debug + 'static>( input: T) {
    println!( "'static value passed in is: {:?}", input );
}

fn print_it1( input: impl Debug + 'static ) {
    println!( "'static value passed in is: {:?}", input );
}


fn print_it2<T: Debug + 'static>( input: &T) {
    println!( "'static value passed in is: {:?}", input );
}

fn main() {
    // i is owned and contains no references, thus it's 'static:
    let i = 5;
    print_it(i);

    // oops, &i only has the lifetime defined by the scope of
    // main(), so it's not 'static:
    print_it(&i);

    print_it1(&i);

    // but this one WORKS !
    print_it2(&i);
}
```


6. 🌟🌟🌟
```rust,editable
use std::fmt::Display;

fn main() {
  let mut string = "First".to_owned();

  string.push_str(string.to_uppercase().as_str());
  print_a(&string);
  print_b(&string);
  print_c(&string); // Compilation error
  print_d(&string); // Compilation error
  print_e(&string);
  print_f(&string);
  print_g(&string); // Compilation error
}

fn print_a<T: Display + 'static>(t: &T) {
  println!("{}", t);
}

fn print_b<T>(t: &T)
where
  T: Display + 'static,
{
  println!("{}", t);
}

fn print_c(t: &'static dyn Display) {
  println!("{}", t)
}

fn print_d(t: &'static impl Display) {
  println!("{}", t)
}

fn print_e(t: &(dyn Display + 'static)) {
  println!("{}", t)
}

fn print_f(t: &(impl Display + 'static)) {
  println!("{}", t)
}

fn print_g(t: &'static String) {
  println!("{}", t);
}
```

> You can find the solutions [here](https://github.com/sunface/rust-by-practice/blob/master/solutions/lifetime/static.md)(under the solutions path), but only use it when you need it :)
