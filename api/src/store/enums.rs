use diesel::deserialize::{self, FromSql, FromSqlRow};
use diesel::expression::AsExpression;
use diesel::pg::Pg;
use diesel::serialize::{self, IsNull, Output, ToSql};
use diesel::sql_types::Text;
use std::io::Write;

#[derive(Debug, Clone, Copy, PartialEq, Eq, AsExpression, FromSqlRow)]
#[diesel(sql_type = Text)]
pub enum EventStatus {
    Ongoing,
    OnHold,
    Ended,
}

impl ToSql<Text, Pg> for EventStatus {
    fn to_sql<'b>(&'b self, out: &mut Output<'b, '_, Pg>) -> serialize::Result {
        let s = match self {
            EventStatus::Ongoing => "ongoing",
            EventStatus::OnHold => "onhold",
            EventStatus::Ended => "ended",
        };
        out.write_all(s.as_bytes())?;
        Ok(IsNull::No)
    }
}

impl FromSql<Text, Pg> for EventStatus {
    fn from_sql(
        bytes: <Pg as diesel::backend::Backend>::RawValue<'_>,
    ) -> deserialize::Result<Self> {
        let s = <String as FromSql<Text, Pg>>::from_sql(bytes)?;
        match s.as_str() {
            "ongoing" => Ok(EventStatus::Ongoing),
            "onhold" => Ok(EventStatus::OnHold),
            "ended" => Ok(EventStatus::Ended),
            _ => Err(format!("Unrecognized enum variant: {}", s).into()),
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, AsExpression, FromSqlRow)]
#[diesel(sql_type = Text)]
pub enum PartyType {
    Light,
    Full,
}

impl ToSql<Text, Pg> for PartyType {
    fn to_sql<'b>(&'b self, out: &mut Output<'b, '_, Pg>) -> serialize::Result {
        let s = match self {
            PartyType::Light => "light",
            PartyType::Full => "full",
        };
        out.write_all(s.as_bytes())?;
        Ok(IsNull::No)
    }
}

impl FromSql<Text, Pg> for PartyType {
    fn from_sql(
        bytes: <Pg as diesel::backend::Backend>::RawValue<'_>,
    ) -> deserialize::Result<Self> {
        let s = <String as FromSql<Text, Pg>>::from_sql(bytes)?;
        match s.as_str() {
            "light" => Ok(PartyType::Light),
            "full" => Ok(PartyType::Full),
            _ => Err(format!("Unrecognized enum variant: {}", s).into()),
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, AsExpression, FromSqlRow)]
#[diesel(sql_type = Text)]
pub enum EventFrequency {
    Weekly,
    BiWeekly,
    Monthly,
    Quarterly,
    Yearly,
}

impl ToSql<Text, Pg> for EventFrequency {
    fn to_sql<'b>(&'b self, out: &mut Output<'b, '_, Pg>) -> serialize::Result {
        let s = match self {
            EventFrequency::Weekly => "weekly",
            EventFrequency::BiWeekly => "bi-weekly",
            EventFrequency::Monthly => "monthly",
            EventFrequency::Quarterly => "quarterly",
            EventFrequency::Yearly => "yearly",
        };
        out.write_all(s.as_bytes())?;
        Ok(IsNull::No)
    }
}

impl FromSql<Text, Pg> for EventFrequency {
    fn from_sql(
        bytes: <Pg as diesel::backend::Backend>::RawValue<'_>,
    ) -> deserialize::Result<Self> {
        let s = <String as FromSql<Text, Pg>>::from_sql(bytes)?;
        match s.as_str() {
            "weekly" => Ok(EventFrequency::Weekly),
            "bi-weekly" => Ok(EventFrequency::BiWeekly),
            "monthly" => Ok(EventFrequency::Monthly),
            "quarterly" => Ok(EventFrequency::Quarterly),
            "yearly" => Ok(EventFrequency::Yearly),
            _ => Err(format!("Unrecognized enum variant: {}", s).into()),
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, AsExpression, FromSqlRow)]
#[diesel(sql_type = Text)]
pub enum PartyRole {
    Tank,
    Healer,
    Dps,
}

impl ToSql<Text, Pg> for PartyRole {
    fn to_sql<'b>(&'b self, out: &mut Output<'b, '_, Pg>) -> serialize::Result {
        let s = match self {
            PartyRole::Tank => "tank",
            PartyRole::Healer => "healer",
            PartyRole::Dps => "dps",
        };
        out.write_all(s.as_bytes())?;
        Ok(IsNull::No)
    }
}

impl FromSql<Text, Pg> for PartyRole {
    fn from_sql(
        bytes: <Pg as diesel::backend::Backend>::RawValue<'_>,
    ) -> deserialize::Result<Self> {
        let s = <String as FromSql<Text, Pg>>::from_sql(bytes)?;
        match s.as_str() {
            "tank" => Ok(PartyRole::Tank),
            "healer" => Ok(PartyRole::Healer),
            "dps" => Ok(PartyRole::Dps),
            _ => Err(format!("Unrecognized enum variant: {}", s).into()),
        }
    }
}
