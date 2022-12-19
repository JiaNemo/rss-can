# RSS Can / RSS 罐头

[![CodeQL](https://github.com/soulteary/RSS-Can/actions/workflows/codeql.yml/badge.svg)](https://github.com/soulteary/RSS-Can/actions/workflows/codeql.yml) ![Go Report Card](https://goreportcard.com/badge/github.com/soulteary/RSS-Can)


📰 🥫 **Got RSS CAN be better and simple.**

<p style="text-align: center;">
  <img src="./assets/images/project.jpg" width="300">
</p>

*image generated by stable diffusion*


## PLAN

- Dockerize: Provide a simple tutorial on how to use Docker images with common technology stacks
- Stateful: Persistent data storage solution, supporting at least two databases
- JS SDK: Games and Fun for everyone
- Golang: Optimized for Golang 1.19
- Pipeline: Support RSS pipeline flow, customize information processing tasks and integrate other open-source software
- AI: NLP tasks
- Rules: Support merge open-source software rules: [rss-bridge](https://github.com/RSS-Bridge/rss-bridge/tree/master/bridges) / [RSSHub](https://github.com/DIYgod/RSSHub/tree/master/lib)
- Tools: Quick RSS rules generator, like: [damoeb/rss-proxy](https://github.com/damoeb/rss-proxy)
- [x] 2022.12.20 Support dynamic loading rules.
- [x] 2022.12.19 Support document charset auto detection.
- [x] 2022.12.19 Support MIX parser, both use CSR and SSR parser, faster batch CSR processing.
- [x] 2022.12.19 Support extract combined data from multiple pages and assemble detailed RSS feed data.
- [x] 2022.12.15 Support websites parsing via CSR render, [Blog](https://soulteary.io/2022/12/15/rsscan-use-golang-rod-to-parse-the-content-dynamically-rendered-in-the-browser-part-4.html)
- [x] 2022.12.14 Support convert website page as RSS feeds, [Blog](https://soulteary.com/2022/12/14/rsscan-convert-website-information-stream-to-rss-feed-part-3.html)
- [x] 2022.12.13 Support dynamic rule capability, [Blog](https://soulteary.com/2022/12/13/rsscan-make-golang-applications-with-v8-part-2.html)
- [x] 2022.12.12 Support websites parsing via SSR render, [Blog](https://soulteary.com/2022/12/12/rsscan-better-rsshub-service-build-with-golang-part-1.html)

## License

This project is licensed under the [MIT License](https://github.com/soulteary/RSS-Can/blob/main/LICENSE)

## Credits

- [@PuerkitoBio](https://github.com/PuerkitoBio), He implements a good DOM parsing tool library [goquery](https://github.com/PuerkitoBio/goquery) for Go under the [BSD-3-Clause license](https://github.com/PuerkitoBio/goquery/blob/master/LICENSE). In the project, it is used as a SSR method to parse remote document data. Because there is no Release for the new version, the code base used by the project is [[#3b7929a](https://github.com/PuerkitoBio/goquery/commit/3b7929a0d759a20968ba605c56bc3027c30d3527)].
- [@andybalholm](https://github.com/andybalholm), He implements a Go implementation of a CSS selector library [cascadia](https://github.com/andybalholm/cascadia), which is the core dependency of goquery under the [BSD-2-Clause license](https://github.com/andybalholm/cascadia/blob/master/LICENSE). Because there is no Release for the new version, the code base used by the project is [[#c6065e4](https://github.com/andybalholm/cascadia/commit/c6065e4618b7f538edf5ca0d6b5b2fd0fe129fdd)]
- [@rogchap](https://github.com/rogchap), He implements a good JavaScript runtime library [https://github.com/rogchap/v8go](https://github.com/rogchap/v8go) under the [BSD-3-Clause license](https://github.com/rogchap/v8go/blob/master/LICENSE). In the project, it used as a dynamic configuration execution sandbox environment with version [[v0.7.0](https://github.com/rogchap/v8go/releases/tag/v0.7.0)].
- [@gorilla](https://github.com/gorilla), Gorilla Web Toolkit Dev Team, they offer an amazing library of great tools, eg. [gorilla/feeds](https://github.com/gorilla/feeds) an tiny RSS generator library under the [BSD-2-Clause license](https://github.com/gorilla/feeds/blob/master/LICENSE). In the project, it is used as RSS generator. Sadly, the team decided to archive all projects on December 9th, 2022, the code base used by the project is [b60f215](https://github.com/gorilla/feeds/commit/b60f215f72c708b0800622c804167bea85539ea5).
- [@gin-gonic](https://github.com/gin-gonic), Gin-Gonic Dev Team, they offer an great HTTP web framework [gin](https://github.com/gin-gonic/gin) under the [MIT license](https://github.com/gin-gonic/gin/blob/master/LICENSE). In the project, it used as Web Server to provides RSS API. The code base is [v1.8.1](https://github.com/gin-gonic/gin/releases/tag/v1.8.1).
- [@go-rod](https://github.com/go-rod/rod), Go-Rod Dev Team, they offer an tiny and high-performance CDP driver [go-rod](https://github.com/go-rod/rod) under the [MIT license](https://github.com/go-rod/rod/blob/master/LICENSE). In the project, it is used as CSR parser processing the content dynamically rendered in the browser. The code base is [v0.112.2](https://github.com/go-rod/rod/releases/tag/v0.112.2).
- [@jquery](https://github.com/jquery/jquery), the them offer an great JavaScript library [jquery](https://github.com/jquery/jquery) under the [MIT license](https://github.com/jquery/jquery/blob/main/LICENSE.txt). In the project, it used as CSR in-browser helper, to helper user simply complete the element positioning and information processing in the page. The code is [v1.12.4](https://github.com/jquery/jquery/releases/tag/1.12.4), avoid affecting the execution of the original program of the page after injecting the page, if the page also relies on the same program.
- [@JohannesKaufmann](https://github.com/JohannesKaufmann), He implements a Good HTML to Markdown converter [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) under the [MIT license](https://github.com/JohannesKaufmann/html-to-markdown/blob/master/LICENSE). In the project, it used for enhanced content processing in the SSR phase to generate content in a clean Markdown format. The code base is [v1.3.6](https://github.com/JohannesKaufmann/html-to-markdown/releases/tag/v1.3.6).