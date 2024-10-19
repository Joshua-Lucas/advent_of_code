use std::{env, fs};

// Function to get the puzzle input as a string for a given year and day
pub fn puzzle_input_string(year: &str, day: &str) -> String {
    // Get the current working directory
    let current_dir = env::current_dir().expect("no current dir");

    let directory_path = format!("src/years/year_{}/days", year);
    let day_file = format!("day{}", day);

    // Construct the path to the nested file
    let file_path = current_dir
        .join(directory_path)
        .join(day_file)
        .join("input.txt");

    // Parse input instructions
    let input =
        fs::read_to_string(&file_path).expect("Can't read input file, to figure out the solution");

    input
}
