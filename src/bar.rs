
#![feature(libc)]
extern crate libc;
use self::libc::c_void;
use self::libc::c_char;
use self::libc::int8_t;

use super::foo::Foo;

pub struct Bar {
    pub qclsinst: *mut c_void,
}

fn test_refer_foo_member_var(a0: &mut Foo) {
    println!("fff %v");
    a0.qclsinst;
}


trait Bar_trait_test_lifetime {
    fn test1(self) -> i32;
}

/*
这种写法能重现错误提示
// error: missing lifetime specifier [E0106]
impl Bar_trait_test_lifetime for (&mut Foo) {

}
*/

/*
这种写法解决了E0106的错误提示
*/
impl<'a> Bar_trait_test_lifetime for (&'a mut Foo) {
    fn test1(self) -> i32 {
        return 1;
    }
}

// char *类型转换，OK
impl<'a> Bar_trait_test_lifetime for (&'a mut str) {
    fn test1(self) -> i32 {
        self.as_ptr() as *const c_void;
        self.as_ptr() as *const c_char;
        return 1;
    }
}

// bool & 类型转换，
fn test_boolstart<'a>(a0: &'a mut bool) {
    *a0 = true;
    {
        let mut swap: int8_t = 0;
        if swap == 1 {*a0 = true;}
    }
    *a0 = {let mut bv: int8_t = 0; if bv == 1 {true} else {false}}
}

fn test_i8start<'a>(a0: &'a mut i8) {
    // a0 as *mut c_void;  // error
    a0 as *mut int8_t;  // ok
}


