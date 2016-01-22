# bacon

## What is bacon?

Delicious strips of cured meat!

## What is _this_ bacon?

Bacon will be a simple `HTTP` based server that will issue commands to the host. It’s a continuous integration model built on top of `SCM` and `HTTP`, that buys in to the notion of archetecture as software.

The basic idea  is that you specify a list of methods in a `Baconfile` - which is a `JSON` structure. This specification is called a `party` and once you’ve given a server (or `piggy`) a the `Baconfile` the `piggy` becomes the `party` `leader`. Then other little `piggies` can request to `join` the `party` over a `JSON HTTP` interface (I think it will actually be `HTTPS`, it should all be wrapped up in `SSL` from the start.)

Once the `piggies` have joined the `party` the `party` `leader` can issue commands to the `party` `members` with the `methods` from the `Baconfile`.

## License

```
Copyright (c) 2016, Myke Atkinson
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```
