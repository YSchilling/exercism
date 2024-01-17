// in minutes
const HOUR: i32 = 60;
const DAY: i32 = 60 * 24;

#[derive(Debug, PartialEq)]
pub struct Clock {
    minutes: u32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        Self {
            minutes: ((((hours * HOUR + minutes) % DAY) + DAY) % DAY) as u32,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(0, self.minutes as i32 + minutes)
    }
}

impl std::fmt::Display for Clock {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "{:0>2}:{:0>2}",
            self.minutes / HOUR as u32,
            self.minutes % HOUR as u32
        )
    }
}
