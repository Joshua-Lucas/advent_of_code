//--- Day 2: Dive! ---
// https://adventofcode.com/2021/day/2

use std::{env, fs};

pub fn run_solution() {
    // Get the current working directory
    let current_dir = env::current_dir().expect("no current dir");

    // Construct the path to the nested file
    let file_path = current_dir
        .join("src/years/year_2021/days")
        .join("day02")
        .join("input.txt");

    // Parse input instructions
    let input =
        fs::read_to_string(&file_path).expect("Can't read input file, to figure out the solution");
    // Interpret input instructions part one
    let mut location = Cordinates { x: 0, y: 0 };

    // Iterates through the instructions and updates location.
    input.lines().for_each(|dir| navigate(dir, &mut location));

    // Get the final location by multiplying the depth by the horizontal
    // position.
    let final_location = location.x * location.y;

    //Maybe return this and have the main function print it out.
    println!("Part One Answer: {}", final_location);
    // Interpret input instructions part two
    println!("Part Two Answer: {}", "");
}

#[derive(Debug)]
struct Cordinates {
    x: i32,
    y: i32,
}

fn navigate(direction: &str, current_cordinates: &mut Cordinates) {
    // Split the direction from the number.
    let (dir, step) = direction.split_once(" ").expect("Inproper direction");

    match dir {
        "forward" => {
            let step = step.parse::<i32>().unwrap_or(0);
            current_cordinates.x += step;
        }
        "up" => {
            let step = step.parse::<i32>().unwrap_or(0);
            current_cordinates.y -= step;
        }
        "down" => {
            let step = step.parse::<i32>().unwrap_or(0);
            current_cordinates.y += step;
        }
        &_ => println!("no instruction"),
    }
}
