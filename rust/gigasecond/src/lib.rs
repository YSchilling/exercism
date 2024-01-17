use time::{OffsetDateTime, PrimitiveDateTime as DateTime};

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    //get unix timestamp
    let mut unix = start.assume_utc().unix_timestamp();

    //increase unix timestamp by one gigasecond
    unix += 1_000_000_000;

    //convert back to PrimitiveDateTime
    let offset = OffsetDateTime::from_unix_timestamp(unix).unwrap();
    let date = offset.date();
    let time = offset.time();

    DateTime::new(date, time)
}
