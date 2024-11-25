1. 

```rust
fn main() {
    let color = String::from("green");

    let print = || println!("`color`: {}", color);

    print();
    print();

    println!("{}",color);
}
```

2. 

```rust
fn main() {
    let mut count = 0;

    let mut inc = move || {
        count += 1;
        println!("`count`: {}", count);
    };

    inc();


    let _reborrow = &count; 

    inc();

    // The closure no longer needs to borrow `&mut count`. Therefore, it is
    // possible to reborrow without an error
    let _count_reborrowed = &mut count; 

    assert_eq!(count, 0);
}
```

3. 

```rust
fn main() {
     // A non-copy type.
     let movable = Box::new(3);

     // A copy type would copy into the closure leaving the original untouched.
     // A non-copy must move and so `movable` immediately moves into
     // the closure.
     let consume = || {
         println!("`movable`: {:?}", movable);
         take(movable);
     };

     consume();
     // consume();
}

fn take<T>(_v: T) {

}
```

```rust
fn main() {
     // A non-copy type.
     let movable = Box::new(3);

     // A copy type would copy into the closure leaving the original untouched.
     // A non-copy must move and so `movable` immediately moves into
     // the closure.
     let consume = || {
         println!("`movable`: {:?}", movable);
         take(&movable);
     };

     consume();
     consume();
}

fn take<T>(_v: &T) {

}
```

4. 

```rust
fn main() {
    let example_closure = |x| x;

    let s = example_closure(String::from("hello"));

    /* Make it work, only changeg the following line */
    let n = example_closure(String::from("5"));
}
```

5. 

```rust
fn fn_once<F>(func: F)
where
    F: Fn(usize) -> bool,
{
    println!("{}", func(3));
    println!("{}", func(4));
}

fn main() {
    let x = vec![1, 2, 3];
    fn_once(|z|{z == x.len()})
}
```

```rust
fn fn_once<F>(func: F)
where
    F: Copy + FnOnce(usize) -> bool,// 改动在这里
{
    println!("{}", func(3));
    println!("{}", func(4));
}

fn main() {
    let x = vec![1, 2, 3];
    fn_once(|z|{z == x.len()})
}
```

6. 

```rust
fn main() {
    let mut s = String::new();

    let update_string = |str| s.push_str(str);

    exec(update_string);

    println!("{:?}",s);
}

fn exec<'a, F: FnMut(&'a str)>(mut f: F)  {
    f("hello")
}
```

7. 

```rust
// A function which takes a closure as an argument and calls it.
// <F> denotes that F is a "Generic type parameter"
fn apply<F>(f: F) where
    // The closure takes no input and returns nothing.
    F: FnOnce() {

    f();
}

// A function which takes a closure and returns an `i32`.
fn apply_to_3<F>(f: F) -> i32 where
    // The closure takes an `i32` and returns an `i32`.
    F: Fn(i32) -> i32 {

    f(3)
}

fn main() {
    use std::mem;

    let greeting = "hello";
    // A non-copy type.
    // `to_owned` creates owned data from borrowed one
    let mut farewell = "goodbye".to_owned();

    // Capture 2 variables: `greeting` by reference and
    // `farewell` by value.
    let diary = || {
        // `greeting` is by reference: requires `Fn`.
        println!("I said {}.", greeting);

        // Mutation forces `farewell` to be captured by
        // mutable reference. Now requires `FnMut`.
        farewell.push_str("!!!");
        println!("Then I screamed {}.", farewell);
        println!("Now I can sleep. zzzzz");

        // Manually calling drop forces `farewell` to
        // be captured by value. Now requires `FnOnce`.
        mem::drop(farewell);
    };

    // Call the function which applies the closure.
    apply(diary);

    // `double` satisfies `apply_to_3`'s trait bound
    let double = |x| 2 * x;

    println!("3 doubled: {}", apply_to_3(double));
}
```

8. 

```rust
fn main() {
    let mut s = String::new();

    let update_string = |str| -> String {s.push_str(str); s };

    exec(update_string);
}

fn exec<'a, F: FnOnce(&'a str) -> String>(mut f: F) {
    f("hello");
}
```

9. 

```rust
// Define a function which takes a generic `F` argument
// bounded by `Fn`, and calls it
fn call_me<F: Fn()>(f: F) {
    f();
}

// Define a wrapper function satisfying the `Fn` bound
fn function() {
    println!("I'm a function!");
}

fn main() {
    // Define a closure satisfying the `Fn` bound
    let closure = || println!("I'm a closure!");

    call_me(closure);
    call_me(function);
}
```

10. 

```rust
/* Fill in the blank and fix the error */
// You can also use `impl FnOnce(i32) -> i32`
fn create_fn() -> impl Fn(i32) -> i32 {
    let num = 5;

    move |x| x + num
}


fn main() {
    let fn_plain = create_fn();
    fn_plain(1);
}
```

```rust
/* Fill in the blank and fix the error */
fn create_fn() -> Box<dyn Fn(i32) -> i32> {
    let num = 5;

    // how does the following closure capture the environment variable `num`
    // &T, &mut T, T ?
    Box::new(move |x| x + num)
}


fn main() {
    let fn_plain = create_fn();
    fn_plain(1);
}
```

11. 

```rust
// Every closure has its own type. Even if one closure has the same representation as another, their types are different.
fn factory(x:i32) -> Box<dyn Fn(i32) -> i32> {

    let num = 5;

    if x > 1{
        Box::new(move |x| x + num)
    } else {
        Box::new(move |x| x + num)
    }
}

fn main() {}
```
