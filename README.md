# Package confeito provides fast ensemble tree inference
![confeito logo](https://rawgit.com/hiro4bbh/confeito/master/logo.svg)

[![Build Status](https://travis-ci.org/hiro4bbh/confeito.svg?branch=master)](https://travis-ci.org/hiro4bbh/confeito)

Copyright 2017- Tatsuhiro Aoshima (hiro4bbh@gmail.com).

# Abstract
Package confeito provides fast ensemble tree inference.
See documents on [GoDoc](https://godoc.org/github.com/hiro4bbh/confeito) for details.

This package is based on QuickScorer __[Lucchese+ 2015]__.
We confirmed that the QuickScorer-based implementation is several times faster than the naive implementation on 65536 depth-12 trees with 65536D sparse feature.

# References
- __[Lucchese+ 2015]__ C. Lucchese, F. M. Nardini, S. Orlando, R. Perego, N. Tonellotto, and R. Venturini. "QuickScorer: A Fast Algorithm to Rank Documents with Additive Ensembles of Regression Trees." _Proceedings of the 38th International ACM SIGIR Conference on Research and Development in Information Retrieval_. ACM, 2015.
