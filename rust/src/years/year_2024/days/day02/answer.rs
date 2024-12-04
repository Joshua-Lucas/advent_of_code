use std::str::FromStr;

use anyhow::{anyhow, Result};

use crate::utils::inputs::puzzle_input_string;

#[derive(Debug)]
struct Report {
    levels: Vec<i32>,
}

impl FromStr for Report {
    type Err = anyhow::Error;

    fn from_str(s: &str) -> Result<Self> {
        let result: Vec<i32> = s.split(' ').flat_map(|x| str::parse(x)).collect();

        if result.is_empty() {
            return Err(anyhow!(
                "Expected a Report to containing space seperated levels"
            ));
        }

        return Ok(Report { levels: result });
    }
}

impl Report {
    fn is_safe(&self) -> bool {
        let mut is_gradual = true;

        // The levels are all increasing
        let mut is_monotonically_increasing = true;

        for (current, next) in self.levels.iter().zip(self.levels.iter().skip(1)) {
            if current < next {
                // Now check that the increase is gradual
                let difference = next - current;

                if difference >= 1 && difference <= 3 {
                    continue;
                } else {
                    is_gradual = false;
                    break;
                }
            } else {
                is_monotonically_increasing = false;
                break;
            }
        }

        // The levels are all decreasing
        let mut is_monotonically_decreasing = true;

        for (current, next) in self.levels.iter().zip(self.levels.iter().skip(1)) {
            if current > next {
                // Now check that the increase is gradual
                let difference = current - next;

                if difference >= 1 && difference <= 3 {
                    continue;
                } else {
                    is_gradual = false;
                    break;
                }
            } else {
                is_monotonically_decreasing = false;
                break;
            }
        }

        if (is_monotonically_increasing || is_monotonically_decreasing) && is_gradual {
            return true;
        } else {
            return false;
        }

        // Any two adjacent levels differ by at least one and at most three.
    }
}

pub fn run_solution() {
    // Process the input
    let input_string = puzzle_input_string("2024", "02");

    let reports: Vec<Report> = input_string
        .lines()
        .flat_map(|rep| str::parse(rep))
        .filter(|x: &Report| x.is_safe())
        .inspect(|x| println!("{:?}", x))
        .collect();

    //Maybe return this and have the main function print it out.
    println!("Part One Answer: {}", reports.len());
    // Interpret input instructions part two
    println!("Part Two Answer: {}", "");
}
