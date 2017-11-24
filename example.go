package gocurrency

func _() {

	/* Example using static methods */
	{
		//Money value to start with
		money := 13.37

		//Convert from euros to us dollars (USD)
		money = ConvertEURTo("USD", money)

		//Convert from us dollars (USD) to singapore dollars (SGD)
		money = ConvertFromTo("USD", "SGD", money)

		//Convert from singapore dollars (SGD) to euros (EUR)
		money = ConvertToEUR("SGD", money)

		//Check if the money is valid (valid means not 0 and not infinite)
		println(IsNumberValid(money))

		//Print the number formatted with 2 floating points
		println(FormatNumber(money, 2))
	}

	/* Example using Money type */
	{
		//Money object to start with, requires the money value and the currency of the money
		tmon := Money{Value: 13.37, Currency: "EUR"}

		//Convert from euros to us dollars (USD)
		tmon.To("USD")

		//Convert from us dollars (USD) to singapore dollars (SGD)
		tmon.To("SGD")

		//Convert from singapore dollars (SGD) to euros (EUR)
		tmon.ToEUR()

		//Check if the money is valid (valid means not 0 and not infinite)
		println(tmon.IsValid())

		//Print the number formatted with 2 floating points
		println(tmon.Format(2))
	}

}
