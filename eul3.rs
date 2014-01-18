use common::is_prime;
mod common;

fn main() {
    let mut i = 2;
    let mut x = 600851475143;

    let mut mpf = 0;
    while x > 1 {
        if x % i == 0 {
            x = x / i;

            if is_prime(i) && i > mpf {
                mpf = i;
            }
        } else {
            i += 1;
        }
    }

    println!("{}", mpf);

}
