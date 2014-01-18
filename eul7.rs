use common::is_prime;
mod common;

fn main() {
    let mut x = 3;
    let mut count = 2;

    while count < 10001 {
        x += 2;
        if is_prime(x) {
            count += 1;
        }
    }

    println!("{}", x);
}
