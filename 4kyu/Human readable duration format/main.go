/* Your task in order to complete this Kata is to write a function which formats a duration, given as a number of seconds, in a human-friendly way.

The function must accept a non-negative integer. If it is zero, it just returns "now". Otherwise, the duration is expressed as a combination of
years, days, hours, minutes and seconds.

It is much easier to understand with an example:

* For seconds = 62, your function should return
    "1 minute and 2 seconds"
* For seconds = 3662, your function should return
    "1 hour, 1 minute and 2 seconds"
For the purpose of this Kata, a year is 365 days and a day is 24 hours.

Note that spaces are important.

Detailed rules
The resulting expression is made of components like 4 seconds, 1 year, etc. In general, a positive integer and one of the valid units of time, separated by a space.
The unit of time is used in plural if the integer is greater than 1.

The components are separated by a comma and a space (", "). Except the last component, which is separated by " and ", just like it would be written in English.

A more significant units of time will occur before than a least significant one. Therefore, 1 second and 1 year is not correct, but 1 year and 1 second is.

Different components have different unit of times. So there is not repeated units like in 5 seconds and 1 second.

A component will not appear at all if its value happens to be zero. Hence, 1 minute and 0 seconds is not valid, but it should be just 1 minute.

A unit of time must be used "as much as possible". It means that the function should not return 61 seconds, but 1 minute and 1 second instead. Formally,
the duration specified by of a component must not be greater than any valid more significant unit of time. */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FormatDuration(1))
	fmt.Println(FormatDuration(62))
	fmt.Println(FormatDuration(120))
	fmt.Println(FormatDuration(3662))
	fmt.Println(FormatDuration(15731080))
}

func FormatDuration(seconds int64) string {
	year := seconds / 31536000
	seconds -= year * 31536000
	day := seconds / 86400
	seconds -= day * 86400
	hour := seconds / 3600
	seconds -= hour * 3600
	minute := seconds / 60
	seconds -= minute * 60

	arr := make([]string, 5)
	var i, count int

	if year > 0 {
		arr[i] = strconv.FormatInt(year, 10) + " years"
		if year == 1 {
			arr[i] = strconv.FormatInt(year, 10) + " year"
		}
		count++
		i++
	}
	if day > 0 {
		arr[i] = strconv.FormatInt(day, 10) + " days"
		if day == 1 {
			arr[i] = strconv.FormatInt(day, 10) + " day"
		}
		count++
		i++
	}
	if hour > 0 {
		arr[i] = strconv.FormatInt(hour, 10) + " hours"
		if hour == 1 {
			arr[i] = strconv.FormatInt(hour, 10) + " hour"
		}
		count++
		i++
	}
	if minute > 0 {
		arr[i] = strconv.FormatInt(minute, 10) + " minutes"
		if minute == 1 {
			arr[i] = strconv.FormatInt(minute, 10) + " minute"
		}
		count++
		i++
	}
	if seconds > 0 {
		arr[i] = strconv.FormatInt(seconds, 10) + " seconds"
		if seconds == 1 {
			arr[i] = strconv.FormatInt(seconds, 10) + " second"
		}
		count++
		i++
	}

	switch count {
	case 1:
		return arr[0]
	case 2:
		return arr[0] + " and " + arr[1]
	case 3:
		return arr[0] + ", " + arr[1] + " and " + arr[2]
	case 4:
		return arr[0] + ", " + arr[1] + ", " + arr[2] + " and " + arr[3]
	case 5:
		return arr[0] + ", " + arr[1] + ", " + arr[2] + ", " + arr[3] + " and " + arr[4]
	}

	return "now"
}
