fn main() {
    let mut a = 1;
    let mut b = 2;
    let mut sum = 0;
    let mut tmp: int;

    while b < 4_000_000 {
        if b % 2 == 0 {
            sum += b;
        }
        tmp = a;
        a = b;
        b = a + tmp;
    }

    println!("sum: {}", sum);
}
