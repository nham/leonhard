// only works for strings where all characters are single bytes
// I can't figure out if there's any builtin method that handles this
fn is_pal(x: &str) -> bool {
    let mut i = 0;
    let mut t1;
    let mut t2;

    while i < x.len() {
        t1 = x[i];
        t2 = x[x.len() - i - 1];

        if t1 != t2 {
            return false;
        }

        i += 1;
    }

    return true;
    
}

// 19 19
// -----
// 19 18
// -----
// 18 18
// 19 17
// -----
// 18 17
// 19 16
// -----
// 17 17
// 18 16
// 19 15
fn main() {
    let mut x = 999 + 999;
    let mut a: int;
    let mut b: int;
    let mut tmp: int;
    'foo: loop {
        a = x / 2;
        b = x - a;
        while b < 1000 {
            tmp = a*b;
            println!("{} {} {}", a, b, tmp);
            if is_pal(tmp.to_str()) {
                break 'foo;
            }
            b += 1;
            a -= 1;
        }
        println!("---------");
        x -= 1;
    }
    //println!("{}", is_pal(y));

    /*
    let mut x = 0;
    'foo: loop {
        if x == 5 {
            break 'foo;
        }
        x += 1;
     }
     */

}
