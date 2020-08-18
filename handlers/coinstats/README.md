# CoinStats Handler

this handler receives the amount to be converted in USD and returns the converted value in btc along with the current exchange rate. In order to do this, it makes two API calls which run concurrently and sync using channels.
It expects 3 query parameters with the field names:

- name
- currency
- amount
Amount gets parsed into a Float64 value, just to be sure an actua amount was passed. This handler depends on various import from models/coinNetwork.go to work. See the [readme](../../models/README.md)
