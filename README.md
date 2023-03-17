# Timeseries util
`timeseries` is equal to `ts`

## Commands

```
timeseries start Refactoring;
timeseries stop;
timeseries status;
timeseries list;
timeseries list --projects;
ts list --tasks --period=<PERIOD>;
ts list --series --period=<PERIOD>;
ts list --projects --tasks --series --period=<PERIOD>;
ts list --csv-out="./report.csv"
```

## Periods
TODAY
YESTERDAY
THIS_WEEK
LAST_WEEK
THIS_MONTH
LAST_MONTH

FROM TODAY - 7 TO TODAY
FROM 01.06.2022 TO TODAY
FROM THIS_MONTH - 3 TO THIS_MONTH

## Examples
`--tasks` closure is used by default

```
ts list --period="FROM TODAY - 3 TO TODAY"
Project1: Task                                                                                6 hours 32 min
Project1: Task2ryf9248ry3n9c238ru4f4hf948hf84h3f3498fh398hf389hd9hqf89ufpa94ufoae8f948eauf    15 min

Total: 6 hours 47 min / $271.33
```