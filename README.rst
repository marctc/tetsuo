Tetsuo
======

Tetsuo is an experimental Go library that allows to report events and metrics of your applications. It's like `Kaneda <http://kaneda.readthedocs.io>`_ but for Go.

Usage
~~~~~~~~~~~

First of all, you need to install `Kaneda` package::

    go get github.com/marctc/tetsuo

Then you need a backend in order to keep data in a persistent storage.
The following example it shows how to send metrics with Elasticsearch as a backend:

.. code-block:: go

  package main

  import (
      "github.com/marctc/tetsuo/backends"
  )
  
  func main() {
      mb := backends.NewMongoBackend("webserver", "users")
      m := Metrics{mb}
      m.Gauge("user_vists", 5)
  }
