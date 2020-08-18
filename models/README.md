# CoinNetwork

This package contains

- **Network type**: This is used mainly by CoinToCurrency method to set up the map value for var network map in CoinToCurrency. This map makes it quicker to select specific cryptocurrency and get the API url for it. It expects a string channel and an error channel as parameters for running concurrently, in order to communicate efficiently

- **CoinQuery type**: This is used to define the structure type for the required Query parameters. Attached to it is the CoinToCurrency method which relies on its parameters to work

- **exchangeRates type**: Is a type structure created to match the Json data returned from the exchange rates API for consumption

- **CoinToCurrency method**: This method does as its' name says, it converts a cryptocoin to USD currency. Available Crypto options are in the network map. After matching a key using the Name value the Coinquery Struct it makes its' request. It runs concurrently and needs a string channel and error channel to sync and communicate properly

- **CoinExchangeRate method**: This method simply returns the exchange rate for USD to Bitcoin. Other currencies can be added too. It also runs concurrently
