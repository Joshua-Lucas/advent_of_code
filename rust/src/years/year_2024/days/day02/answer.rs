use std::str::FromStr;

use anyhow::{anyhow, Result};

use crate::utils::inputs::puzzle_input_string;

#[derive(Debug, Clone)]
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
        // The levels are all increasing
        let is_safe = match (self.levels[0], self.levels[1]) {
            (a, b) if a < b => {
                let mut safe = true;
                for (current, next) in self.levels.iter().zip(self.levels.iter().skip(1)) {
                    if current < next {
                        // Now check that the increase is gradual
                        let difference = next - current;

                        if difference >= 1 && difference <= 3 {
                            continue;
                        } else {
                            safe = false;
                            break;
                        }
                    } else {
                        safe = false;
                        break;
                    }
                }
                return safe;
            }
            (a, b) if a > b => {
                let mut safe = true;
                for (current, next) in self.levels.iter().zip(self.levels.iter().skip(1)) {
                    if current > next {
                        // Now check that the increase is gradual
                        let difference = current - next;

                        if difference >= 1 && difference <= 3 {
                            continue;
                        } else {
                            safe = false;
                            break;
                        }
                    } else {
                        safe = false;
                        break;
                    }
                }

                return safe;
            }
            _ => false,
        };

        return is_safe;
    }

    //Removes a single unsafe condition
    fn is_safe_after_dampener(&mut self) -> Self {
        match (&self.levels[0], &self.levels[1]) {
            (a, b) if a < b => {
                println!("increase");
                let mut index = 0;
                // Iterate through the vector using enumerate to get both index and element
                while index < self.levels.len() - 1 {
                    let current = &self.levels[index];
                    let next = &self.levels[index + 1];

                    let difference = next - current;
                    println!("{}", difference);
                    // Check if the current element is greater than the next
                    // this makes the report unsafe since the levels should
                    // increase in value
                    if current > next || (difference < 1 || difference > 3) {
                        println!("Got here");
                        // Remove the element at the next index (index + 1)
                        self.levels.remove(index);

                        // break when a item is removed.
                        break;

                        // After removal, we don't increment `index` since we want to check the next element
                        // in the same position (as the next element was shifted).
                    } else {
                        // If not, just move to the next index
                        index += 1;
                    }
                }
            }
            (a, b) if a > b => {
                println!("decrease");
                let mut index = 0;
                // Iterate through the vector using enumerate to get both index and element
                while index < self.levels.len() - 1 {
                    let current = &self.levels[index];
                    let next = &self.levels[index + 1];

                    let difference = current - next;

                    println!("{}", difference);

                    // Check if the current element is greater than the next
                    // this makes the report unsafe since the levels should
                    // increase in value
                    if current < next || (difference < 1 || difference > 3) {
                        // Remove the element at the next index (index + 1)
                        self.levels.remove(index);

                        // break when a item is removed.
                        break;

                        // After removal, we don't increment `index` since we want to check the next element
                        // in the same position (as the next element was shifted).
                    } else {
                        // If not, just move to the next index
                        index += 1;
                    }
                }
            }
            (a, b) if a == b => {
                self.levels.remove(0);
            }
            _ => {}
        };

        return self.clone();
    }
}

pub fn run_solution() {
    // Process the input
    let input_string = puzzle_input_string("2024", "02");

    let reports_part_one = input_string
        .lines()
        .flat_map(|rep| str::parse(rep))
        .filter(|x: &Report| x.is_safe())
        .count();

    let reports_part_two = input_string
        .lines()
        .flat_map(|rep| str::parse(rep))
        .map(|mut report: Report| {
            // If the report is not safe we can remove the first unsafe level
            if !report.is_safe() {
                println!("Before mutation: {:?}", report);
                report.is_safe_after_dampener(); // Mutate the report
                println!("After mutation: {:?}", report);
            }
            report // Return the mutated report
        })
        .inspect(|x| println!("{:?}", x.is_safe()))
        .filter(|x| x.is_safe()) // Filter the reports after mutation
        .count();

    //Maybe return this and have the main function print it out.
    println!("Part One Answer: {}", reports_part_one);
    // Interpret input instructions part two
    println!("Part Two Answer: {}", reports_part_two);
}
