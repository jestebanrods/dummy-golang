package pkg

const ErrorPkg = "Error pkg"

func CallMe() (string, int) {
	return callMe()
}

func callMe() (string, int) {
	return getCosaError(), 0
}
