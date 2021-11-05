========================
Golang Employee Open API
========================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/go-employee-api?status.svg
   :target: https://godoc.org/github.com/siongui/go-employee-api

.. image:: https://github.com/siongui/go-employee-api/workflows/ci/badge.svg
    :target: https://github.com/siongui/go-employee-api/blob/master/.github/workflows/ci.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/go-employee-api
   :target: https://goreportcard.com/report/github.com/siongui/go-employee-api

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/go-employee-api/blob/master/UNLICENSE


Create/Update/Delete API by Go_.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17`_


Requirement
+++++++++++

- Create API

  * Request - RequestId, Emp name, age, address, gender, department, mobile number (The mandatory fields are based on your assumption)
  * Response - RequestId, Errorcode, Errordesc, empid, empname

- Update API

  * Request - RequestId, Emp id, Emp name, age, address, gender, department, mobile number
  * Response - RequestId, Errorcode, Errordesc, empid

- Delete API

  * Request - RequestId, Emp id
  * Response - RequestId, Errorcode, Errordesc, empid

- Use Postgres DB, bun_ to connect DB.
- Log debug / info is required. (Using logrus_)
- Unit Testing is required.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] | `golang crud api - Google search <https://www.google.com/search?q=golang+crud+api>`_
       | `golang crud api - DuckDuckGo search <https://duckduckgo.com/?q=golang+crud+api>`_
       | `golang crud api - Ecosia search <https://www.ecosia.org/search?q=golang+crud+api>`_
       | `golang crud api - Qwant search <https://www.qwant.com/?q=golang+crud+api>`_
       | `golang crud api - Bing search <https://www.bing.com/search?q=golang+crud+api>`_
       | `golang crud api - Yahoo search <https://search.yahoo.com/search?p=golang+crud+api>`_
       | `golang crud api - Baidu search <https://www.baidu.com/s?wd=golang+crud+api>`_
       | `golang crud api - Yandex search <https://www.yandex.com/search/?text=golang+crud+api>`_


.. _Go: https://golang.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
.. _bun: https://github.com/uptrace/bun
.. _logrus: https://github.com/sirupsen/logrus
