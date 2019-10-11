********************************************************************************
Go Snowflake Driver
********************************************************************************

.. image:: https://travis-ci.org/snowflakedb/gosnowflake.svg?branch=master
    :target: https://travis-ci.org/snowflakedb/gosnowflake

.. image:: https://codecov.io/gh/snowflakedb/gosnowflake/branch/master/graph/badge.svg
    :target: https://codecov.io/gh/snowflakedb/gosnowflake

.. image:: http://img.shields.io/:license-Apache%202-brightgreen.svg
    :target: http://www.apache.org/licenses/LICENSE-2.0.txt

.. image:: https://goreportcard.com/badge/github.com/snowflakedb/gosnowflake
    :target: https://goreportcard.com/report/github.com/snowflakedb/gosnowflake

This topic provides instructions for installing, running, and modifying the Go Snowflake Driver. The driver supports Go's `database/sql <https://golang.org/pkg/database/sql/>`_ package.

Prerequisites
================================================================================

The following software packages are required to use the Go Snowflake Driver.

Go
----------------------------------------------------------------------

The latest driver requires the `Go language <https://golang.org/>`_ 1.11 or higher. The supported operating systems are Linux, Mac OS, and Windows, but you may run the driver on other platforms if the Go language works correctly on those platforms.


Installation
================================================================================

Get Gosnowflake source code and `dep <https://github.com/golang/dep>`_ (dependency managment tool), if not installed, and ensure the dependent libraries are installed.

.. code-block:: bash

    go get -u github.com/snowflakedb/gosnowflake
    go get -u github.com/golang/dep/cmd/dep
    cd $GOPATH/src/github.com/snowflakedb/gosnowflake/
    dep ensure

Logging Settings
--------------------------------------------------------------------

Unless you implement your own glog and would like to plug in that glog to dump messages generated from the Gosnowflake driver, we strongly recommend using snowflake's implementation of glog.

To use snowflake's glog, add the following block to your ``Gopkg.toml``

.. code-block::

    [[override]]
    name = "github.com/golang/glog"
    source = "https://github.com/snowflakedb/glog.git"
    revision = "f5055e6f21ce71153a199e6460664045b3af8e50"

Run ``dep ensure`` to download the snowflake's glog

Docs
====

For detailed documentation and basic usage examples, please see the documentation at
`godoc.org <https://godoc.org/github.com/snowflakedb/gosnowflake/>`_.

Sample Programs
================================================================================

Snowflake provides a set of sample programs to test with. Set the environment variable ``$GOPATH`` to the top directory of your workspace, e.g., ``~/go`` and make certain to 
include ``$GOPATH/bin`` in the environment variable ``$PATH``. Run the ``make`` command to build all sample programs.

.. code-block:: go

    make install

In the following example, the program ``select1.go`` is built and installed in ``$GOPATH/bin`` and can be run from the command line:

.. code-block:: bash

    SNOWFLAKE_TEST_ACCOUNT=<your_account> \
    SNOWFLAKE_TEST_USER=<your_user> \
    SNOWFLAKE_TEST_PASSWORD=<your_password> \
    select1
    Congrats! You have successfully run SELECT 1 with Snowflake DB!

Development
================================================================================

The developer notes are hosted with the source code on `GitHub <https://github.com/snowflakedb/gosnowflake>`_.

Testing Code
----------------------------------------------------------------------

Set the Snowflake connection info in ``parameters.json``:

.. code-block:: json

    {
        "testconnection": {
            "SNOWFLAKE_TEST_USER":      "<your_user>",
            "SNOWFLAKE_TEST_PASSWORD":  "<your_password>",
            "SNOWFLAKE_TEST_ACCOUNT":   "<your_account>",
            "SNOWFLAKE_TEST_WAREHOUSE": "<your_warehouse>",
            "SNOWFLAKE_TEST_DATABASE":  "<your_database>",
            "SNOWFLAKE_TEST_SCHEMA":    "<your_schema>",
            "SNOWFLAKE_TEST_ROLE":      "<your_role>"
        }
    }

Install `jq <https://stedolan.github.io/jq/>`_ so that the parameters can get parsed correctly, and run ``make test`` in your Go development environment:

.. code-block:: bash

    make test

Submitting Pull Requests
----------------------------------------------------------------------

You may use your preferred editor to edit the driver code. Make certain to run ``make fmt lint`` before submitting any pull request to Snowflake. This command formats your source code according to the standard Go style and detects any coding style issues.

Support
----------------------------------------------------------------------

For official support, contact Snowflake support at:
https://support.snowflake.net/

