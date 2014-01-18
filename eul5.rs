fn gcd(a: int, b: int) -> int {
    if a == 0 && b == 0 {
        return 0; // should really be an error
    }

    if a == 0 {
        return b;
    } else if b == 0 {
        return a;
    }

    if a < b {
        return gcd(b, a);
    }

    gcd(b, a % b)
}

fn main() {
    let mut res = 1;
    let mut i = 20;
    let mut d: int;
    while i > 1 {
        d = gcd(res, i);
        if d != i {
            res *= i / d;
        }
        i -= 1;
    }

    println!("{}", res);
}
