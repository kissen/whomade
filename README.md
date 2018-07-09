whomade
=======

Like `whois` but for MAC addresses. Given a MAC address, it tells you
who was responsible for manufacturing a (network) device.

Installation
------------

	$ go get github.com/kissen/whomade
	$ go install github.com/kissen/whomade

Usage
-----

	$ whomade b8:ae:ed:af:fe:92
	b8:ae:ed        Elitegroup Computer Systems Co.,Ltd.

Credit and License
------------------

All hard work is done by the [klauspost/oui library](https://github.com/klauspost/oui).

The code for `whomade` is licensed under the MIT license. See
`LICENSE` for more information.
