use std::fmt;

const MINUTES_DAY: i64 = 60*24;

#[derive(Debug)]
pub struct Clock {
    minutes: i64,
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let (h,m) = self.hour_minute();
        write!(f, "{:02}:{:02}", h, m)
    }
}

impl PartialEq for Clock {
    fn eq(&self, other: &Self) -> bool {
        self.minutes == other.minutes
    }
}

impl Clock {

    fn hour_minute(&self) -> (i64, i64) {
        (self.minutes / 60, self.minutes % 60)
    }

    pub fn new(hours: i64, minutes: i64) -> Self {
        Clock {
            minutes: 0,
        }.add_minutes(hours*60 + minutes) 
    }

    pub fn add_minutes(&self, minutes: i64) -> Self {
        let time_to_add = (self.minutes + minutes) % MINUTES_DAY;
        let finalminutes = if time_to_add < 0 {
            MINUTES_DAY + time_to_add
        } else {
            (self.minutes + minutes) % MINUTES_DAY
        };
        Clock { minutes: finalminutes }
    }
}
