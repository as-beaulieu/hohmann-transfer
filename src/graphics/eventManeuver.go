package graphics

import "fmt"

func EarthLaunch() {
	borderedFiller()
	fmt.Println("|  You are ready for launch!                                                     |")
	fmt.Println("|  You use 5 fuel!                                                               |")
	borderedFiller()
}

func EarthLaunchFailure() {
	borderedFiller()
	fmt.Println("|        You did not have enough fuel for launch!                                |")
	fmt.Println("|        Your rocket does not launch from the pad!                               |")
	borderedFiller()
}

func EarthOrbitFailure() {
	borderedFiller()
	fmt.Println("|        You did not have enough fuel for orbit!                                 |")
	fmt.Println("|        Your rocket burns up in reentry                                         |")
	borderedFiller()
}

func MarsEnterOrbit() {
	borderedFiller()
	fmt.Println("|  You begin your orbital burn for Mars!                                         |")
	fmt.Println("|  You use 5 fuel!                                                               |")
	borderedFiller()
}

func ManeuverOffCourse() {
	borderedFiller()
	fmt.Println("   After the maneuver, your crew reports that they are off course.   ")
	borderedFiller()
	fmt.Println("  What would you like to do? ")
	fmt.Println("  (1) - Trust your pilot to make a correction maneuver.                ")
	fmt.Println("  (2) - Trust your mission commander to adjust the remaining maneuvers to correct. ")
}
