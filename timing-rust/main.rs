use std::env;
use std::thread::{sleep};
use chrono::prelude::*;
use std::time::Duration;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() == 1 {
        println!("参数：\n   now 立刻执行\n   执行时间的小时数\n   执行时间的小时数 执行时间的分钟数");
    } else {
        if args[1] == "now" {
            return;
        } else if args.len() == 3 {
            loop {
                if Local::now().hour() >= args[1].parse().expect("error") {
                    loop {
                        if Local::now().minute() >= args[2].parse().expect("error") {
                            return;
                        } else {
                            sleep(Duration::from_millis(1000));
                        }
                    }
                } else {
                    sleep(Duration::from_millis(5000));
                }
            }
        } else if args.len() == 2 {
            loop {
                if Local::now().hour() >= args[1].parse().expect("error") {
                    return;
                } else {
                    sleep(Duration::from_millis(1000));
                }
            }
        }
    }
}
